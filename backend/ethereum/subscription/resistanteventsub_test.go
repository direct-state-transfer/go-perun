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

package subscription_test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"perun.network/go-perun/backend/ethereum/bindings/peruntoken"
	"perun.network/go-perun/backend/ethereum/channel/test"
	"perun.network/go-perun/backend/ethereum/subscription"
	"perun.network/go-perun/log"
	pctx "perun.network/go-perun/pkg/context"
	pkgtest "perun.network/go-perun/pkg/test"
)

var event = func() *subscription.Event {
	return &subscription.Event{
		Name: "Approval",
		Data: new(peruntoken.PerunTokenApproval),
	}
}

// TestResistantEventSub_Confirmations tests that a TX is confirmed exactly
// after being included in `confirmations` many blocks.
func TestResistantEventSub_Confirmations(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	rng := pkgtest.Prng(t)
	require := require.New(t)
	s := test.NewTokenSetup(ctx, t, rng)

	for i := 0; i < 10; i++ {
		// Create a generic event sub.
		rawSub, err := subscription.NewEventSub(ctx, s.CB, s.Contract, event, 0)
		require.NoError(err)

		confirmations := rng.Int31n(10) + 1
		sub, err := subscription.NewResistantEventSub(ctx, rawSub, s.CB.ContractInterface, uint64(confirmations))
		require.NoError(err)
		log.Debugf("Needed confirmations: %d", confirmations)
		// Send and Confirm the TX. The simulated backend already mined a block here,
		// so the TX has 1 confirmation now.
		s.ConfirmTx(s.IncAllowance(ctx), true)
		// Wait `confirmations-1` blocks.
		for j := int32(0); j < confirmations-1; j++ {
			NoEvent(require, sub)
			log.Debug("Commit")
			s.SB.Commit()
		}
		log.Debug("Waiting for confirm")
		OneEvent(require, sub)
		rawSub.Close()
		sub.Close()
		s.SB.Commit() // Make sure the next EventSub does not get some old block.
	}
}

// TestResistantEventSub_FinalityDepth tests that a TX is confirmed exactly
// after `finalityDepth` blocks even when reorgs occur.
func TestResistantEventSub_Reorg(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	rng := pkgtest.Prng(t)
	require := require.New(t)
	s := test.NewTokenSetup(ctx, t, rng)
	// Create a raw event sub.

	for i := 0; i < 10; i++ {
		rawSub, err := subscription.NewEventSub(ctx, s.CB, s.Contract, event, 0)
		require.NoError(err)

		finality := rng.Int31n(100) + 2
		sub, err := subscription.NewResistantEventSub(ctx, rawSub, s.CB.ContractInterface, uint64(finality))
		require.NoError(err)

		for i := 0; i < 100; i++ {
			s.SB.Commit()
		}
		// Send and Confirm the TX
		s.IncAllowance(ctx)
		// Reorg until the block hight hits `finality` blocks.
		for h := int64(0); h < int64(finality-1); {
			d := rng.Int63n(10) + 1
			l := rng.Int63n(10) + 1
			NoEvent(require, sub)
			log.Debugf("[h=%d] Reorg with depth: %d, length: %d", h, d, l)
			s.SB.Reorg(ctx, uint64(d), uint64(d+l), func(txs []types.Transactions) []types.Transactions {
				return txs
			})
			h += l
		}
		log.Debug("done")
		// Verify that the event arrived.
		OneEvent(require, sub)
		rawSub.Close()
		sub.Close()
	}
}

// TestResistantEventSub_New checks that `NewResistantEventSub` panics for
// `confirmations` < 1.
func TestResistantEventSub_New(t *testing.T) {
	require.PanicsWithValue(t, "confirmations needs to be at least 1", func() {
		subscription.NewResistantEventSub(context.Background(), nil, nil, 0)
	})
}

// NoEvent checks that no event can be read from `sub`.
func NoEvent(require *require.Assertions, sub *subscription.ResistantEventSub) {
	nEvents(require, 0, sub)
}

// OneEvent checks that exactly one event can be read from `sub`.
func OneEvent(require *require.Assertions, sub *subscription.ResistantEventSub) {
	nEvents(require, 1, sub)
}

// nEvents checks that exactly `n` events can be read from `sub`.
func nEvents(require *require.Assertions, n int, sub *subscription.ResistantEventSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	sink := make(chan *subscription.Event, n+1)

	err := sub.Read(ctx, sink)
	require.True(pctx.IsContextError(err))

	for i := 0; i < n; i++ {
		require.NotNil(<-sink)
	}

	select {
	case event := <-sink:
		require.Nil(event)
	default:
	}
}
