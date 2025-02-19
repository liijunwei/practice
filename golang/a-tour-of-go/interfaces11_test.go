package main

import (
	"testing"

	"golang.org/x/tour/reader"
)

type MyReader struct{}
type MyReaderError bool

func (MyReaderError) Error() string {
	return "capicity not enough"
}

func (r MyReader) Read(p []byte) (int, error) {
	if cap(p) < 1 {
		return 0, MyReaderError(false)
	}

	p[0] = 'A'
	return 1, nil
}

func TestI11(t *testing.T) {
	reader.Validate(MyReader{})
}
