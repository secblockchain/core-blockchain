// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
// bug across the entire project files fixed and high tx per block feature added  by EtherAuthority <https://etherauthority.io/>

// Package checkpointoracle is a an on-chain light client checkpoint oracle.
package checkpointoracle

//go:generate abigen --sol contract/oracle.sol --pkg contract --out contract/oracle.go

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/checkpointoracle/contract"
	"github.com/ethereum/go-ethereum/core/types"
)

// CheckpointOracle is a Go wrapper around an on-chain checkpoint oracle contract.
type CheckpointOracle struct {
	address  common.Address
	contract *contract.CheckpointOracle
}

// NewCheckpointOracle binds checkpoint contract and returns a registrar instance.
func NewCheckpointOracle(contractAddr common.Address, backend bind.ContractBackend) (*CheckpointOracle, error) {
	c, err := contract.NewCheckpointOracle(contractAddr, backend)
	if err != nil {
		return nil, err
	}
	return &CheckpointOracle{address: contractAddr, contract: c}, nil
}

// ContractAddr returns the address of contract.
func (oracle *CheckpointOracle) ContractAddr() common.Address {
	return oracle.address
}

// Contract returns the underlying contract instance.
func (oracle *CheckpointOracle) Contract() *contract.CheckpointOracle {
	return oracle.contract
}

// LookupCheckpointEvents searches checkpoint event for specific section in the
// given log batches.
func (oracle *CheckpointOracle) LookupCheckpointEvents(blockLogs [][]*types.Log, section uint64, hash common.Hash) []*contract.CheckpointOracleNewCheckpointVote {
	var votes []*contract.CheckpointOracleNewCheckpointVote

	for _, logs := range blockLogs {
		for _, log := range logs {
			event, err := oracle.contract.ParseNewCheckpointVote(*log)
			if err != nil {
				continue
			}
			if event.Index == section && event.CheckpointHash == hash {
				votes = append(votes, event)
			}
		}
	}
	return votes
}

// RegisterCheckpoint registers the checkpoint with a batch of associated signatures
// that are collected off-chain and sorted by lexicographical order.
//
// Notably all signatures given should be transformed to "ethereum style" which transforms
// v from 0/1 to 27/28 according to the yellow paper.
//
// The function now uses EIP-712 signature validation with nonce-based replay protection.
// Each signer must provide their current nonce, which will be verified and incremented.
func (oracle *CheckpointOracle) RegisterCheckpoint(opts *bind.TransactOpts, index uint64, hash []byte, nonces []*big.Int, sigs [][]byte) (*types.Transaction, error) {
	var (
		r [][32]byte
		s [][32]byte
		v []uint8
	)
	
	// Validate that we have the same number of nonces as signatures
	if len(nonces) != len(sigs) {
		return nil, errors.New("number of nonces must match number of signatures")
	}
	
	for i := 0; i < len(sigs); i++ {
		if len(sigs[i]) != 65 {
			return nil, errors.New("invalid signature")
		}
		r = append(r, common.BytesToHash(sigs[i][:32]))
		s = append(s, common.BytesToHash(sigs[i][32:64]))
		v = append(v, sigs[i][64])
	}
	return oracle.contract.SetCheckpoint(opts, common.BytesToHash(hash), index, nonces, v, r, s)
}

// GetAdminNonce retrieves the current nonce for a specific admin address.
// This is used for EIP-712 signature creation and replay protection.
func (oracle *CheckpointOracle) GetAdminNonce(opts *bind.CallOpts, admin common.Address) (*big.Int, error) {
	return oracle.contract.GetAdminNonce(opts, admin)
}

// GetChainId retrieves the chain ID used in the EIP-712 domain separator.
// This is important for cross-chain replay protection.
func (oracle *CheckpointOracle) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	return oracle.contract.ChainId(opts)
}

// GetLatestCheckpoint retrieves the latest checkpoint information.
func (oracle *CheckpointOracle) GetLatestCheckpoint(opts *bind.CallOpts) (uint64, common.Hash, *big.Int, error) {
	return oracle.contract.GetLatestCheckpoint(opts)
}

// GetAllAdmin retrieves all admin addresses registered with the oracle.
func (oracle *CheckpointOracle) GetAllAdmin(opts *bind.CallOpts) ([]common.Address, error) {
	return oracle.contract.GetAllAdmin(opts)
}
