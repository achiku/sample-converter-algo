package main

import (
	"fmt"
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
