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
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/contracts/wrkchainroot/contract"
)

type WrkchainRoot struct {
	*contract.WRKChainRootSession
	contractBackend bind.ContractBackend
}

func NewWrkchainRoot(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*WrkchainRoot, error) {
	wrkchainRoot, err := contract.NewWRKChainRoot(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &WrkchainRoot{
		&contract.WRKChainRootSession{
			Contract:     wrkchainRoot,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}


func DeployWrkchainRoot(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, ownerAddress common.Address) (common.Address, *WrkchainRoot, error) {
	wrkchainRootAddr, _, _, err := contract.DeployWRKChainRoot(transactOpts, contractBackend, ownerAddress)
	if err != nil {
		return wrkchainRootAddr, nil, err
	}

	wrkchainRoot, err := NewWrkchainRoot(transactOpts, wrkchainRootAddr, contractBackend)
	if err != nil {
		return wrkchainRootAddr, nil, err
	}

	return wrkchainRootAddr, wrkchainRoot, nil
}
