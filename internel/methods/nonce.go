package methods

import (
	"context"
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	errTransactionTooMuch = errors.New("transaction too much")
)

type NonceCache struct {
	sync.Mutex
	fromAddress  common.Address
	currentNonce uint64
}

func newNonceCache(fromAddress common.Address, currentNonce uint64) *NonceCache {
	return &NonceCache{
		fromAddress:  fromAddress,
		currentNonce: currentNonce,
	}
}

func (cache *NonceCache) getNonceAndAdd() (uint64, error) {
	cache.Lock()
	defer cache.Unlock()

	res := cache.currentNonce
	cache.currentNonce++
	return res, nil
}

func (cache *NonceCache) flushNonce(client *ethclient.Client) error {
	cache.Lock()
	defer cache.Unlock()
	res, err := client.PendingNonceAt(context.Background(), cache.fromAddress)
	if err != nil {
		return err
	}
	cache.currentNonce = res
	return nil
}
