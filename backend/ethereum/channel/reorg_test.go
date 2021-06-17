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

package channel_test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"perun.network/go-perun/backend/ethereum/channel"
	"perun.network/go-perun/backend/ethereum/channel/test"
	pkgtest "perun.network/go-perun/pkg/test"
)

// TestSimBackend_Reorg tests the reorg capability of the SimBackend.
// It checks the assumptions that we make about how go-ethereum handles
// chain reorganizations.
func TestSimBackend_Reorg(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// These tests do not work with finality depth > 1.
	oldFd := channel.TxFinalityDepth
	channel.TxFinalityDepth = 1
	defer func() {
		channel.TxFinalityDepth = oldFd
	}()

	t.Run("same-order", func(t *testing.T) {
		rng := pkgtest.Prng(t)
		s := test.NewTokenSetup(ctx, t, rng)
		s.StartSubs()
		defer s.CloseSubs()

		// Send TXs.
		tx1, tx2 := sendTXs(ctx, s)
		// Reorg
		require.NoError(t, s.SB.Reorg(ctx, 2, 3, func(txs []types.Transactions) []types.Transactions {
			return txs
		}))

		// Both TX still valid.
		s.ConfirmTx(tx1, true)
		s.ConfirmTx(tx2, true)
		// No events emitted.
		s.NoMoreEvents()
	})
	t.Run("remove-approval", func(t *testing.T) {
		rng := pkgtest.Prng(t)
		s := test.NewTokenSetup(ctx, t, rng)
		s.StartSubs()
		defer s.CloseSubs()

		// Send TXs.
		tx1, tx2 := sendTXs(ctx, s)
		// Reorg, remove Approval.
		require.NoError(t, s.SB.Reorg(ctx, 2, 3, func(txs []types.Transactions) []types.Transactions {
			return []types.Transactions{nil, nil, txs[1]}
		}))
		// Both TX invalid.
		s.ConfirmTx(tx1, false)
		s.ConfirmTx(tx2, false)
		// All events removed.
		s.AllowanceEvent(1, false)
		s.TransferEvent(false)
		s.AllowanceEvent(0, false)
		s.NoMoreEvents()
	})
	t.Run("remove-transfer", func(t *testing.T) {
		rng := pkgtest.Prng(t)
		s := test.NewTokenSetup(ctx, t, rng)
		s.StartSubs()
		defer s.CloseSubs()

		// Send TXs.
		tx1, tx2 := sendTXs(ctx, s)
		// Reorg, remove transfer.
		require.NoError(t, s.SB.Reorg(ctx, 2, 3, func(txs []types.Transactions) []types.Transactions {
			return []types.Transactions{txs[0]}
		}))
		// Wait for approval, but not transfer.
		s.ConfirmTx(tx1, true)
		s.ConfirmTx(tx2, false)
		// Check that the transfer rollbacked.
		s.TransferEvent(false)
		s.AllowanceEvent(0, false) // allowance is 1 again
		s.NoMoreEvents()
	})
	t.Run("remove-transfer-rebirth", func(t *testing.T) {
		rng := pkgtest.Prng(t)
		s := test.NewTokenSetup(ctx, t, rng)
		s.StartSubs()
		defer s.CloseSubs()

		// Send TXs.
		tx1, tx2 := sendTXs(ctx, s)
		// Reorg, remove transfer.
		require.NoError(t, s.SB.Reorg(ctx, 2, 3, func(txs []types.Transactions) []types.Transactions {
			return []types.Transactions{nil, nil, txs[0]}
		}))
		// Wait for approval, but not transfer.
		s.ConfirmTx(tx1, true)
		s.ConfirmTx(tx2, false)
		// Check that the transfer rollbacked.
		s.AllowanceEvent(1, false) // allowance is 0
		s.TransferEvent(false)
		s.AllowanceEvent(0, false)
		// Rebirth
		s.AllowanceEvent(1, true) // allowance is 1 again
		s.NoMoreEvents()
	})
	t.Run("remove-both", func(t *testing.T) {
		rng := pkgtest.Prng(t)
		s := test.NewTokenSetup(ctx, t, rng)
		s.StartSubs()
		defer s.CloseSubs()

		// Send TXs.
		tx1, tx2 := sendTXs(ctx, s)
		// Reorg, remove transfer.
		require.NoError(t, s.SB.Reorg(ctx, 2, 3, func(txs []types.Transactions) []types.Transactions {
			return nil
		}))
		// Both TX invalid.
		s.ConfirmTx(tx1, false)
		s.ConfirmTx(tx2, false)
		// All events removed.
		s.AllowanceEvent(1, false)
		s.TransferEvent(false)
		s.AllowanceEvent(0, false)
		s.NoMoreEvents()
	})
	t.Run("switch-tx", func(t *testing.T) {
		rng := pkgtest.Prng(t)
		s := test.NewTokenSetup(ctx, t, rng)
		s.StartSubs()
		defer s.CloseSubs()

		// Send TXs.
		tx1, tx2 := sendTXs(ctx, s)
		// Reorg, remove transfer.
		require.NoError(t, s.SB.Reorg(ctx, 2, 3, func(txs []types.Transactions) []types.Transactions {
			return []types.Transactions{txs[1], txs[0]}
		}))
		// Wait for approval, but not transfer.
		s.ConfirmTx(tx1, true)
		s.ConfirmTx(tx2, false)
		// All events removed.
		s.AllowanceEvent(1, false)
		s.TransferEvent(false)
		s.AllowanceEvent(0, false)
		// Allowance event reborn.
		s.AllowanceEvent(1, true)
		s.NoMoreEvents()
	})
}

// sendTXs sends an IncreaseAllowance and a Transfer TX and asserts that both
// and their events are confirmed.
func sendTXs(ctx context.Context, s *test.TokenSetup) (tx1, tx2 *types.Transaction) {
	// Send TXs.
	tx1 = s.IncAllowance(ctx)
	tx2 = s.Transfer(ctx)
	// Wait for TXs confirm.
	s.ConfirmTx(tx1, true)
	s.ConfirmTx(tx2, true)
	// Wait for Events.
	s.AllowanceEvent(1, true) // allowance increase to 1
	s.TransferEvent(true)
	s.AllowanceEvent(0, true) // allowance decrease to 0 after transfer
	s.NoMoreEvents()
	return
}
