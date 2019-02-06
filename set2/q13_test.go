package set2

import (
	"testing"
)

func TestQ13(t *testing.T) {
	err := SolveQ13()
	if err != nil {
		t.Error(err)
	}
}
