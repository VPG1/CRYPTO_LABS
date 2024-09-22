package main

import (
	"Lab2/cipher"
	"encoding/binary"
	"fmt"
)

func main() {
	data := []uint32{0, 0, 0, 1}
	key := []uint32{0, 0, 0, 0, 0, 0, 0, 1}
	fmt.Println(data)
	fmt.Println(key)
	encryptData := cipher.EncryptBlock(data, key)
	fmt.Println(encryptData)
	decryptData := cipher.DecryptBlock(encryptData, key)
	fmt.Println(decryptData)

	fmt.Println(uint32(1) & 0xFF0000)

	bytesBuffer := make([]byte, 4)
	binary.BigEndian.PutUint32(bytesBuffer, 1)
	fmt.Println(bytesBuffer)
}
