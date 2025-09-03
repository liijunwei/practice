package main

import "fmt"

func main() {
	m := newBitmap(10)
	cassert(m.Exists(10) == false)
	m.Set(10)
	cassert(m.Exists(10) == true)
	fmt.Println(m.flags)
}

type bitmap struct {
	flags []byte
}

func newBitmap(max int64) *bitmap {
	flaglen := max/8 + 1
	return &bitmap{
		flags: make([]byte, flaglen),
	}
}

func (m *bitmap) Set(n int) {
	arridx, offset := locate(n)
	m.flags[arridx] |= bitmask(offset)
}

func (m *bitmap) Del(n int) {
	arridx, offset := locate(n)
	m.flags[arridx] &= clearmask(offset)
}

func (m *bitmap) Exists(n int) bool {
	arridx, offset := locate(n)
	res := m.flags[arridx] & bitmask(offset)
	return res != 0
}

func bitmask(offset int) byte {
	cassert(offset >= 0 && offset < 8)
	return 0x1 << offset
}

// inverse bitmask
// complement bitmask
func clearmask(offset int) byte {
	cassert(offset >= 0 && offset < 8)
	return ^bitmask(offset)
}

func locate(n int) (int, int) {
	cassert(n > 0)
	arridx := n / 8
	offset := n % 8
	return arridx, offset
}

func cassert(ok bool) {
	if !ok {
		panic("assersion faied")
	}
}
