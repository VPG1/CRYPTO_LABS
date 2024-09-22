package cipher

import "CRYPTO_LABS/converters"

func DecryptBlock(block []byte, key []byte) ([]byte, error) {
	blockUint32, err := converters.ConvertByteArrToUint32Arr(block)
	if err != nil {
		return nil, err
	}

	keyUint32, err := converters.ConvertByteArrToUint32Arr(key)
	if err != nil {
		return nil, err
	}

	tactKeys := make([]uint32, 57)
	for i := 1; i < len(tactKeys); i++ {
		tactKeys[i] = keyUint32[(i-1)%8]
	}
	a, b, c, d := blockUint32[0], blockUint32[1], blockUint32[2], blockUint32[3]

	for i := 8; i >= 1; i-- {
		b = b ^ MappingG(sum(a, tactKeys[7*i]), 5)                               // 1
		c = c ^ MappingG(sum(d, tactKeys[7*i-1]), 21)                            // 2
		a = sub(a, MappingG(sum(b, tactKeys[7*i-2]), 13))                        // 3
		e := MappingG(sum(b, c, tactKeys[7*i-3]), 21) ^ reverseUint32(uint32(i)) // 4
		b = sum(b, e)                                                            // 5
		c = sub(c, e)                                                            // 6
		d = sum(d, MappingG(sum(c, tactKeys[7*i-4]), 13))                        // 7
		b = b ^ MappingG(sum(a, tactKeys[7*i-5]), 21)                            // 8
		c = c ^ MappingG(sum(d, tactKeys[7*i-6]), 5)                             // 9
		a, b = b, a                                                              // 10
		c, d = d, c                                                              // 11
		a, d = d, a                                                              // 12
	}

	resBlockUint32 := make([]uint32, 4)
	resBlockUint32[0], resBlockUint32[1], resBlockUint32[2], resBlockUint32[3] = c, a, d, b

	return converters.ConvertUint32ArrToByteArr(resBlockUint32), nil
}
