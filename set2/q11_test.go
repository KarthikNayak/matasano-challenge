package set2

import (
	"testing"
)

func TestQ11(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, err := SolveQ11()
		if err != nil {
			t.Error(err)
		}
	}
}
