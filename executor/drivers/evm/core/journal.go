// Copyright 2016 The go-ethereum Authors
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

package core

import (
	"gitlab.33.cn/chain33/chain33/executor/drivers/evm/vm/common"
	"gitlab.33.cn/chain33/chain33/types"
)

type journalEntry interface {
	undo(mdb *MemoryStateDB)
	getData(mdb *MemoryStateDB) []*types.KeyValue
	getLog(mdb *MemoryStateDB) []*types.ReceiptLog
}

type journal []journalEntry

type (
	baseChange struct {
	}
	// 创建合约对象变更事件
	createAccountChange struct {
		baseChange
		account *common.Address
	}

	// 自杀事件
	suicideChange struct {
		baseChange
		account     *common.Address
		prev        bool // whether account had already suicided
	}

	// nonce变更事件
	nonceChange struct {
		baseChange
		account *common.Address
		prev    uint64
	}

	// 存储状态变更事件
	storageChange struct {
		baseChange
		account       *common.Address
		key, prevalue common.Hash
	}

	// 合约代码状态变更事件
	codeChange struct {
		baseChange
		account            *common.Address
		prevcode, prevhash []byte
	}

	// 返还金额变更事件
	refundChange struct {
		baseChange
		prev uint64
	}

	// 金额变更事件
	balanceChange struct {
		baseChange
		amount int64
		addr string
		data []*types.KeyValue
		logs []*types.ReceiptLog
	}


	addLogChange struct {
		baseChange
		txhash common.Hash
	}

	addPreimageChange struct {
		baseChange
		hash common.Hash
	}
)
func (ch baseChange) undo(s *MemoryStateDB) {
}

func (ch baseChange) getData(s *MemoryStateDB) (kvset []*types.KeyValue) {
	return nil
}

func (ch baseChange) getLog(s *MemoryStateDB) (logs []*types.ReceiptLog) {
	return nil
}


// 创建账户对象的回滚，需要删除缓存中的账户和变更标记
func (ch createAccountChange) undo(s *MemoryStateDB) {
	delete(s.accounts, *ch.account)
	delete(s.accountsDirty, *ch.account)
	delete(s.contractsDirty, *ch.account)
}

// 创建账户对象的数据集
func (ch createAccountChange) getData(s *MemoryStateDB) (kvset []*types.KeyValue) {
	acc := s.accounts[*ch.account]
	if acc != nil {
		kvset = append(kvset, acc.GetDataKV()...)
		kvset = append(kvset, acc.GetStateKV()...)
		return kvset
	}
	return nil
}


func (ch suicideChange) undo(mdb *MemoryStateDB) {
	// 如果已经自杀过了，不处理
	if ch.prev {
		return
	}
	acc := mdb.accounts[*ch.account]
	if acc != nil {
		acc.State.Suicided = ch.prev
	}
}

func (ch suicideChange) getData(mdb *MemoryStateDB) []*types.KeyValue {
	// 如果已经自杀过了，不处理
	if ch.prev {
		return nil
	}
	acc := mdb.accounts[*ch.account]
	if acc != nil {
		return acc.GetStateKV()
	}
	return nil
}

func (ch nonceChange) undo(mdb *MemoryStateDB) {
	acc := mdb.accounts[*ch.account]
	if acc != nil {
		acc.State.Nonce = ch.prev
	}
}

func (ch nonceChange) getData(mdb *MemoryStateDB) []*types.KeyValue {
	acc := mdb.accounts[*ch.account]
	if acc != nil {
		return acc.GetStateKV()
	}
	return nil
}


func (ch codeChange) undo(mdb *MemoryStateDB) {
	acc := mdb.accounts[*ch.account]
	if acc != nil {
		acc.Data.Code = ch.prevcode
		acc.Data.CodeHash = ch.prevhash
	}
}

func (ch codeChange) getData(mdb *MemoryStateDB) (kvset []*types.KeyValue) {
	acc := mdb.accounts[*ch.account]
	if acc != nil {
		kvset = append(kvset, acc.GetDataKV()...)
		kvset = append(kvset, acc.GetStateKV()...)
		return kvset
	}
	return nil
}

func (ch storageChange) undo(mdb *MemoryStateDB) {
	acc := mdb.accounts[*ch.account]
	if acc != nil {
		acc.SetState(ch.key, ch.prevalue)
	}
}

func (ch storageChange) getData(mdb *MemoryStateDB) []*types.KeyValue {
	acc := mdb.accounts[*ch.account]
	if acc != nil {
		return acc.GetStateKV()
	}
	return nil
}

func (ch refundChange) undo(mdb *MemoryStateDB) {
	mdb.refund = ch.prev
}
func (ch refundChange) getData(mdb *MemoryStateDB) []*types.KeyValue {
	return nil
}


func (ch addLogChange) undo(mdb *MemoryStateDB) {
	logs := mdb.logs[ch.txhash]
	if len(logs) == 1 {
		delete(mdb.logs, ch.txhash)
	} else {
		mdb.logs[ch.txhash] = logs[:len(logs)-1]
	}
	mdb.logSize--
}

func (ch addLogChange) getData(mdb *MemoryStateDB) []*types.KeyValue {
	return nil
}


func (ch addPreimageChange) undo(mdb *MemoryStateDB) {
	delete(mdb.preimages, ch.hash)
}
func (ch addPreimageChange) getData(mdb *MemoryStateDB) []*types.KeyValue {
	return nil
}

// 设置成变化钱的金额
func (ch balanceChange) undo(mdb *MemoryStateDB) {
	acc := mdb.CoinsAccount.LoadAccount(ch.addr)
	acc.Balance -= ch.amount
	mdb.CoinsAccount.SaveAccount(acc)
}

func (ch balanceChange) getData(mdb *MemoryStateDB) []*types.KeyValue {
	return ch.data
}
func (ch balanceChange) getLog(mdb *MemoryStateDB) []*types.ReceiptLog {
	return ch.logs
}