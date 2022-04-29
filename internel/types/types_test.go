package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestTypes(t *testing.T) {
	tx1 := &TransferNativeTx{
		AssetContract: common.BigToAddress(big.NewInt(7489474889940949889)),
		RegisterEmail: "0xjdjdjdj",
		Nonce:         "1233",
		To:            common.BigToAddress(big.NewInt(7489474889940949889)),
		Amount:        "2344",
		FeeToken:      common.BigToAddress(big.NewInt(7489474889940949889)),
		FeeAmount:     "6438984903003",
		Sig:           "3848457",
		SigKeyIndex:   "4585985859",
	}
	a, _ := json.Marshal(tx1)
	fmt.Println(string(a))
	fmt.Println("------")
	tx2 := &TransferNativeTx{}
	json.Unmarshal(a, &tx2)

	fmt.Println(tx2)
	t.Error("a")
}
