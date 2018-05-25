package set1

import (
	"gitlab.com/karthiknayak/matasano/operations"
	"gitlab.com/karthiknayak/matasano/types"
)

func SolveQ2(a, b string) (string, error) {
	h1 := types.Hex{S: a}
	h2 := types.Hex{S: b}

	c, err := operations.Xor(&h1, &h2)
	return c.Get(), err
}