package main

import "runtime"

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("caller not found")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	println("enter: ", name)
	return func() { println("exit:  ", name) }
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

func main() {
	defer Trace()()
	foo()
}
