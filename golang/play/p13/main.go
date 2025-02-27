package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

func toBase62_1(s string) string {
	var i big.Int
	i.SetBytes([]byte(s))
	return i.Text(62)
}

func toBase62_2(s []byte) string {
	var i big.Int
	i.SetBytes(s)
	return i.Text(62)
}

func toBase62_3(s int) string {
	var i big.Int
	i.SetUint64(uint64(s))
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
	fmt.Println(toBase62_1("61"))
	fmt.Println(toBase62_2([]byte{0}))
	fmt.Println(toBase62_2([]byte{1}))
	fmt.Println(toBase62_2([]byte{61}))

	fmt.Println("toBase62_3")
	fmt.Println(toBase62_3(0))
	fmt.Println(toBase62_3(1))
	fmt.Println(toBase62_3(61))
	fmt.Println(toBase62_3(62))
}
