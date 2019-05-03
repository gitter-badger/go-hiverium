// Copyright (c) 2019 Unification Foundation
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package wrkchainroot

import (
	"context"
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/accounts/abi/bind/backends"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core"
	"github.com/unification-com/mainchain/crypto"
	"math/big"
	"testing"
	"time"
)

var (
	deployerKey, _         = crypto.HexToECDSA("e7b1f342c77881ff425d09b83892cc0683435878cf1d3ff2b6e015ada9378ddf")
	deployerAddr           = crypto.PubkeyToAddress(deployerKey.PublicKey)
	wrkchainOwnerKey, _    = crypto.HexToECDSA("8ad09b7c563340f3216650b3007f87723b2b2f3fd4c81f31aa91801bf31d404b")
	wrkchainOwnerAddr      = crypto.PubkeyToAddress(wrkchainOwnerKey.PublicKey)
	accKey1, _             = crypto.HexToECDSA("d7c15922c5371f9536f8b1cef750d5451b01f57b2865278fef6bd65c9a17a3a3")
	accAddr1               = crypto.PubkeyToAddress(accKey1.PublicKey)
	accKey2, _             = crypto.HexToECDSA("e6d90f41cebb11eb6f5f5d7b71a49f3f4309d78ffb3e66eaa6c7af94c7d99d41")
	accAddr2               = crypto.PubkeyToAddress(accKey2.PublicKey)
	accKey3, _             = crypto.HexToECDSA("dc23751c35323b70ad70ec2d08d9be59b74f8491baccc87123a0c0ffdfa24936")
	accAddr3               = crypto.PubkeyToAddress(accKey3.PublicKey)
	wrkchainId1            = new(big.Int).SetUint64(12345)
	wrkchainAuthAddresses1 = []common.Address{accAddr1, wrkchainOwnerAddr}
	wrkchainGenesisHash1   = [32]byte{'x'}
	genesisBalance         = new(big.Int).SetUint64(1000000000000000000)
	depositAmt             = new(big.Int).SetUint64(50000000000000000)
	minBlocksInt           = uint64(5)
	minBlocks              = new(big.Int).SetUint64(minBlocksInt)
)

func TestWrkchainRootDeploy(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOpts := bind.NewKeyedTransactor(deployerKey)

	ownerBalance, _ := contractBackend.BalanceAt(context.Background(), deployerAddr, big.NewInt(0))

	t.Log("Deployer", deployerAddr.Hex())
	t.Log("Deployer balance", ownerBalance)

	wrkchainRootAddress, _, err := DeployWrkchainRoot(transactOpts, contractBackend, depositAmt, minBlocks)
	if err != nil {
		t.Fatalf("can't deploy wrkchainroot: %v", err)
	}
	contractBackend.Commit()
	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	code, err := contractBackend.CodeAt(ctx, wrkchainRootAddress, nil)
	if err != nil {
		t.Fatalf("can't get wrkchainroot code: %v", err)
	}

	t.Log("contract code", common.ToHex(code))
	t.Log("contract address", wrkchainRootAddress.Hex())
}

func TestRegisterWrkchain(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)

	wrkchainRootAddress, _, err  := DeployWrkchainRoot(transactOptsDeployer, contractBackend, depositAmt, minBlocks)
	if err != nil {
		t.Fatalf("can't deploy wrkchainroot: %v", err)
	}
	contractBackend.Commit()

	// wrkchain owner registers a WRKChain
	transactOptsWrkchainOwner.Value = depositAmt
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	tx, err := wrkchainOwner.RegisterWrkChain(wrkchainId1, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	if err != nil {
		t.Fatalf("can't register wrkchainroot: %v", err)
	}
	contractBackend.Commit()

	t.Log("RegisterWrkChain Tx", tx)

	registerredHash, _ := wrkchainOwner.GetGenesis(wrkchainId1)
	t.Log("Registered Genesis Hash", registerredHash)
	wrkchainList, _ := wrkchainOwner.WrkchainList(wrkchainId1)

	t.Log(wrkchainList)

}

func TestRegisterWrkchainWrongDeposit(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)

	wrkchainRootAddress, _, _  := DeployWrkchainRoot(transactOptsDeployer, contractBackend, depositAmt, minBlocks)

	contractBackend.Commit()

	// wrkchain owner tries to register a WRKChain without sending deposit
	transactOptsWrkchainOwner.Value = new(big.Int).SetUint64(0)
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, err := wrkchainOwner.RegisterWrkChain(wrkchainId1, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	if err == nil {
		t.Fatal("should fail to register wrkchainroot - no deposit sent")
	}

	// wrkchain owner tries to register a WRKChain with too small deposit
	transactOptsWrkchainOwner.Value = new(big.Int).SetUint64(3)
	wrkchainOwner, _ = NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, err = wrkchainOwner.RegisterWrkChain(wrkchainId1, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	if err == nil {
		t.Fatal("should fail to register wrkchainroot - no deposit sent")
	}

	// wrkchain owner tries to register a WRKChain with too much deposit
	transactOptsWrkchainOwner.Value = new(big.Int).SetUint64(9)
	wrkchainOwner, _ = NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, err = wrkchainOwner.RegisterWrkChain(wrkchainId1, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	if err == nil {
		t.Fatal("should fail to register wrkchainroot - no deposit sent")
	}
}

func TestWritingHashes(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)

	wrkchainRootAddress, _, _  := DeployWrkchainRoot(transactOptsDeployer, contractBackend, depositAmt, minBlocks)

	contractBackend.Commit()

	// wrkchain owner registers a WRKChain
	transactOptsWrkchainOwner.Value = depositAmt
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, _ = wrkchainOwner.RegisterWrkChain(wrkchainId1, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	contractBackend.Commit()

	//reset Value to 0, since it's not calling a payable function
	transactOptsWrkchainOwner.Value = new(big.Int).SetUint64(0)
	wrkchainOwner, _ = NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)
	tx, err := wrkchainOwner.RecordHeader(
		wrkchainId1,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
		wrkchainOwnerAddr,
		)

	if err != nil {
		t.Fatalf("can't record header: %v", err)
	}
	contractBackend.Commit()

	t.Log("RecordHeader Tx", tx)

}

func TestOnlyAuthCanWriteHashes(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
		accAddr2: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)
	transactOptsAcc2 := bind.NewKeyedTransactor(accKey2)

	wrkchainRootAddress, _, _  := DeployWrkchainRoot(transactOptsDeployer, contractBackend, depositAmt, minBlocks)

	contractBackend.Commit()

	// wrkchain owner registers a WRKChain
	transactOptsWrkchainOwner.Value = depositAmt
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, _ = wrkchainOwner.RegisterWrkChain(wrkchainId1, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	contractBackend.Commit()

	acc2, _ := NewWrkchainRoot(transactOptsAcc2, wrkchainRootAddress, contractBackend)

	_, err := acc2.RecordHeader(
		wrkchainId1,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
		wrkchainOwnerAddr,
	)

	if err == nil {
		t.Fatalf("Should have failed...")
	}
}

func TestNewAuthCanWrite(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
		accAddr2: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)
	transactOptsAcc2 := bind.NewKeyedTransactor(accKey2)

	wrkchainRootAddress, _, _  := DeployWrkchainRoot(transactOptsDeployer, contractBackend, depositAmt, minBlocks)

	contractBackend.Commit()

	// wrkchain owner registers a WRKChain
	transactOptsWrkchainOwner.Value = depositAmt
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, _ = wrkchainOwner.RegisterWrkChain(wrkchainId1, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	contractBackend.Commit()

	//WRKChain owner adds new Auth
	transactOptsWrkchainOwner.Value = new(big.Int).SetUint64(0)
	wrkchainOwner, _ = NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)
	_, err := wrkchainOwner.AddAuthAddresses(wrkchainId1, []common.Address{accAddr2})

	if err != nil {
		t.Fatalf("Failed to authorise new address: %v", err)
	}
	contractBackend.Commit()

	acc2, _ := NewWrkchainRoot(transactOptsAcc2, wrkchainRootAddress, contractBackend)

	tx, err := acc2.RecordHeader(
		wrkchainId1,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
		wrkchainOwnerAddr,
	)
	if err != nil {
		t.Fatalf("New authorised account failed to write hashes: %v", err)
	}

	t.Logf("Tx: %v", tx)

}

func TestDepositRefund(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)

	wrkchainRootAddress, _, _  := DeployWrkchainRoot(transactOptsDeployer, contractBackend, depositAmt, minBlocks)

	contractBackend.Commit()

	// wrkchain owner registers a WRKChain
	transactOptsWrkchainOwner.Value = depositAmt
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, _ = wrkchainOwner.RegisterWrkChain(wrkchainId1, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	contractBackend.Commit()

	//reset Value to 0, since it's not calling a payable function
	transactOptsWrkchainOwner.Value = new(big.Int).SetUint64(0)
	wrkchainOwner, _ = NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	for i := uint64(1); i <= minBlocksInt +1; i++ {

		wrkchainOwnerBalanceBefore, _ := contractBackend.BalanceAt(context.Background(), wrkchainOwnerAddr, nil)

		_, err := wrkchainOwner.RecordHeader(
			wrkchainId1,
			new(big.Int).SetUint64(i),
			[32]byte{'a'},
			[32]byte{'b'},
			[32]byte{'c'},
			[32]byte{'d'},
			[32]byte{'e'},
			wrkchainOwnerAddr,
		)

		if err != nil {
			t.Fatalf("can't record header: %v", err)
		}

		contractBackend.Commit()

		t.Log("Sent hashes for block #", i)

		wrkchainOwnerBalanceAfter, _ := contractBackend.BalanceAt(context.Background(), wrkchainOwnerAddr, nil)

		numBlocksSubmitted, _ := wrkchainOwner.GetNumBlocksSubmitted(wrkchainId1)
		lastBlock, _ := wrkchainOwner.GetNumBlocksSubmitted(wrkchainId1)

		balanceDiff := new(big.Int).Sub(wrkchainOwnerBalanceAfter, wrkchainOwnerBalanceBefore)
		t.Log("balanceDiff", balanceDiff)

		// Check refund has been sent. Balance difference before and after Tx should be > 0
		// if refunded successfully.
		if numBlocksSubmitted.Uint64() == minBlocksInt {
             if balanceDiff.Uint64() <= 0 {
           	     t.Fatalf("Deposit not returned after block #%v", minBlocksInt)
		     }
		}

		t.Log("numBlocksSubmitted", numBlocksSubmitted)
		t.Log("lastBlock", lastBlock)
		t.Log("wrkchainOwnerBalanceAfter", wrkchainOwnerBalanceAfter)
	}
}
