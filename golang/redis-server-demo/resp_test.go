package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Note: https://go.dev/ref/spec#String_literals
// https://stackoverflow.com/questions/46917331/what-is-the-difference-between-backticks-double-quotes-in-golang
//
// There are two forms: raw string literals and interpreted string literals.
// Raw string literals are character sequences between back quotes, as in `foo`.
// Within the quotes, any character may appear except back quote.
//
// a := "\n" // This is one character, a line break.
// b := `\n` // These are two characters, backslash followed by letter n.

func TestParseRespMessage(t *testing.T) {
	// OK what's the difference between "*1\r\n$4\r\nping\r\n" and `*1\r\n$4\r\nping\r\n` ?
	//
	// cuz I found `\r\n` needs 4 reader.ReadByte(), while "\r\n" only needs 2
	// here I need the interpreted string literal actually

	input := "*1\r\n$4\r\nping\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	// arrays type in RESP
	// *<number-of-elements>\r\n<element-1>...<element-n>
	// https://redis.io/docs/latest/develop/reference/protocol-spec/#arrays
	arrType, err := reader.ReadByte()
	require.NoError(t, err)
	require.Equal(t, byte('*'), arrType)

	sizeInByte, err := reader.ReadByte()
	require.NoError(t, err)

	size, err := strconv.Atoi(string(sizeInByte))
	require.NoError(t, err)
	assert.Equal(t, 1, size, string(sizeInByte))

	// consume \r\n in buffered reader
	reader.ReadByte()
	reader.ReadByte()

	// bulk-strings type in RESP
	// $<length>\r\n<data>\r\n
	// https://redis.io/docs/latest/develop/reference/protocol-spec/#bulk-strings
	bulkStrType, err2 := reader.ReadByte()
	require.NoError(t, err2)
	require.Equal(t, byte('$'), bulkStrType, string(bulkStrType))

	lengthInByte, errL := reader.ReadByte()
	require.NoError(t, errL)
	length, err := strconv.Atoi(string(lengthInByte))
	require.NoError(t, err)
	assert.Equal(t, 4, length, string(lengthInByte))

	// consume \r\n in buffered reader
	reader.ReadByte()
	reader.ReadByte()

	buf := make([]byte, length)
	n, err := reader.Read(buf)
	require.NoError(t, err)
	require.Equal(t, length, n)

	assert.Equal(t, "ping", string(buf))
}

// 0x2A in ASCII -> *
// 0x31 in ASCII -> 1
// 0x0D in ASCII -> \r
// 0x0A in ASCII -> \n
// 0x24 in ASCII -> $
// 0x34 in ASCII -> 4
// 0x0D in ASCII -> \r
// 0x0A in ASCII -> \n
// 0x70 in ASCII -> p
// 0x69 in ASCII -> i
// 0x6E in ASCII -> n
// 0x67 in ASCII -> g
// 0x0D in ASCII -> \r
// 0x0A in ASCII -> \n
func TestParsePing(t *testing.T) {
	input := "*1\r\n$4\r\nping\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	var tmp byte

	typ, err := reader.ReadByte()
	require.NoError(t, err)
	require.Equal(t, byte('*'), typ)

	tmp, _ = reader.ReadByte()
	// numElements, _ := strconv.Atoi(string(tmp))
	assert.Equal(t, byte('1'), tmp, string(tmp))

	// \r\n
	reader.ReadByte()
	reader.ReadByte()

	elementType, err := reader.ReadByte()
	require.NoError(t, err)
	require.Equal(t, byte('$'), elementType, string(elementType))

	tmp, err = reader.ReadByte()
	require.NoError(t, err)
	assert.Equal(t, byte('4'), tmp, string(tmp))
	// length, _ := strconv.Atoi(string(tmp))

	reader.ReadByte()
	reader.ReadByte()

	str, _ := reader.ReadString('\r')
	str = strings.TrimSuffix(str, "\r")
	assert.Equal(t, "ping", str)

	reader.ReadByte()
	reader.ReadByte()

	_, err = reader.ReadByte()
	require.True(t, err == io.EOF, err.Error())
}

func TestByteToASCII(t *testing.T) {
	// byte to ascii, byte <=> unit8
	fmt.Printf("%q %q %v\n", byte(42), byte(36), string(byte(36)))
}

func Test_readLine(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		input := "*1\r\n$5\r\nhello\r\n"
		reader := bufio.NewReader(strings.NewReader(input))
		resp := newResp(reader)

		data, nbyte, err := resp.readLine()
		require.NoError(t, err)
		assert.Equal(t, []byte{'*', '1'}, data)
		assert.Equal(t, 4, nbyte)

		data, nbyte, err = resp.readLine()
		require.NoError(t, err)
		assert.Equal(t, []byte{'$', '5'}, data)
		assert.Equal(t, 4, nbyte)
	})

	t.Run("case2", func(t *testing.T) {
		input := "\r\n"
		reader := bufio.NewReader(strings.NewReader(input))
		resp := newResp(reader)

		data, nbyte, err := resp.readLine()
		require.NoError(t, err)
		assert.Equal(t, []byte{}, data)
		assert.Equal(t, 2, nbyte)

		data, nbyte, err = resp.readLine()
		require.ErrorIs(t, err, io.EOF)
		assert.Equal(t, 0, len(data))
		assert.Equal(t, 0, nbyte)
	})
}
