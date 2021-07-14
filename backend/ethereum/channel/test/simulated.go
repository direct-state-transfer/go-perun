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

	"perun.network/go-perun/channel/test"
	"perun.network/go-perun/log"
	"perun.network/go-perun/pkg/sync"
)

// GasLimit is the max amount of gas we want to send per transaction.
const GasLimit = 500000

// SimulatedBackend provides a simulated ethereum blockchain for tests.
type SimulatedBackend struct {
	backends.SimulatedBackend
	faucetKey  *ecdsa.PrivateKey
	faucetAddr common.Address
	clockMu    sync.Mutex    // Mutex for clock adjustments. Locked by SimTimeouts.
	mining     chan struct{} // Used for auto-mining blocks.
}

// NewSimulatedBackend creates a new Simulated Backend.
func NewSimulatedBackend() *SimulatedBackend {
	sk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	faucetAddr := crypto.PubkeyToAddress(sk.PublicKey)

	genesis := core.DeveloperGenesisBlock(15, faucetAddr)
	return &SimulatedBackend{
		SimulatedBackend: *backends.NewSimulatedBackend(genesis.Alloc, genesis.GasLimit),
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
