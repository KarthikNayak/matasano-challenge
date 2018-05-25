package types

import (
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestHex_Decode(t *testing.T) {
	tests := []struct {
		hex     string
		decoded string
		err     error
		name    string
	}{
		{
			"48656c6c6f20746865726521",
			"Hello there!",
			nil,
			"Simple Text to hex",
		},
		{
			"48656c6c6f2074686572652",
			"",
			errors.New("hex array length is odd"),
			"odd hex array length",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hex := Hex{S: test.hex}
			output, err := hex.Decode()
			if output != test.decoded {
				t.Errorf("Expected output: %v output recieved: %v", test.decoded, output)
			}
			if (err != nil) != (test.err != nil) {
				t.Errorf("Expected error: %v error recieved: %v", test.err, err)
			}
		})
	}
}

func TestHex_Encode(t *testing.T) {
	tests := []struct {
		s       string
		encoded string
		err     error
		name    string
	}{
		{
			"Hello there!",
			"48656c6c6f20746865726521",
			nil,
			"Simple Text to hex",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hex := Hex{}
			err := hex.Encode(test.s)
			if !strings.EqualFold(hex.S, test.encoded) {
				t.Errorf("Expected output: %v output recieved: %v", test.encoded, hex.S)
			}
			if (err != nil) != (test.err != nil) {
				t.Errorf("Expected error: %v error recieved: %v", test.err, err)
			}
		})
	}
}
