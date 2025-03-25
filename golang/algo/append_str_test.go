package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_appendStr(t *testing.T) {
	a := appendStr()
	assert.Equal(t, "Hello World", a("World"))
	assert.Equal(t, "Hello World !", a("!"))

	b := appendStr()
	assert.Equal(t, "Hello World", b("World"))
	assert.Equal(t, "Hello World Gopher", b("Gopher"))
}

func appendStr() func(a string) string {
	t := "Hello"

	f := func(b string) string {
		t = t + " " + b
		return t
	}

	return f
}
