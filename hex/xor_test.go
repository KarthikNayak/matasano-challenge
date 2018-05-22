package hex

import (
	"testing"
)

func TestXor(t *testing.T) {
	table := []struct{
		a, b string
		output string
		err error
	}{
		{
			"1c0111001f010100061a024b53535009181c",
			"686974207468652062756c6c277320657965",
			"746865206B696420646F6E277420706C6179",
			nil,
		},
	}

	for _, val := range table {
		output, err := Xor(val.a, val.b)
		if output != val.output {
			t.Errorf("Expected output: %v deviates from obtained output: %v", val.output, output)
		}
		if err != val.err {
			t.Errorf("Expected error: %v deviates from obtained error: %v", val.err, err)
		}
	}
}
