package cipher

func CmEncrypt(data []byte, key []byte, S []byte) ([]byte, error) {
	s, err := EncryptBlock(S, key)
	if err != nil {
		return nil, err
	}

	n := (len(data) + 15) / 16

	alignedData := make([]byte, n*16)
	copy(alignedData[:len(data)], data)

	res := make([]byte, len(alignedData))

	for i := 1; i <= n; i++ {
		s = add128Bit(s, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
		newS, err := EncryptBlock(s, key)
		if err != nil {
			return nil, err
		}

		copy(res[(i-1)*16:i*16], xor128Bit(alignedData[(i-1)*16:i*16], newS))
	}

	return res[:len(data)], nil
}
