package set1

import (
	"gitlab.com/karthiknayak/matasano/operations"
	"gitlab.com/karthiknayak/matasano/types"
)

func SolveQ2(a, b []byte) (string, error) {
	h1 := types.Hex{B: a}
	h2 := types.Hex{B: b}

	c, err := operations.Xor(&h1, &h2)
	return string(c.Get()), err
}