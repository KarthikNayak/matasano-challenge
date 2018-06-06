package set1

import (
	"testing"
)

func TestSolveQ3(t *testing.T) {
	h := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expectedOutput := "Cooking MC's like a pound of bacon"

	output, err := SolveQ3(h)
	if err != nil {
		t.Errorf("Got an unexpected error %v", err)
	}
	if output != expectedOutput {
		t.Errorf("Expected output: %v obtained output: %v", expectedOutput, output)
	}
}
