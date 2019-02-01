package set2

import "matasano/cipher"

func SolveQ11() (bool, error) {
	var oracle cipher.Oracle
	oracle = cipher.EncryptionOracle

	return cipher.DetectECB(oracle)
}
