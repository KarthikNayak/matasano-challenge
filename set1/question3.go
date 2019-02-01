package set1

import (
	"matasano/cipher"
	"matasano/types"
)

func SolveQ3(b []byte) (string, error) {
	h := types.Hex{B: b}
	s, _, err := cipher.DecodeSingleByteXOR(&h)
	return s, err
}
