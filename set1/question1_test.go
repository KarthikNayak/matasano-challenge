package set1

import (
	"testing"
)

func TestSolveQ1(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	output, err := SolveQ1([]byte(input))
	if err != nil {
		t.Errorf("Got an unexpected error %v", err)
	}
	if string(output) != expectedOutput {
		t.Errorf("Expected output: %v obtained output: %v", expectedOutput, output)
	}
}
