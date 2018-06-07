package set1

import (
	"gitlab.com/karthiknayak/matasano/cipher"
	"gitlab.com/karthiknayak/matasano/metrics/frequency"
	"gitlab.com/karthiknayak/matasano/types"
)

func SolveQ3(s string) (string, error) {
	h := types.Hex{S: s}
	f := frequency.CharacterFrequency{}
	s, _, err := cipher.SingleByteXOR(&h, &f)
	return s, err
}