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
	ownerKey, _ = crypto.HexToECDSA("02c10f620820398d7df4c6157e48f86d65456346ff5633483a9e2384de732795")
	ownerAddr   = crypto.PubkeyToAddress(ownerKey.PublicKey)
)

func TestWrkchainRootDeploy(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{ownerAddr: {Balance: big.NewInt(999999999999999999)}})
	transactOpts := bind.NewKeyedTransactor(ownerKey)

	ownerBalance, _ := contractBackend.BalanceAt(context.Background(), ownerAddr, big.NewInt(0))

	t.Log("master registrar", ownerAddr.Hex())
	t.Log("master registrar balance", ownerBalance)

	wrkchainRootAddress, _, err := DeployWrkchainRoot(transactOpts, contractBackend, ownerAddr)
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
}
