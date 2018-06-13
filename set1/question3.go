package set1

import (
	"matasano/cipher"
	"matasano/metrics/frequency"
	"matasano/types"
)

func SolveQ3(b []byte) (string, error) {
	h := types.Hex{B: b}
	f := frequency.CharacterFrequency{}
	s, _, err := cipher.DecodeSingleByteXOR(&h, &f)
	return s, err
}
