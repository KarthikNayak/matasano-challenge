package set1

import (
	"testing"
)

func TestSolveQ8(t *testing.T) {
	value := SolveQ8("question8_data")
	if value != true {
		t.Errorf("Expected value: True; Got: %v", value)
	}
}
