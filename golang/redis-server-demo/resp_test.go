package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRespMessage(t *testing.T) {
	// TODO what's the difference between "*1\r\n$4\r\nping\r\n" and `*1\r\n$4\r\nping\r\n` ?
	// cuz I found `\r\n` needs 4 reader.ReadByte(), while "\r\n" only needs 2

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
