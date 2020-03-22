package padding

func PKCS7(b []byte, blockSize int) []byte {

	srcLen := len(b)
	padding := 0

	if srcLen > blockSize {
		extra := srcLen % blockSize
		if extra > 0 {
			padding = blockSize - extra
		}
	} else {
		padding = blockSize - srcLen
	}

	output := make([]byte, srcLen+padding)
	copy(output, b)

	for i := 0; i < padding; i++ {
		output[srcLen+i] = byte(padding)
	}

	return output
}
