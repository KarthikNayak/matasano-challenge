package set1

import (
	"gitlab.com/karthiknayak/matasano/cipher"
	"gitlab.com/karthiknayak/matasano/metrics/frequency"
	"gitlab.com/karthiknayak/matasano/types"
)

func SolveQ3(b []byte) (string, error) {
	h := types.Hex{B: b}
	f := frequency.CharacterFrequency{}
	s, _, err := cipher.DecodeSingleByteXOR(&h, &f)
	return s, err
}