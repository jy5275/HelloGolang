package c

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"unsafe"
)

func IsLittleEndian() bool {
	var n int32 = 0x01020304
	u := unsafe.Pointer(&n)
	pb := (*byte)(u)
	b := *pb

	// 由于b是byte类型，最多保存8位，那么只能取得开始的8位
	// Little endian: 04 (03 02 01)
	// Big endian: 01 (02 03 04)
	return b == 0x04
}

func IntToBytes(n uint32) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})

	var order binary.ByteOrder = binary.BigEndian
	if IsLittleEndian() {
		order = binary.LittleEndian
	}

	binary.Write(bytesBuffer, order, x)
	return bytesBuffer.Bytes()
}

func TestMemoryC(t *testing.T) {
	data := Malloc(4)
	fmt.Printf("data %+v, %T\n", data, data)

	myData := (*uint32)(data)
	*myData = 5
	fmt.Printf("data %+v, %T\n", *myData, *myData)

	var a uint32 = 100
	Memcpy(data, IntToBytes(a), 4)
	fmt.Printf("data %+v, %T\n", *myData, *myData)

	Free(data)
}
