package cipher

import "encoding/binary"

func sum(args ...uint32) uint32 {
	res := uint32(0)
	for _, el := range args {
		res += reverseUint32(el)
	}
	return reverseUint32(res)
}

func sub(x, y uint32) uint32 {
	return reverseUint32(reverseUint32(x) - reverseUint32(y))
}

func reverseUint32(value uint32) uint32 {
	// Преобразуем int32 в байты
	bytesBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytesBuffer, value)

	// Реверсируем байты
	for i, j := 0, len(bytesBuffer)-1; i < j; i, j = i+1, j-1 {
		bytesBuffer[i], bytesBuffer[j] = bytesBuffer[j], bytesBuffer[i]
	}

	// Преобразуем обратно в int32
	return binary.LittleEndian.Uint32(bytesBuffer)
}
