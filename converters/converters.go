package converters

import (
	"encoding/binary"
	"errors"
)

func ConvertByteArrToUint64Arr(arr []byte) ([]uint64, error) {
	if len(arr)%8 != 0 {
		return nil, errors.New("invalid byte array")
	}

	res := make([]uint64, (len(arr)+7)/8)

	for i := 0; i < len(arr); i += 8 {
		res[i/8] = binary.LittleEndian.Uint64(arr[i : i+8])
	}

	return res, nil
}

func ConvertByteArrToUint32Arr(arr []byte) ([]uint32, error) {
	if len(arr)%4 != 0 {
		return nil, errors.New("invalid byte array")
	}

	res := make([]uint32, (len(arr)+3)/4)

	for i := 0; i < len(arr); i += 4 {
		res[i/4] = binary.LittleEndian.Uint32(arr[i : i+4])
	}

	return res, nil
}

func ConvertUint32ArrToByteArr(arr []uint32) []byte {
	res := make([]byte, len(arr)*4)

	for i := 0; i < len(arr); i++ {
		binary.LittleEndian.PutUint32(res[i*4:(i+1)*4], arr[i])
	}

	return res
}
