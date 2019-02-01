package set1

import (
	"matasano/helpers"
	"matasano/types"
)

func SolveQ2(a, b []byte) (string, error) {
	h1 := types.Hex{B: a}
	h2 := types.Hex{B: b}

	c, err := helpers.Xor(&h1, &h2)
	return string(c.Get()), err
}
