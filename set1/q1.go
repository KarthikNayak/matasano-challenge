package set1

import (
	"cryptopals/types"
)

func Q1(h types.Hex) types.Base64 {
	b, _ := h.Decode()
	return types.EncodeBase64(b)
}
