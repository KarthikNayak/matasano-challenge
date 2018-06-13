package set1

import (
	"matasano/types"
)

func SolveQ1(h []byte) ([]byte, error) {
	hex := types.Hex{B: h}
	b, err := hex.Decode()
	if err != nil {
		return []byte{}, err
	}

	b64 := types.Base64{}
	b64.Encode(b)
	return b64.B, nil
}
