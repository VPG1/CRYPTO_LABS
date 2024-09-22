package cipher

func EncryptBlock(block []uint32, key []uint32) []uint32 {
	tactKeys := make([]uint32, 57)
	for i := 1; i < len(tactKeys); i++ {
		tactKeys[i] = key[(i-1)%8]
	}
	a, b, c, d := block[0], block[1], block[2], block[3]

	for i := 1; i <= 8; i++ {
		b = b ^ MappingG(sum(a, tactKeys[7*i-6]), 5)                             // 1
		c = c ^ MappingG(sum(d, tactKeys[7*i-5]), 21)                            // 2
		a = sub(a, MappingG(sum(b, tactKeys[7*i-4]), 13))                        // 3
		e := MappingG(sum(b, c, tactKeys[7*i-3]), 21) ^ reverseUint32(uint32(i)) // 4
		b = sum(b, e)                                                            // 5
		c = sub(c, e)                                                            // 6
		d = sum(d, MappingG(sum(c, tactKeys[7*i-2]), 13))                        // 7
		b = b ^ MappingG(sum(a, tactKeys[7*i-1]), 21)                            // 8
		c = c ^ MappingG(sum(d, tactKeys[7*i]), 5)                               // 9
		a, b = b, a                                                              // 10
		c, d = d, c                                                              // 11
		b, c = c, b                                                              // 12
	}

	resBlock := make([]uint32, 4)
	resBlock[0], resBlock[1], resBlock[2], resBlock[3] = b, d, a, c
	return resBlock
}
