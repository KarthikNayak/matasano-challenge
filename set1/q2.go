package set1

import (
	"cryptopals/helpers"
	"cryptopals/types"
)

func Q2(h1, h2 types.Hex) types.Hex {
	b1, _ := h1.Decode()
	b2, _ := h2.Decode()
	b := helpers.Xor(b1, b2)
	return types.EncodeHex(b)
}
