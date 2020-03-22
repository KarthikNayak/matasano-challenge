package padding

func PKCS7(b []byte, l int) []byte {
	var f []byte = make([]byte, l)

	for i := 0; i < len(b); i++ {
		f[i] = b[i]
	}

	for i := len(b); i < l; i++ {
		f[i] = byte(l - len(b))
	}
	return f
}
