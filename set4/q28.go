package set4

import (
	"errors"
	"matasano/hash"
)

func PrefixMAC(key, message []byte) string {
	final := append(key, message...)
	return hash.Sha1(final)
}

func SolveQ28() error {
	key := []byte("IamASmartKey")
	data := []byte("IamSomeSecretData")
	sha1 := PrefixMAC(key, data)

	data2 := []byte("IamSomeSecretDate")
	sha2 := PrefixMAC(key, data2)

	if sha1 == sha2 {
		return errors.New("same sha1")
	}
	return nil
}
