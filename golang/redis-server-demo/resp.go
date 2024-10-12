package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	TYPE_STRING = '+'
	TYPE_ERROR  = '-'
	TYPE_BULK   = '$'
	TYPE_ARRAY  = '*'
)

// RESP spec
// https://redis.io/docs/latest/develop/reference/protocol-spec/
// The first byte in an RESP-serialized payload always identifies its type. Subsequent bytes constitute the type's contents.

// deserilize RESP request message(binary) into value data structure
type value struct {
	Typ   string  `json:"type,omitempty"`
	Str   string  `json:"string,omitempty"`
	Num   int     `json:"number,omitempty"`
	Bulk  string  `json:"bulk,omitempty"`
	Array []value `json:"array,omitempty"`
}

type resp struct {
	reader *bufio.Reader
}

func newResp(reader io.Reader) *resp {
	return &resp{
		reader: bufio.NewReader(reader),
	}
}

func (r *resp) read() (value, error) {
	typ, err := r.reader.ReadByte()
	if err != nil {
		return value{}, err
	}

	fmt.Printf("_type: %v -> %q\n", string(typ), typ)

	switch typ {
	case TYPE_ARRAY:
		return r.readArray()
	case TYPE_BULK:
		return r.readBulk()
	default:
		fmt.Printf("Unknown type: %v", string(typ))
		return value{}, nil
	}
}

// format: *<number-of-elements>\r\n<element-1>...<element-n>
//
// e.g. *3\r\n$3\r\nset\r\n$3\r\nkey\r\n$5\r\nvalue\r\n
// len = 3
// element[0] = set
// element[1] = key
// element[2] = value
func (r *resp) readArray() (value, error) {
	v := value{}
	v.Typ = "array"

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	v.Array = make([]value, 0)
	for i := 0; i < len; i++ {
		val, err := r.read() // recursive call
		if err != nil {
			return v, err
		}

		v.Array = append(v.Array, val)
	}

	return v, nil
}

func (r *resp) readLine() ([]byte, int, error) {
	var nbyte int
	var line []byte

	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		nbyte += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' { // \r\n
			break
		}
	}

	return line[:len(line)-2], nbyte, nil
}

func (r *resp) readInteger() (int, int, error) {
	line, n, err := r.readLine() // content/data between \r\n
	if err != nil {
		return 0, 0, err
	}

	num, err := strconv.Atoi(string(line))
	if err != nil {
		return 0, n, err
	}
	return int(num), n, nil
}

// format: $<length>\r\n<data>\r\n
//
// e.g. $5\r\nvalue\r\n
// len  = 5
// bulk = value
func (r *resp) readBulk() (value, error) {
	v := value{}
	v.Typ = "bulk"

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	bulk := make([]byte, len)
	r.reader.Read(bulk) // read data into buffer
	v.Bulk = string(bulk)

	// see: Test_readLine
	r.readLine()

	return v, nil
}
