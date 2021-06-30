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

package client

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wallet"
	"perun.network/go-perun/wire"
)

// IsVirtualChannel returns whether the channel is a virtual channel.
// A virtual channel is a channel that has a parent channel with different
// participants.
func (c *Channel) IsVirtualChannel() bool {
	return c.Parent() != nil && !c.equalParticipants(c.Parent())
}

func (c *Client) fundVirtualChannel(ctx context.Context, virtual *Channel, prop *VirtualChannelProposal) error {
	parentID := prop.Parents[virtual.Idx()]
	parent, ok := c.channels.Get(parentID)
	if !ok {
		return errors.New("referenced parent channel not found")
	}

	indexMap := prop.IndexMaps[virtual.Idx()]
	err := parent.proposeVirtualChannelFunding(ctx, virtual, indexMap)
	if err != nil {
		return errors.WithMessage(err, "proposing channel funding")
	}

	return c.completeFunding(ctx, virtual)
}

func (c *Channel) proposeVirtualChannelFunding(ctx context.Context, virtual *Channel, indexMap []channel.Index) error {
	// We assume that the channel is locked.

	state := c.state().Clone()
	state.Version++

	// Deposit initial balances into sub-allocation
	balances := virtual.translateBalances(indexMap)
	state.Allocation.Balances = state.Allocation.Balances.Sub(balances)
	state.AddSubAlloc(*channel.NewSubAlloc(virtual.ID(), balances.Sum(), indexMap))

	err := c.updateGeneric(ctx, state, func(mcu *msgChannelUpdate) wire.Msg {
		return &virtualChannelFundingProposal{
			msgChannelUpdate: *mcu,
			Initial: channel.SignedState{
				Params: virtual.Params(),
				State:  virtual.State(),
				Sigs:   virtual.machine.CurrentTX().Sigs,
			},
			IndexMap: indexMap,
		}
	})
	return err
}

const responseTimeout = 10 * time.Second          // How long we wait until the proposal response must be transmitted.
const virtualFundingTimeout = 10 * time.Second    // How long we wait for a matching funding proposal.
const virtualSettlementTimeout = 10 * time.Second // How long we wait for a matching settlement proposal.

func (c *Client) handleVirtualChannelFundingProposal(
	ch *Channel,
	prop *virtualChannelFundingProposal,
	responder *UpdateResponder,
) {
	err := c.validateVirtualChannelFundingProposal(ch, prop)
	if err != nil {
		c.rejectProposal(responder, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Ctx(), virtualFundingTimeout)
	defer cancel()

	err = c.fundingWatcher.Await(ctx, prop)
	if err != nil {
		c.rejectProposal(responder, err.Error())
	}

	c.acceptProposal(responder)
}

// dummyAcount represents an address but cannot be used for signing.
type dummyAccount struct {
	address wallet.Address
}

func (a *dummyAccount) Address() wallet.Address {
	return a.address
}

func (a *dummyAccount) SignData([]byte) ([]byte, error) {
	panic("dummy")
}

const hubIndex = 0 // The hub's index in a virtual channel machine.

func (c *Client) persistVirtualChannel(ctx context.Context, parent *Channel, peers []wire.Address, params channel.Params, state channel.State, sigs []wallet.Sig) (*Channel, error) {
	cID := params.ID()
	if _, err := c.Channel(cID); err == nil {
		return nil, errors.New("channel already exists")
	}

	// We use a dummy account because we don't have the keys to the account.
	ch, err := c.newChannel(&dummyAccount{params.Parts[hubIndex]}, parent, peers, params)
	if err != nil {
		return nil, err
	}

	err = ch.init(ctx, &state.Allocation, state.Data)
	if err != nil {
		return nil, err
	}

	for i, sig := range sigs {
		err = ch.machine.AddSig(ctx, channel.Index(i), sig)
		if err != nil {
			return nil, err
		}
	}

	err = ch.machine.EnableInit(ctx)
	if err != nil {
		return nil, err
	}

	err = ch.machine.SetFunded(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.pr.ChannelCreated(ctx, ch.machine, peers, nil); err != nil {
		return ch, errors.WithMessage(err, "persisting new channel")
	}
	ok := c.channels.Put(cID, ch)
	if !ok {
		return nil, errors.Errorf("failed to put channel into registry: %v", cID)
	}
	return ch, nil
}

func (c *Client) validateVirtualChannelFundingProposal(
	ch *Channel,
	prop *virtualChannelFundingProposal,
) error {
	switch {
	case prop.Initial.Params.ID() != prop.Initial.State.ID:
		return errors.New("state does not match parameters")
	case !prop.Initial.Params.VirtualChannel:
		return errors.New("virtual channel flag not set")
	case len(prop.Initial.State.Locked) > 0:
		return errors.New("cannot have locked funds")
	}

	// Validate signatures.
	for i, sig := range prop.Initial.Sigs {
		ok, err := channel.Verify(
			prop.Initial.Params.Parts[i],
			prop.Initial.Params,
			prop.Initial.State,
			sig,
		)
		if err != nil {
			return err
		} else if !ok {
			return errors.New("invalid signature")
		}
	}

	// Validate index map.
	if len(prop.Initial.Params.Parts) != len(prop.IndexMap) {
		return errors.New("index map: invalid length")
	}

	// Assert not contained before
	_, containedBefore := ch.state().SubAlloc(prop.Initial.Params.ID())
	if containedBefore {
		return errors.New("virtual channel already allocated")
	}

	// Assert contained after with correct balances
	expected := channel.NewSubAlloc(prop.Initial.Params.ID(), prop.Initial.State.Sum(), prop.IndexMap)
	subAlloc, containedAfter := prop.State.SubAlloc(expected.ID)
	if !containedAfter || subAlloc.Equal(expected) != nil {
		return errors.New("invalid allocation")
	}

	// Validate allocation.

	// Assert equal assets.
	if err := channel.AssetsAssertEqual(ch.state().Assets, prop.Initial.State.Assets); err != nil {
		return errors.WithMessage(err, "assets do not match")
	}

	// Assert sufficient funds in parent channel.
	virtual := transformBalances(prop.Initial.State.Balances, ch.state().NumParts(), subAlloc.IndexMap)
	if err := ch.state().Balances.AssertGreaterOrEqual(virtual); err != nil {
		return errors.WithMessage(err, "insufficient funds")
	}

	return nil
}

func (c *Client) matchFundingProposal(ctx context.Context, a, b interface{}) bool {
	var err error
	defer func() {
		if err != nil {
			c.log.Debug("matching funding proposal:", err)
		}
	}()

	// Cast.
	props, err := castToFundingProposals(a, b)
	if err != nil {
		return false
	}
	prop0 := props[0]

	// Check initial state.
	for i, prop := range props {
		if prop.Initial.State.Equal(prop0.Initial.State) != nil {
			err = errors.Errorf("checking state equality %d", i)
			return false
		}
	}

	channels, err := c.gatherChannels(props...)
	if err != nil {
		return false
	}

	// Check index map.
	indices := make([]bool, len(prop0.IndexMap))
	for i, prop := range props {
		for j, idx := range prop.IndexMap {
			if idx == channels[i].Idx() {
				indices[j] = true
			}
		}
	}
	for i, ok := range indices {
		if !ok {
			err = errors.Errorf("checking index map %d", i)
			return false
		}
	}

	// Store state for withdrawal after dispute.
	parent := channels[0]
	peers := c.gatherPeers(channels...)
	_, err = c.persistVirtualChannel(ctx, parent, peers, *prop0.Initial.Params, *prop0.Initial.State, prop0.Initial.Sigs)
	return err == nil
}

func castToFundingProposals(inputs ...interface{}) ([]*virtualChannelFundingProposal, error) {
	var ok bool
	props := make([]*virtualChannelFundingProposal, len(inputs))
	for i, x := range inputs {
		props[i], ok = x.(*virtualChannelFundingProposal)
		if !ok {
			return nil, errors.Errorf("casting %d", i)
		}
	}
	return props, nil
}

func (c *Client) gatherChannels(props ...*virtualChannelFundingProposal) ([]*Channel, error) {
	var err error
	channels := make([]*Channel, len(props))
	for i, prop := range props {
		channels[i], err = c.Channel(prop.ID())
		if err != nil {
			return nil, err
		}
	}
	return channels, nil
}

func (c *Client) gatherPeers(channels ...*Channel) (peers []wire.Address) {
	peers = make([]wire.Address, len(channels))
	for i, ch := range channels {
		chPeers := ch.Peers()
		if len(chPeers) != 2 {
			panic("unsupported number of participants")
		}
		peers[i] = chPeers[1-ch.Idx()]
	}
	return
}