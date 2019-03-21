package set6

import (
	"crypto/sha256"
	"errors"
	"matasano/helpers"
	"matasano/keys"
	"math/big"
)

func signVerifyFunctions() (func(msg string) (*big.Int, error), func(msg string, b *big.Int) bool, int) {
	r := new(keys.RSA)
	r.GenerateKeys()

	s := func(msg string) (*big.Int, error) {
		return r.SignDataSha256([]byte(msg))
	}

	v := func(msg string, b *big.Int) bool {
		d := r.EncryptBigInt(b)
		return checkSignedDataSha256(msg, d.Bytes())
	}

	return s, v, (r.N.BitLen() + 7) / 8
}

func checkSignedDataSha256(msg string, block []byte) bool {
	h := sha256.Sum256([]byte(msg))

	i := 0

	// byte 0 -> 1
	if block[i] != 1 {
		return false
	}
	i += 1

	// followed by 0xff
	for block[i] == 0xff {
		i++
	}

	// followed by 0
	if block[i] != 0 {
		return false
	}
	i += 1

	// followed by the prefix for sha256
	prefix := []byte{0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20}
	for _, val := range prefix {
		if block[i] != val {
			return false
		}
		i += 1
	}

	// check hash
	for _, val := range h {
		if block[i] != val {
			return false
		}
		i += 1
	}
	return true
}

func createFakePadding(h []byte, count int, size int) []byte {
	em := make([]byte, 0)
	em = append(em, byte(1))
	for j := 0; j < count; j++ {
		em = append(em, byte(0xff))
	}
	em = append(em, byte(0))

	prefix := []byte{0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20}
	em = append(em, prefix...)
	em = append(em, h...)

	em = append(em, make([]byte, size-len(em))...)

	return em
}

func createFakeSignature(msg string, size int) *big.Int {
	h := sha256.Sum256([]byte(msg))

	f := new(big.Int).SetBytes(createFakePadding(h[:], 1, size))
	val, _ := helpers.CubeRoot(f)
	return val
}

func SolveQ42() error {
	msg := "hi mom"

	_, v, size := signVerifyFunctions()

	d := createFakeSignature(msg, size)

	if !v(msg, d) {
		return errors.New("expected a pass")
	}

	return nil
}
