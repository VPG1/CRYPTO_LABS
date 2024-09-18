package main

import (
	"ISM_LAB1/converters"
	"ISM_LAB1/imitation_insert"
	"encoding/hex"
	"fmt"
	"strconv"
)

// key: 1f2e3f4f5f6f7f8f1f2e3f4f5f6f7f8f1f2e3f4f5f6f7f8f1f2e3f4f5f6f7f89

// 6972694b 6972694b 67ce1c6c
// cdce1c6c

// 6b279800

// 34279800

// a61d351a

func main() {
	fmt.Println("Enter data:")
	var data string
	_, err := fmt.Scanln(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter key(len 32 bytes):")
	var keyStr string
	_, err = fmt.Scanln(&keyStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	hexArr, err := hex.DecodeString(keyStr)
	if err != nil {
		fmt.Println("The key is invalid")
		return
	}

	if len(hexArr) != 32 {
		fmt.Println("The key must be 32 bytes long")
		return
	}

	key, err := converters.ConvertByteArrToUint32Arr(hexArr)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := imitation_insert.DevImitationInsert([]byte(data), [8]uint32(key))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Imitation insert:", strconv.FormatUint(uint64(res), 16))

}
