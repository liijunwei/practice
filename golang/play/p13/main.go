package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

func toBase62(s string) string {
	var i big.Int
	i.SetBytes([]byte(s))
	return i.Text(62)
}

func toBase62_1(s []byte) string {
	var i big.Int
	i.SetBytes(s)
	return i.Text(62)
}

func main() {
	h := sha256.New()
	h.Write([]byte("foo"))
	r := h.Sum(nil)
	num := new(big.Int).SetBytes(r)
	fmt.Println(num.String())

	fmt.Println([]byte("61"))
	fmt.Println([]byte{61})
	fmt.Println(toBase62("61"))
	fmt.Println(toBase62_1([]byte{0}))
	fmt.Println(toBase62_1([]byte{1}))
	fmt.Println(toBase62_1([]byte{61}))
}
