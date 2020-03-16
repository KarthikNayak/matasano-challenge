package helpers

func Xor(b1, b2 []byte) (b []byte) {
	b = make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		b[i] = b1[i] ^ b2[i]
	}
	return
}
