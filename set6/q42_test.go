package set6

import "testing"

func TestSolveQ42(t *testing.T) {
	err := SolveQ42()
	if err != nil {
		t.Fatal(err)
	}
}
