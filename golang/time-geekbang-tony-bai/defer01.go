package main

// https://time.geekbang.org/column/article/471138

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("caller not found")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	gid := curGoroutineID()
	fmt.Printf("g[%05d] enter: [%s]\n", gid, name)
	return func() { fmt.Printf("g[%05d] exit:  [%s]\n", gid, name) }
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func main() {
	defer Trace()()
	foo()
}
