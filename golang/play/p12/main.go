package main

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
)

type demo1 struct{}

func (d *demo1) Error() string {
	return "demo1"
}

type demo2 struct{}

func (d demo2) Error() string {
	return "demo2"
}

func main() {
	var f1 error = &demo1{}
	f1 = &demo2{}

	spew.Dump(f1)

	err := demo1Error()
	assert(errors.As(err, f1) == false)

	spew.Dump(f1)
}

func demo1Error() error {
	return &demo1{}
}

func demo2Error() error {
	return demo2{}
}

func assert(ok bool) {
	if !ok {
		panic("assertion failed")
	}
}
