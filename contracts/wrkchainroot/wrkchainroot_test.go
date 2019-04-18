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
	masterRegistrarKey, _ = crypto.HexToECDSA("02c10f620820398d7df4c6157e48f86d65456346ff5633483a9e2384de732795")
	masterRegistrarAddr   = crypto.PubkeyToAddress(masterRegistrarKey.PublicKey)
	wrkchainKey, _ = crypto.HexToECDSA("8ad09b7c563340f3216650b3007f87723b2b2f3fd4c81f31aa91801bf31d404b")
	wrkchainAddr = crypto.PubkeyToAddress(wrkchainKey.PublicKey)
	accKey1, _ = crypto.HexToECDSA("d7c15922c5371f9536f8b1cef750d5451b01f57b2865278fef6bd65c9a17a3a3")
	accAddr1 = crypto.PubkeyToAddress(accKey1.PublicKey)
	accKey2, _ = crypto.HexToECDSA("e6d90f41cebb11eb6f5f5d7b71a49f3f4309d78ffb3e66eaa6c7af94c7d99d41")
	accAddr2 = crypto.PubkeyToAddress(accKey2.PublicKey)
	accKey3, _ = crypto.HexToECDSA("dc23751c35323b70ad70ec2d08d9be59b74f8491baccc87123a0c0ffdfa24936")
	accAddr3 = crypto.PubkeyToAddress(accKey3.PublicKey)
	wrkchainId1 = new(big.Int).SetUint64(12345)
	wrkchainAuthAddresses1 = []common.Address{accAddr1}
	wrkchainGenesisHash1 = [32]byte{'x'}
	genesisBalance = big.NewInt(1000000000000000000)
)

func TestWrkchainRootDeploy(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		masterRegistrarAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOpts := bind.NewKeyedTransactor(masterRegistrarKey)

	ownerBalance, _ := contractBackend.BalanceAt(context.Background(), masterRegistrarAddr, big.NewInt(0))

	t.Log("master registrar", masterRegistrarAddr.Hex())
	t.Log("master registrar balance", ownerBalance)

	wrkchainRootAddress, _, err := DeployWrkchainRoot(transactOpts, contractBackend, masterRegistrarAddr)
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

func TestAddRemoveRegistrar(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		masterRegistrarAddr: {
			Balance: genesisBalance,
		},
		accAddr1: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOptsMaster := bind.NewKeyedTransactor(masterRegistrarKey)
	transactOptsAcc1 := bind.NewKeyedTransactor(accKey1)

	wrkchainRootAddress, wrkchainRoot, err := DeployWrkchainRoot(transactOptsMaster, contractBackend, masterRegistrarAddr)
	if err != nil {
		t.Fatalf("can't deploy wrkchainroot: %v", err)
	}
	contractBackend.Commit()


	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	f := func(key, val common.Hash) bool {
		t.Log(key.Hex(), val.Hex())
		return true
	}
	contractBackend.ForEachStorageAt(ctx, wrkchainRootAddress, nil, f)

	masterRegistrar, err := wrkchainRoot.GetMasterRegistrar()

	t.Logf("masterRegistrar: %v", masterRegistrar.Hex())

	// Master Registrar adds accAddr1
	tx, err := wrkchainRoot.AddRegistrar(accAddr1)

	if err != nil {
		t.Fatalf("can't add registrar: %v", err)
	}
	contractBackend.Commit()
	t.Log("AddRegistrar Tx", tx)

	contractBackend.ForEachStorageAt(ctx, wrkchainRootAddress, nil, f)

	// accAddr1 adds accAddr2 as registrar
	acc1Registrar, _ := NewWrkchainRoot(transactOptsAcc1, wrkchainRootAddress, contractBackend)

	tx1, err1 := acc1Registrar.AddRegistrar(accAddr2)

	if err1 != nil {
		t.Fatalf("can't add registrar: %v", err1)
	}
	contractBackend.Commit()
	t.Log("AddRegistrar Tx", tx1)

	// Master removes accAddr1
	tx2, err2 := wrkchainRoot.RemoveRegistrar(accAddr1)

	if err2 != nil {
		t.Fatalf("can't remove registrar: %v", err2)
	}
	contractBackend.Commit()
	t.Log("RemoveRegistrar Tx", tx2)

}

func TestUnauthAddRegistrar(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		masterRegistrarAddr: {
			Balance: genesisBalance,
		},
		accAddr1: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOptsMaster := bind.NewKeyedTransactor(masterRegistrarKey)
	transactOptsAcc1 := bind.NewKeyedTransactor(accKey1)

	wrkchainRootAddress, _, err := DeployWrkchainRoot(transactOptsMaster, contractBackend, masterRegistrarAddr)
	if err != nil {
		t.Fatalf("can't deploy wrkchainroot: %v", err)
	}
	contractBackend.Commit()

	acc1Registrar, _ := NewWrkchainRoot(transactOptsAcc1, wrkchainRootAddress, contractBackend)

	// unauthorised accAddr1 tries to add accAddr3
	_, err1 := acc1Registrar.AddRegistrar(accAddr3)

	contractBackend.Commit()
	if err1 == nil {
		t.Fatalf("accAddr1 was able to add accAddr3: %v", err1)
	}
}

func TestRegisterWrkchain(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		masterRegistrarAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOptsMaster := bind.NewKeyedTransactor(masterRegistrarKey)

	_, wrkchainRoot, err := DeployWrkchainRoot(transactOptsMaster, contractBackend, masterRegistrarAddr)
	if err != nil {
		t.Fatalf("can't deploy wrkchainroot: %v", err)
	}
	contractBackend.Commit()

	// Master Registrar registers a WRKChain
	tx, err := wrkchainRoot.RegisterWrkChain(wrkchainId1, wrkchainAuthAddresses1, wrkchainGenesisHash1)
	contractBackend.Commit()
	if err != nil {
		t.Fatalf("can't register wrkchain: %v", err)
	}

	t.Log("RegisterWrkChain Tx", tx)

}

func TestRecordHeader(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		masterRegistrarAddr: {
			Balance: genesisBalance,
		},
		accAddr1: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc)

	transactOptsMaster := bind.NewKeyedTransactor(masterRegistrarKey)
	transactOptsAcc1 := bind.NewKeyedTransactor(accKey1)

	wrkchainRootAddress, wrkchainRoot, _ := DeployWrkchainRoot(transactOptsMaster, contractBackend, masterRegistrarAddr)
	contractBackend.Commit()
	// Master Registrar registers a WRKChain
	tx, _ := wrkchainRoot.RegisterWrkChain(wrkchainId1, wrkchainAuthAddresses1, wrkchainGenesisHash1)
	contractBackend.Commit()
	t.Log("RegisterWrkChain Tx", tx)

	acc1RecordHeader, _ := NewWrkchainRoot(transactOptsAcc1, wrkchainRootAddress, contractBackend)

	tx, err := acc1RecordHeader.RecordHeader(
		wrkchainId1,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
		accAddr1,
	)

	if err != nil {
		t.Fatalf("can't record wrkchain header: %v", err)
	}

	t.Log("RecordHeader Tx", tx)

}
