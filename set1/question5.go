package set1

import (
	"gitlab.com/karthiknayak/matasano/cipher"
	"gitlab.com/karthiknayak/matasano/types"
)

func SolveQ5(input, key []byte) ([]byte, error) {
	h := types.Hex{}
	h.Encode(input)

	return cipher.EncodeRepeatingXor(&h, key)
}
