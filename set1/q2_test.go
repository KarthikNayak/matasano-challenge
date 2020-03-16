package set1

import (
	"bytes"
	"cryptopals/types"
	"testing"
)

func TestQ2(t *testing.T) {
	h1 := types.Hex("1c0111001f010100061a024b53535009181c")
	h2 := types.Hex("686974207468652062756c6c277320657965")
	expected, _ := types.Hex("746865206b696420646f6e277420706c6179").Decode()

	output := Q2(h1, h2)
	got, _ := output.Decode()
	if bytes.Compare(expected, got) != 0 {
		t.Fatalf("expected: %v got: %v", expected, got)
	}
}
