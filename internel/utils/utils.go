package utils

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
)

func HexToBytes(s string) ([]byte, error) {
	res, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func HexToBytes32(s string) ([32]byte, error) {
	res, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
	if err != nil {
		return [32]byte{0}, err
	}

	if len(res) != 32 {
		return [32]byte{0}, errors.New("params error")
	}
	var buf [32]byte
	copy(buf[:], res)
	return buf, nil
}

func BigIntDecode(str string) (*big.Int, error) {
	res, ok := new(big.Int).SetString(str, 0)
	if !ok {
		return nil, errors.New("wrong number")
	}
	return res, nil
}
