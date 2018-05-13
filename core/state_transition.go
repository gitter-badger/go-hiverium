// Copyright 2014 The happyuc-go Authors
// This file is part of the happyuc-go library.
//
// The happyuc-go library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The happyuc-go library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the happyuc-go library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"errors"
	"math/big"

	"github.com/happyuc-project/happyuc-go/common"
	"github.com/happyuc-project/happyuc-go/core/vm"
	"github.com/happyuc-project/happyuc-go/log"
	"github.com/happyuc-project/happyuc-go/params"
	"math"
)

var (
	errInsufficientBalanceForGas = errors.New("insufficient balance to pay for gas")
)

/*
The State Transitioning Model

A state transition is a change made when a transaction is applied to the current world state
The state transitioning model does all all the necessary work to work out a valid new state root.

1) Nonce handling
2) Pre pay gas
3) Create a new state object if the recipient is \0*32
4) Value transfer
== If contract creation ==
  4a) Attempt to run transaction data
  4b) If valid, use result as code for the new state object
== end ==
5) Run Script section
6) Derive new state root
*/
type StateTransition struct {
	gp         *GasPool
	msg        Message
	gas        uint64
	gasPrice   *big.Int
	initialGas uint64
	value      *big.Int
	actValue   *big.Int
	data       []byte
	state      vm.StateDB
	evm        *vm.EVM
}

// Message represents a message sent to a contract.
type Message interface {
	From() common.Address
	// FromFrontier() (common.Address, error)
	To() *common.Address

	GasPrice() *big.Int
	Gas() uint64
	Value() *big.Int

	Nonce() uint64
	CheckNonce() bool
	Data() []byte
}

// IntrinsicGas computes the 'intrinsic gas' for a message with the given data.
func IntrinsicGas(data []byte, contractCreation, homestead bool) (uint64, error) {
	// Set the starting gas for the raw transaction
	var gas uint64
	if contractCreation && homestead {
		gas = params.TxGasContractCreation
	} else {
		gas = params.TxGas
	}
	// Bump the required gas by the amount of transactional data
	if len(data) > 0 {
		// Zero and non-zero bytes are priced differently
		var nz uint64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}
		// Make sure we don't exceed uint64 for all data combinations
		if (math.MaxUint64-gas)/params.TxDataNonZeroGas < nz {
			return 0, vm.ErrOutOfGas
		}
		gas += nz * params.TxDataNonZeroGas

		z := uint64(len(data)) - nz
		if (math.MaxUint64-gas)/params.TxDataZeroGas < z {
			return 0, vm.ErrOutOfGas
		}
		gas += z * params.TxDataZeroGas
	}
	return gas, nil
}

// NewStateTransition initialises and returns a new state transition object.
func NewStateTransition(evm *vm.EVM, msg Message, gp *GasPool) *StateTransition {
	return &StateTransition{
		gp:       gp,
		evm:      evm,
		msg:      msg,
		gasPrice: msg.GasPrice(),
		value:    msg.Value(),
		data:     msg.Data(),
		state:    evm.StateDB,
	}
}

// ApplyMessage computes the new state by applying the given message
// against the old state within the environment.
//
// ApplyMessage returns the bytes returned by any EVM execution (if it took place),
// the gas used (which includes gas refunds) and an error if it failed. An error always
// indicates a core error meaning that the message would always fail for that particular
// state and would never be accepted within a block.
func ApplyMessage(evm *vm.EVM, msg Message, gp *GasPool) ([]byte, uint64, bool, error) {
	return NewStateTransition(evm, msg, gp).TransitionDb()
}

func (st *StateTransition) from() vm.AccountRef {
	f := st.msg.From()
	if !st.state.Exist(f) {
		st.state.CreateAccount(f)
	}
	return vm.AccountRef(f)
}

func (st *StateTransition) to() vm.AccountRef {
	if st.msg == nil {
		return vm.AccountRef{}
	}
	to := st.msg.To()
	if to == nil {
		return vm.AccountRef{} // contract creation
	}

	reference := vm.AccountRef(*to)
	if !st.state.Exist(*to) {
		st.state.CreateAccount(*to)
	}
	return reference
}

func (st *StateTransition) useGas(amount uint64) error {
	if st.gas < amount {
		return vm.ErrOutOfGas
	}
	st.gas -= amount

	return nil
}

// Transactions fee will be deducted from the recipient. Consider the recipient may
// not have hucer balance, fee will deducted from this transfer.
func (st *StateTransition) buyGas() error {
	var (
		state = st.state
		from  = st.from().Address()
		hucTx = len(st.state.GetCode(st.to().Address())) == 0
	)
	mgval := new(big.Int).Mul(new(big.Int).SetUint64(st.msg.Gas()), st.gasPrice)
	if state.GetBalance(from).Cmp(mgval) == -1 && st.value.Cmp(mgval) <= 0 && hucTx {
		return errInsufficientBalanceForGas
	}
	if err := st.gp.SubGas(st.msg.Gas()); err != nil {
		return err
	}
	st.gas += st.msg.Gas()
	st.initialGas = st.msg.Gas()

	if hucTx {
		state.SubBalance(from, mgval)
		if st.msg.To() == nil {
			st.actValue = st.value
		} else {
			st.actValue = new(big.Int).Sub(st.value, mgval)
		}
	}
	return nil
}

func (st *StateTransition) preCheck() error {
	msg := st.msg
	sender := st.from()

	// Make sure this transaction's nonce is correct
	if msg.CheckNonce() {
		nonce := st.state.GetNonce(sender.Address())
		if nonce < msg.Nonce() {
			return ErrNonceTooHigh
		} else if nonce > msg.Nonce() {
			return ErrNonceTooLow
		}
	}
	return st.buyGas()
}

// TransitionDb will transition the state by applying the current message and
// returning the result including the the used gas. It returns an error if it
// failed. An error indicates a consensus issue.
func (st *StateTransition) TransitionDb() (ret []byte, usedGas uint64, failed bool, err error) {
	// pre check and pay deposit
	if err = st.preCheck(); err != nil {
		return nil, 0, false, err
	}

	// Pay intrinsic gas
	homestead := st.evm.ChainConfig().IsHomestead(st.evm.BlockNumber)
	contractCreation := st.msg.To() == nil
	if gas, err := IntrinsicGas(st.data, contractCreation, homestead); err != nil {
		return nil, 0, false, err
	} else if err = st.useGas(gas); err != nil {
		return nil, 0, false, err
	}

	// do transaction
	ret, failed, err = st.transitionDb()
	if err == vm.ErrInsufficientBalance {
		return nil, 0, false, err
	}

	// refund deposit
	st.refundGas()
	if len(st.state.GetCode(st.to().Address())) == 0 {
		// If just normal tx, like transfer, contract creation etc...
		// We need to pay processing fees to miners.
		st.state.AddBalance(st.evm.Coinbase, new(big.Int).Mul(new(big.Int).SetUint64(st.gasUsed()), st.gasPrice))
	}

	return ret, st.gasUsed(), failed, err
}

// vm errors do not effect consensus and are therefor not
// assigned to err, except for insufficient balance error.
func (st *StateTransition) transitionDb() (ret []byte, failed bool, err error) {
	sender := st.from()
	if st.msg.To() == nil {
		ret, _, st.gas, err = st.evm.Create(sender, st.data, st.gas, st.actValue)
	} else {
		// Increment the nonce for the next transaction
		st.state.SetNonce(sender.Address(), st.state.GetNonce(sender.Address())+1)
		ret, st.gas, err = st.evm.Call(sender, st.to().Address(), st.data, st.gas, st.actValue)
	}
	if err != nil {
		log.Debug("VM returned with error", "err", err)
		if err == vm.ErrInsufficientBalance {
			return nil, false, err
		}
	}
	return ret, err != nil, err
}

func (st *StateTransition) refundGas() {
	// Apply refund counter, capped to half of the used gas.
	refund := st.gasUsed() / 2
	if refund > st.state.GetRefund() {
		refund = st.state.GetRefund()
	}
	st.gas += refund

	// Return HUC for remaining gas, exchanged at the original rate.
	if len(st.state.GetCode(st.to().Address())) == 0 {
		// If just normal tx, like transfer, contract creation etc...
		remaining := new(big.Int).Mul(new(big.Int).SetUint64(st.gas), st.gasPrice)
		var refundee common.Address
		if st.msg.To() == nil {
			refundee = st.from().Address()
		} else {
			refundee = st.to().Address()
		}
		st.state.AddBalance(refundee, remaining)
	}

	// Also return remaining gas to the block gas counter so it is
	// available for the next transaction.
	st.gp.AddGas(st.gas)
}

// gasUsed returns the amount of gas used up by the state transition.
func (st *StateTransition) gasUsed() uint64 {
	return st.initialGas - st.gas
}
