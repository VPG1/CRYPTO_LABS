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

func add128Bit(a, b []byte) []byte {
	// Длина 128-битного числа (16 байт)
	n := 16
	result := make([]byte, n)
	var carry uint16 = 0

	// Складываем числа побайтно начиная с младших байтов
	for i := n - 1; i >= 0; i-- {
		// Складываем байты как uint16, чтобы учесть перенос
		sum := uint16(a[i]) + uint16(b[i]) + carry
		result[i] = byte(sum) // Младший байт суммы сохраняем в результат
		carry = sum >> 8      // Старший байт используется как перенос
	}

	// Результат автоматически "обрезан" до 128 бит, так как излишний перенос просто игнорируется
	return result
}

func xor128Bit(a, b []byte) []byte {
	// Длина 128-битного числа (16 байт)
	n := 16
	result := make([]byte, n)

	// Выполняем XOR для каждого байта
	for i := 0; i < n; i++ {
		result[i] = a[i] ^ b[i]
	}

	return result
}
