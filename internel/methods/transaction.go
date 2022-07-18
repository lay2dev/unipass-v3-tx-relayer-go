package methods

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionList struct {
	sync.Mutex
	fromAddress  common.Address
	currentNonce uint64
	minNumber    int64
	maxNumber    int64
	txs          map[int64]*types.Transaction
}

func NewTransactionList(fromAddress common.Address, currentNonce uint64) *TransactionList {
	return &TransactionList{
		fromAddress:  fromAddress,
		currentNonce: currentNonce,
		minNumber:    0,
		maxNumber:    0,
		txs:          make(map[int64]*types.Transaction),
	}
}

func (txList *TransactionList) getNonceAndAdd() (uint64, error) {
	res := txList.currentNonce
	txList.currentNonce++
	return res, nil
}

func (txList *TransactionList) AddTx(txs ...*types.Transaction) {
	for _, tx := range txs {
		txList.txs[txList.maxNumber] = tx
		txList.maxNumber++
	}
}

func (txList *TransactionList) GetTx(count int64) []*types.Transaction {
	txList.Lock()
	defer txList.Unlock()
	if txList.maxNumber == txList.minNumber {
		return nil
	}

	println("len: ", count, txList.maxNumber, txList.minNumber)
	if txList.maxNumber-txList.minNumber < count {
		count = txList.maxNumber - txList.minNumber
	}

	res := make([]*types.Transaction, 0, count)
	for i := int64(0); i < count; i++ {
		res = append(res, txList.txs[txList.minNumber+i])
	}

	return res
}

func (txList *TransactionList) FinishTx(count int64) {
	txList.Lock()
	defer txList.Unlock()
	for i := int64(0); i < count; i++ {
		delete(txList.txs, txList.minNumber+i)
	}

	txList.minNumber += count
	if txList.maxNumber == txList.minNumber {
		txList.maxNumber = 0
		txList.minNumber = 0
	}
}
