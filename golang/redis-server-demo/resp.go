package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []value
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

	switch typ {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		fmt.Printf("Unknown type: %v", string(typ))
		return value{}, nil
	}
}

func (r *resp) readLine() ([]byte, int, error) {
	var counter int
	var line []byte

	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		counter += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' { // \r\n
			break
		}
	}

	return line, counter, nil
}

func (r *resp) readInteger() (int, int, error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}

	num, err := strconv.Atoi(string(line))
	if err != nil {
		return 0, n, err
	}
	return int(num), n, nil
}

func (r *resp) readArray() (value, error) {
	panic("TODO")
}

func (r *resp) readBulk() (value, error) {
	panic("TODO")
}
