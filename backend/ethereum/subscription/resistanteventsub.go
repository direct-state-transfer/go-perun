// Copyright 2021 - See NOTICE file for copyright holders.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package subscription

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	cherrors "perun.network/go-perun/backend/ethereum/channel/errors"
	"perun.network/go-perun/log"
	pkgsync "perun.network/go-perun/pkg/sync"
)

type (
	// Wraps an `EventSub` and makes it resistant to chain reorgs.
	// It handles `removed` and `rebirth` events and has a `finalityDepth`
	// threshold to decide when an Event is final.
	// Its will never emit the same event twice unless the connected geth node
	// is faulty.
	ResistantEventSub struct {
		closer        pkgsync.Closer
		sub           *EventSub
		confirmations uint64

		lastBlockNum *big.Int
		heads        chan *types.Header
		headSub      ethereum.Subscription
		events       map[common.Hash]*Event
	}
)

// Subscribe is a convenience function which returns a `ResistantEventSub`.
// It is equivalent to manually calling `NewEventSub` and `NewResistantEventSub`
// with the given parameters.
func Subscribe(ctx context.Context, cr ethereum.ChainReader, contract *bind.BoundContract, eFact EventFactory, startBlockOffset, confirmations uint64) (*ResistantEventSub, error) {
	_sub, err := NewEventSub(ctx, cr, contract, eFact, startBlockOffset)
	if err != nil {
		return nil, errors.WithMessage(err, "creating filter-watch event subscription")
	}
	sub, err := NewResistantEventSub(ctx, _sub, cr, confirmations)
	if err != nil {
		_sub.Close()
		return nil, errors.WithMessage(err, "creating filter-watch event subscription")
	}
	return sub, nil
}

// NewResistantEventSub creates a new `ResistantEventSub` from the given
// `EventSub`. Closes the passed `EventSub` when done.
// `confirmations` defines in how many blocks the events needs to be included.
// `confirmations` can not be smaller than 1.
// The passed `EventSub` should query more than `confirmations` blocks into
// the past.
func NewResistantEventSub(ctx context.Context, sub *EventSub, cr ethereum.ChainReader, confirmations uint64) (*ResistantEventSub, error) {
	if confirmations < 1 {
		panic("confirmations needs to be at least 1")
	}
	last, err := cr.HeaderByNumber(ctx, nil)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return nil, errors.WithMessage(err, "subscribing to headers")
	}
	log.Debugf("Resistant Event sub started at block: %v", last.Number)
	// Use a large buffer to not block geth.
	heads := make(chan *types.Header, 128)
	headSub, err := cr.SubscribeNewHead(ctx, heads)
	if err != nil {
		headSub.Unsubscribe()
		err = cherrors.CheckIsChainNotReachableError(err)
		return nil, errors.WithMessage(err, "subscribing to headers")
	}

	ret := &ResistantEventSub{
		sub:           sub,
		lastBlockNum:  last.Number,
		heads:         heads,
		headSub:       headSub,
		confirmations: confirmations,
		events:        make(map[common.Hash]*Event),
	}
	ret.closer.OnCloseAlways(func() {
		headSub.Unsubscribe()
		sub.Close()
	})
	return ret, nil
}

// Read reads all past and future events into `sink`.
// Can be aborted by cancelling `ctx` or `Close()`.
// All events can be considered final.
func (s *ResistantEventSub) Read(_ctx context.Context, sink chan<- *Event) error {
	ctx, cancel := context.WithCancel(_ctx)
	defer cancel()
	subErr := make(chan error, 1)
	rawEvents := make(chan *Event, 128)
	// Read events from the underlying event subscription.
	go func() {
		subErr <- s.sub.Read(ctx, rawEvents)
	}()

	for {
		select {
		case head := <-s.heads:
			if head == nil {
				return errors.New("head sub returned nil")
			}
			s.processHead(head, sink)
		case event := <-rawEvents:
			s.processEvent(event, sink)
		case e := <-s.headSub.Err():
			return errors.WithMessage(e, "underlying head subscription")
		case e := <-subErr:
			if e != nil {
				return errors.WithMessage(e, "underlying EventSub.Read")
			}
			return errors.New("underlying event sub terminated")
		case <-ctx.Done():
			return ctx.Err()
		case <-s.closer.Closed():
			return nil
		}
	}
}

func (s *ResistantEventSub) ReadPast(_ctx context.Context, sink chan<- *Event) error {
	ctx, cancel := context.WithCancel(_ctx)
	defer cancel()
	subErr := make(chan error, 1)
	rawEvents := make(chan *Event, 128)
	// Read events from the underlying event subscription.
	go func() {
		defer close(rawEvents)
		subErr <- s.sub.ReadPast(ctx, rawEvents)
	}()

	for {
		select {
		case head := <-s.heads:
			if head == nil {
				return errors.New("head sub returned nil")
			}
			s.processHead(head, sink)
		case event := <-rawEvents:
			if event == nil {
				return nil
			}
			s.processEvent(event, sink)
		case e := <-s.headSub.Err():
			return errors.WithMessage(e, "underlying head subscription")
		case e := <-subErr:
			if e != nil {
				return errors.WithMessage(e, "underlying EventSub.Read")
			}
			continue
		case <-ctx.Done():
			return ctx.Err()
		case <-s.closer.Closed():
			return nil
		}
	}
}

func (s *ResistantEventSub) processEvent(event *Event, sink chan<- *Event) {
	hash := event.Log.TxHash
	log := log.WithField("hash", hash.Hex())

	if event.Log.Removed {
		log.Trace("Event preliminary excluded")
		delete(s.events, hash)
	} else {
		if s.isFinal(event) {
			sink <- event
			delete(s.events, hash)
		} else {
			log.Trace("Event preliminary included")
			s.events[hash] = event
		}
	}
}

func (s *ResistantEventSub) processHead(head *types.Header, sink chan<- *Event) {
	diff := new(big.Int).Sub(head.Number, s.lastBlockNum)
	log.Tracef("Received new block. From %v to %v, increment: %v", s.lastBlockNum, head.Number, diff)
	s.lastBlockNum.Set(head.Number)

	// Events can only become final if the block number increases.
	if diff.Sign() > 0 {
		for _, event := range s.events {
			if s.isFinal(event) {
				sink <- event
				delete(s.events, event.Log.TxHash)
			}
		}
	} else {
		log.Debugf("Block number did not increase. Increment: %v", diff)
	}
}

func (s *ResistantEventSub) isFinal(event *Event) bool {
	log := log.WithField("hash", event.Log.TxHash.Hex())
	diff := new(big.Int).Sub(s.lastBlockNum, big.NewInt(int64(event.Log.BlockNumber)))

	if diff.Sign() < 0 {
		log.Tracef("Event sub was faster than head sub, ignored")
		return false
	} else {
		included := new(big.Int).Add(diff, big.NewInt(1))
		if included.Cmp(big.NewInt(int64(s.confirmations))) >= 0 {
			log.Debugf("Event final after %d block(s)", included)
			return true
		} else {
			log.Tracef("Event included %d time(s)", included)
			return false
		}
	}
}

// Close closes the sub and frees associated resources.
// Close the underlying `EventSub`.
// Can be called more than once, but should be called at least once.
// Is thread safe.
func (s *ResistantEventSub) Close() {
	s.closer.Close() // noline: errcheck
}
