package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func intToHex(i int64) string {
	return fmt.Sprintf("%x", i)
}

func hexToInt(src string) (int64, error) {
	n, err := strconv.ParseInt(src, 16, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func intToBase58(i int64) string {
	var dst []byte
	b := EncodeBig(dst, big.NewInt(i))
	return string(b)
}

func base58ToInt(src string) (int64, error) {
	i, err := DecodeToBig([]byte(src))
	if err != nil {
		return 0, err
	}
	return i.Int64(), nil
}
