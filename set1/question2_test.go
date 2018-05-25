package set1

import (
	"testing"
)

func TestSolveQ2(t *testing.T) {
	h1 := "1c0111001f010100061a024b53535009181c"
	h2 := "686974207468652062756c6c277320657965"
	expectedOutput := "746865206B696420646F6E277420706C6179"

	output, err := SolveQ2(h1, h2)
	if err != nil {
		t.Errorf("Got an unexpected error %v", err)
	}
	if output != expectedOutput {
		t.Errorf("Expected output: %v obtained output: %v", expectedOutput, output)
	}
}

