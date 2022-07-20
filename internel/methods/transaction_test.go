package methods

import (
	"fmt"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestTxListConcurrency(t *testing.T) {
	a := NewTransactionList(common.Address{}, 0)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		n := i
		go func() {
			a.Lock()
			defer a.Unlock()
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("panic err:", err)
				}
			}()
			println(n)
			if n%17 == 0 && n != 0 {
				panic(n)
			}
		}()
	}

	wg.Wait()
}
