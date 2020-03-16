package set1

import (
	"cryptopals/types"
	"testing"
)

func TestQ3(t *testing.T) {
	h1 := types.Hex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	_, msg := Q3(h1)
	if msg != "Cooking MC's like a pound of bacon" {
		t.Fatalf("didn't get expected result")
	}
}
