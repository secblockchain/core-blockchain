// Copyright 2019 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/contracts/checkpointoracle"
	"github.com/ethereum/go-ethereum/contracts/checkpointoracle/contract"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"gopkg.in/urfave/cli.v1"
)

var commandDeploy = cli.Command{
	Name:  "deploy",
	Usage: "Deploy a new checkpoint oracle contract",
	Flags: []cli.Flag{
		nodeURLFlag,
		clefURLFlag,
		signerFlag,
		signersFlag,
		thresholdFlag,
	},
	Action: utils.MigrateFlags(deploy),
}

var commandSign = cli.Command{
	Name:  "sign",
	Usage: "Sign the checkpoint with the specified key",
	Flags: []cli.Flag{
		nodeURLFlag,
		clefURLFlag,
		signerFlag,
		indexFlag,
		hashFlag,
		oracleFlag,
	},
	Action: utils.MigrateFlags(sign),
}

var commandPublish = cli.Command{
	Name:  "publish",
	Usage: "Publish a checkpoint into the oracle",
	Flags: []cli.Flag{
		nodeURLFlag,
		clefURLFlag,
		signerFlag,
		indexFlag,
		signaturesFlag,
	},
	Action: utils.MigrateFlags(publish),
}

// deploy deploys the checkpoint registrar contract.
//
// Note the network where the contract is deployed depends on
// the network where the connected node is located.
func deploy(ctx *cli.Context) error {
	// Gather all the addresses that should be permitted to sign
	var addrs []common.Address
	for _, account := range strings.Split(ctx.String(signersFlag.Name), ",") {
		if trimmed := strings.TrimSpace(account); !common.IsHexAddress(trimmed) {
			utils.Fatalf("Invalid account in --signers: '%s'", trimmed)
		}
		addrs = append(addrs, common.HexToAddress(account))
	}
	// Retrieve and validate the signing threshold
	needed := ctx.Int(thresholdFlag.Name)
	if needed == 0 || needed > len(addrs) {
		utils.Fatalf("Invalid signature threshold %d", needed)
	}
	// Print a summary to ensure the user understands what they're signing
	fmt.Printf("Deploying new checkpoint oracle:\n\n")
	for i, addr := range addrs {
		fmt.Printf("Admin %d => %s\n", i+1, addr.Hex())
	}
	fmt.Printf("\nSignatures needed to publish: %d\n", needed)

	// setup clef signer, create an abigen transactor and an RPC client
	transactor, client := newClefSigner(ctx), newClient(ctx)

	// Get chain ID from client for EIP-712 domain separator
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		utils.Fatalf("Failed to get chain ID: %v", err)
	}

	// Deploy the checkpoint oracle with chain ID
	fmt.Println("Sending deploy request to Clef...")
	oracle, tx, _, err := contract.DeployCheckpointOracle(transactor, client, addrs, big.NewInt(int64(params.CheckpointFrequency)),
		big.NewInt(int64(params.CheckpointProcessConfirmations)), big.NewInt(int64(needed)), chainID)
	if err != nil {
		utils.Fatalf("Failed to deploy checkpoint oracle %v", err)
	}
	log.Info("Deployed checkpoint oracle", "address", oracle, "tx", tx.Hash().Hex())

	return nil
}

// sign creates the signature for specific checkpoint
// with local key. Only contract admins have the permission to
// sign checkpoint.
func sign(ctx *cli.Context) error {
	var (
		offline bool // The indicator whether we sign checkpoint by offline.
		chash   common.Hash
		cindex  uint64
		address common.Address

		node   *rpc.Client
		oracle *checkpointoracle.CheckpointOracle
		err    error
	)
	if !ctx.GlobalIsSet(nodeURLFlag.Name) {
		// Offline mode signing
		offline = true
		if !ctx.IsSet(hashFlag.Name) {
			utils.Fatalf("Please specify the checkpoint hash (--hash) to sign in offline mode")
		}
		chash = common.HexToHash(ctx.String(hashFlag.Name))

		if !ctx.IsSet(indexFlag.Name) {
			utils.Fatalf("Please specify checkpoint index (--index) to sign in offline mode")
		}
		cindex = ctx.Uint64(indexFlag.Name)

		if !ctx.IsSet(oracleFlag.Name) {
			utils.Fatalf("Please specify oracle address (--oracle) to sign in offline mode")
		}
		address = common.HexToAddress(ctx.String(oracleFlag.Name))
	} else {
		// Interactive mode signing, retrieve the data from the remote node
		node = newRPCClient(ctx.GlobalString(nodeURLFlag.Name))

		checkpoint := getCheckpoint(ctx, node)
		chash, cindex, address = checkpoint.Hash(), checkpoint.SectionIndex, getContractAddr(node)

		// Check the validity of checkpoint
		reqCtx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFn()

		head, err := ethclient.NewClient(node).HeaderByNumber(reqCtx, nil)
		if err != nil {
			return err
		}
		num := head.Number.Uint64()
		if num < ((cindex+1)*params.CheckpointFrequency + params.CheckpointProcessConfirmations) {
			utils.Fatalf("Invalid future checkpoint")
		}
		_, oracle = newContract(node)
		latest, _, h, err := oracle.Contract().GetLatestCheckpoint(nil)
		if err != nil {
			return err
		}
		if cindex < latest {
			utils.Fatalf("Checkpoint is too old")
		}
		if cindex == latest && (latest != 0 || h.Uint64() != 0) {
			utils.Fatalf("Stale checkpoint, latest registered %d, given %d", latest, cindex)
		}
	}
	var (
		signature string
		signer    string
	)
	// isAdmin checks whether the specified signer is admin.
	isAdmin := func(addr common.Address) error {
		signers, err := oracle.Contract().GetAllAdmin(nil)
		if err != nil {
			return err
		}
		for _, s := range signers {
			if s == addr {
				return nil
			}
		}
		return fmt.Errorf("signer %v is not the admin", addr.Hex())
	}
	// Print to the user the data thy are about to sign
	fmt.Printf("Oracle     => %s\n", address.Hex())
	fmt.Printf("Index %4d => %s\n", cindex, chash.Hex())

	// Sign checkpoint in clef mode.
	signer = ctx.String(signerFlag.Name)

	if !offline {
		if err := isAdmin(common.HexToAddress(signer)); err != nil {
			return err
		}
	}

	// Get nonce for EIP-712 signature
	var nonce *big.Int
	if !offline {
		nonce, err = oracle.GetAdminNonce(nil, common.HexToAddress(signer))
		if err != nil {
			utils.Fatalf("Failed to get admin nonce: %v", err)
		}
	} else {
		// For offline mode, we need to get the nonce from the user or use 0
		// This is a limitation of offline mode - the user should provide the current nonce
		nonce = big.NewInt(0)
		fmt.Println("Warning: Using nonce 0 for offline mode. Ensure this is the correct nonce.")
	}

	// Get chain ID for EIP-712 domain separator
	var chainID *big.Int
	if !offline {
		chainID, err = oracle.GetChainId(nil)
		if err != nil {
			utils.Fatalf("Failed to get chain ID: %v", err)
		}
	} else {
		// For offline mode, we need to get the chain ID from the user
		// This is a limitation of offline mode
		chainID = big.NewInt(1) // Default to mainnet
		fmt.Println("Warning: Using chain ID 1 for offline mode. Ensure this is correct.")
	}

	// Create EIP-712 hash
	signedHash := createEIP712Hash(cindex, address, chash, nonce, chainID)

	clef := newRPCClient(ctx.String(clefURLFlag.Name))
	p := make(map[string]string)
	p["address"] = address.Hex()
	p["message"] = hexutil.Encode(signedHash)

	fmt.Println("Sending signing request to Clef...")
	if err := clef.Call(&signature, "account_signData", accounts.MimetypeDataWithValidator, signer, p); err != nil {
		utils.Fatalf("Failed to sign checkpoint, err %v", err)
	}
	fmt.Printf("Signer     => %s\n", signer)
	fmt.Printf("Signature  => %s\n", signature)
	fmt.Printf("Nonce      => %s\n", nonce.String())
	return nil
}

// createEIP712Hash creates the EIP-712 hash for the checkpoint oracle.
func createEIP712Hash(index uint64, oracle common.Address, hash common.Hash, nonce *big.Int, chainID *big.Int) []byte {
	// EIP-712 domain separator type hash
	domainSeparatorTypeHash := crypto.Keccak256([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))

	// Checkpoint type hash
	checkpointTypeHash := crypto.Keccak256([]byte("Checkpoint(uint64 sectionIndex,bytes32 hash,uint256 nonce)"))

	// Domain name and version
	domainName := crypto.Keccak256([]byte("CheckpointOracle"))
	domainVersion := crypto.Keccak256([]byte("1"))

	// Create domain separator
	domainSeparator := crypto.Keccak256(abiEncode(
		domainSeparatorTypeHash,
		domainName,
		domainVersion,
		chainID,
		oracle,
	))

	// Create checkpoint struct hash
	checkpointHash := crypto.Keccak256(abiEncode(
		checkpointTypeHash,
		big.NewInt(int64(index)),
		hash,
		nonce,
	))

	// Create final signed hash (EIP-712 format)
	signedHash := crypto.Keccak256(append(
		[]byte{0x19, 0x01},
		append(domainSeparator, checkpointHash...)...,
	))

	return signedHash
}

// abiEncode is a simplified ABI encoder for our specific use case
func abiEncode(types ...interface{}) []byte {
	var result []byte
	for _, t := range types {
		switch v := t.(type) {
		case []byte:
			result = append(result, v...)
		case common.Address:
			result = append(result, v.Bytes()...)
		case common.Hash:
			result = append(result, v.Bytes()...)
		case *big.Int:
			result = append(result, v.Bytes()...)
		default:
			panic(fmt.Sprintf("unsupported type: %T", t))
		}
	}
	return result
}

// ecrecover calculates the sender address from a sighash and signature combo.
func ecrecover(sighash []byte, sig []byte) common.Address {
	sig[64] -= 27
	defer func() { sig[64] += 27 }()

	signer, err := crypto.SigToPub(sighash, sig)
	if err != nil {
		utils.Fatalf("Failed to recover sender from signature %x: %v", sig, err)
	}
	return crypto.PubkeyToAddress(*signer)
}

// publish registers the specified checkpoint which generated by connected node
// with a authorised private key.
func publish(ctx *cli.Context) error {
	// Print the checkpoint oracle's current status to make sure we're interacting
	// with the correct network and contract.
	status(ctx)

	// Gather the signatures from the CLI
	var sigs [][]byte
	for _, sig := range strings.Split(ctx.String(signaturesFlag.Name), ",") {
		trimmed := strings.TrimPrefix(strings.TrimSpace(sig), "0x")
		if len(trimmed) != 130 {
			utils.Fatalf("Invalid signature in --signature: '%s'", trimmed)
		} else {
			sigs = append(sigs, common.Hex2Bytes(trimmed))
		}
	}
	// Retrieve the checkpoint we want to sign to sort the signatures
	var (
		client       = newRPCClient(ctx.GlobalString(nodeURLFlag.Name))
		addr, oracle = newContract(client)
		checkpoint   = getCheckpoint(ctx, client)
	)

	// Get chain ID for EIP-712 hash creation
	chainID, err := oracle.GetChainId(nil)
	if err != nil {
		utils.Fatalf("Failed to get chain ID: %v", err)
	}

	// Fetch nonces for each signer
	var nonces []*big.Int
	for _, sig := range sigs {
		// Create a temporary hash for signature recovery (using nonce 0 as placeholder)
		tempHash := createEIP712Hash(checkpoint.SectionIndex, addr, checkpoint.Hash(), big.NewInt(0), chainID)
		signer := ecrecover(tempHash, sig)
		nonce, err := oracle.GetAdminNonce(nil, signer)
		if err != nil {
			utils.Fatalf("Failed to get nonce for signer %s: %v", signer.Hex(), err)
		}
		nonces = append(nonces, nonce)
	}

	// Sort signatures and nonces by signer address
	for i := 0; i < len(sigs); i++ {
		for j := i + 1; j < len(sigs); j++ {
			// Create temporary hashes for comparison (using nonce 0 as placeholder)
			tempHash := createEIP712Hash(checkpoint.SectionIndex, addr, checkpoint.Hash(), big.NewInt(0), chainID)
			signerA := ecrecover(tempHash, sigs[i])
			signerB := ecrecover(tempHash, sigs[j])
			if bytes.Compare(signerA.Bytes(), signerB.Bytes()) > 0 {
				sigs[i], sigs[j] = sigs[j], sigs[i]
				nonces[i], nonces[j] = nonces[j], nonces[i]
			}
		}
	}

	// Convert signatures to v, r, s arrays
	var v []uint8
	var r [][32]byte
	var s [][32]byte
	for _, sig := range sigs {
		if len(sig) != 65 {
			utils.Fatalf("Invalid signature length")
		}
		v = append(v, sig[64])
		r = append(r, common.BytesToHash(sig[:32]))
		s = append(s, common.BytesToHash(sig[32:64]))
	}

	// Print a summary of the operation that's going to be performed
	fmt.Printf("Publishing %d => %s:\n\n", checkpoint.SectionIndex, checkpoint.Hash().Hex())
	for _, sig := range sigs {
		tempHash := createEIP712Hash(checkpoint.SectionIndex, addr, checkpoint.Hash(), big.NewInt(0), chainID)
		fmt.Printf("Signer => %s\n", ecrecover(tempHash, sig).Hex())
	}
	fmt.Println()

	// Publish the checkpoint into the oracle using EIP-712 with nonces
	fmt.Println("Sending publish request to Clef...")
	tx, err := oracle.Contract().SetCheckpoint(newClefSigner(ctx), checkpoint.Hash(), checkpoint.SectionIndex, nonces, v, r, s)
	if err != nil {
		utils.Fatalf("Register contract failed %v", err)
	}
	log.Info("Successfully registered checkpoint", "tx", tx.Hash().Hex())
	return nil
}
