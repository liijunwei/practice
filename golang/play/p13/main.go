package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	h := sha256.New()
	h.Write([]byte("foo"))
	r := h.Sum(nil)
	num := new(big.Int).SetBytes(r)
	fmt.Println(num.String())
}
