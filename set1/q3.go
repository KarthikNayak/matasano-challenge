package set1

import (
	"cryptopals/cipher"
	"cryptopals/types"
)

func Q3(h types.Hex) (float64, string) {
	b, _ := h.Decode()
	score, ans := cipher.DecodeSingleByteXOR(b)
	return score, string(ans)
}
