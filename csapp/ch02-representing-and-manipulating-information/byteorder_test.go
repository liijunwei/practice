package main

import "github.com/stretchr/testify/assert"

func Test_ByteOrder(t *testing.T) {
	bufSize := 2

	tests := []struct {
		number                uint16
		bigEndianByteOrder    []byte
		littleEndianByteOrder []byte
	}{
		// 0000 0100 0011 1000 (1 187)
		// 数据的高位字节存放在低位内存地址->bigEndian
		// 数据的地位字节存放在低位内存地址->littleEndian
		{number: 443, bigEndianByteOrder: []byte{1, 187}, littleEndianByteOrder: []byte{187, 1}},
		// 0000 0100 0011 1000
		{number: 1080, bigEndianByteOrder: []byte{4, 56}, littleEndianByteOrder: []byte{56, 4}},
	}

	for _, tt := range tests {
		fmt.Println(strconv.FormatInt(int64(tt.number), 2)) // 1111011

		buf1 := make([]byte, bufSize)
		binary.BigEndian.PutUint16(buf1, tt.number)
		assert.Equal(t, tt.bigEndianByteOrder, buf1)

		buf2 := make([]byte, bufSize)
		binary.LittleEndian.PutUint16(buf2, tt.number)
		assert.Equal(t, tt.littleEndianByteOrder, buf2)
	}
}
