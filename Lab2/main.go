package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	fmt.Println(uint32(1) & 0xFF0000)

	bytesBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(bytesBuffer, 1)
	fmt.Println(bytesBuffer)
}
