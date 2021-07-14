// Copyright 2019 - See NOTICE file for copyright holders.
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

package test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/channel/test"
	"perun.network/go-perun/log"
	"perun.network/go-perun/pkg/sync"
)

// GasLimit is the max amount of gas we want to send per transaction.
const GasLimit = 500000

// SimulatedBackend provides a simulated ethereum blockchain for tests.
type SimulatedBackend struct {
	backends.SimulatedBackend
	sbMtx sync.Mutex

	faucetKey  *ecdsa.PrivateKey
	faucetAddr common.Address
	clockMu    sync.Mutex    // Mutex for clock adjustments. Locked by SimTimeouts.
	mining     chan struct{} // Used for auto-mining blocks.
}

// Reorder can be used to insert, reorder and exclude transactions in
// combination with `Reorg`.
type Reordered func([]types.Transactions) []types.Transactions

// NewSimulatedBackend creates a new Simulated Backend.
func NewSimulatedBackend() *SimulatedBackend {
	sk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	faucetAddr := crypto.PubkeyToAddress(sk.PublicKey)
	addr := map[common.Address]core.GenesisAccount{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
		faucetAddr:                       {Balance: new(big.Int).Sub(channel.MaxBalance, big.NewInt(9))},
	}
	alloc := core.GenesisAlloc(addr)
	return &SimulatedBackend{
		SimulatedBackend: *backends.NewSimulatedBackend(alloc, 8000000),
		faucetKey:        sk,
		faucetAddr:       faucetAddr,
	}
}

// SendTransaction executes a transaction.
func (s *SimulatedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if err := s.SimulatedBackend.SendTransaction(ctx, tx); err != nil {
		return errors.WithStack(err)
	}
	s.Commit()
	return nil
}

// FundAddress funds a given address with `test.MaxBalance` eth from a faucet.
func (s *SimulatedBackend) FundAddress(ctx context.Context, addr common.Address) {
	nonce, err := s.PendingNonceAt(context.Background(), s.faucetAddr)
	if err != nil {
		panic(err)
	}
	tx := types.NewTransaction(nonce, addr, test.MaxBalance, GasLimit, big.NewInt(1), nil)
	signer := types.NewEIP155Signer(big.NewInt(1337))
	signedTX, err := types.SignTx(tx, signer, s.faucetKey)
	if err != nil {
		panic(err)
	}
	if err := s.SendTransaction(ctx, signedTX); err != nil {
		panic(err)
	}
	bind.WaitMined(context.Background(), s, signedTX)
}

// StartMining makes the simulated blockchain auto-mine blocks with the given
// interval. Must be stopped with `StopMining`.
// The block time of generated blocks will always increase by 10 seconds.
func (s *SimulatedBackend) StartMining(interval time.Duration) {
	if interval == 0 {
		panic("blockTime can not be zero")
	}

	s.mining = make(chan struct{})
	go func() {
		log.Trace("Started mining")
		defer log.Trace("Stopped mining")

		for {
			s.Commit()
			log.Trace("Mined simulated block")

			select {
			case <-time.After(interval):
			case <-s.mining: // stopped
				return
			}
		}
	}()
}

// StopMining stops the auto-mining of the simulated blockchain.
// Must be called exactly once to free resources iff `StartMining` was called.
func (s *SimulatedBackend) StopMining() {
	close(s.mining)
}

// Reorg simulates a chain reorg with parameters depth and length.
// reorder can be used to insert, reorder and exclude transactions.
// reorder receives a slice of Transactions where each slice entry contains
// the transactions that were contained in the orphaned block.
// The slice that is passed to reorder has length depth, and the one returned
// must have at least length `length`.
// Panics if length is not greater than depth.
// The account nonce prevents transactions of the same account from being re-ordered.
// Trying to do this will panic.
func (s *SimulatedBackend) Reorg(ctx context.Context, depth, length uint64, reorder Reordered) error {
	// The chain inserter always chooses the most difficult chain.
	// In the simulated case we expect that a longer chain is always
	// more difficult and will therefore be picked in a reorg case.
	// This check makes the function deterministic in the length<=depth case.
	if length <= depth {
		panic("reorg length must be greater than depth")
	}
	if !s.sbMtx.TryLockCtx(ctx) {
		return errors.Errorf("locking mutex: %v", ctx.Err())
	}
	defer s.sbMtx.Unlock()

	// parent at current - depth.
	parentN := new(big.Int).Sub(s.Blockchain().CurrentBlock().Number(), big.NewInt(int64(depth)))
	parent, err := s.BlockByNumber(ctx, parentN)
	if err != nil {
		return errors.Wrap(err, "retrieving reorg parent")
	}
	// Collect orphaned TXs.
	txs := make([]types.Transactions, depth)
	for i := uint64(0); i < depth; i++ {
		blockN := new(big.Int).Add(parentN, big.NewInt(int64(i+1)))
		block, err := s.BlockByNumber(ctx, blockN)
		if err != nil {
			return errors.Wrap(err, "retrieving block")
		}
		// Add the TXs from block parent + 1 + i.
		txs[i] = block.Transactions()
	}
	// Modify the TXs with the reorder callback.
	newTXs := reorder(txs)
	if uint64(len(newTXs)) > length {
		panic("more new transactions than blocks")
	}
	// Reset the chain to the parent block.
	if err := s.Fork(ctx, parent.Hash()); err != nil {
		return errors.Wrap(err, "forking")
	}
	// Add the modified TXs to their new blocks, if any.
	for i := 0; i < int(length); i++ {
		if i < len(newTXs) {
			for _, tx := range newTXs[i] {
				if err := s.SimulatedBackend.SendTransaction(ctx, tx); err != nil {
					return errors.Wrap(err, "re-sending transaction")
				}
			}
		}
		s.Commit()
	}

	return nil
}
