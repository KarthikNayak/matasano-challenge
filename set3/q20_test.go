package set3

import (
	"fmt"
	"testing"
)

func TestSolveQ20(t *testing.T) {
	str, err := SolveQ20()
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(str); i += 16 {
		fmt.Println(str[i : i+16])
	}
}
