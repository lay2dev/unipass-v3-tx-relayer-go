package methods

import (
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionList struct {
	sync.Mutex
	minNumber int64
	maxNumber int64
	txs       map[int64]*types.Transaction
}

func NewTransactionList() *TransactionList {
	return &TransactionList{
		minNumber: 0,
		maxNumber: 0,
		txs:       make(map[int64]*types.Transaction),
	}
}

func (txList *TransactionList) AddTx(tx *types.Transaction) {
	txList.Lock()
	defer txList.Unlock()
	txList.txs[txList.maxNumber] = tx
	txList.maxNumber++
}

func (txList *TransactionList) GetTx() *types.Transaction {
	txList.Lock()
	defer txList.Unlock()
	if txList.maxNumber == txList.minNumber {
		return nil
	}
	return txList.txs[txList.minNumber]
}

func (txList *TransactionList) FinishTx() {
	txList.Lock()
	defer txList.Unlock()
	delete(txList.txs, txList.minNumber)
	txList.minNumber++
	if txList.maxNumber == txList.minNumber {
		txList.maxNumber = 0
		txList.minNumber = 0
	}
}
