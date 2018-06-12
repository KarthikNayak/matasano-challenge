package types

import (
	"strings"
	"testing"
)

func TestBase64_Decode(t *testing.T) {
	tests := []struct {
		base64     string
		decoded string
		err     error
		name    string
	}{
		{
			"SGVsbG8gdGhlcmUh",
			"Hello there!",
			nil,
			"Test 1",
		},
		{
			"YW55IGNhcm5hbCBwbGVhcw==",
			"any carnal pleas",
			nil,
			"Test 2",
		},
		{
			"YW55IGNhcm5hbCBwbGVhc3U=",
			"any carnal pleasu",
			nil,
			"Test 3",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b64 := Base64{S: test.base64}
			output, err := b64.Decode()
			if output != test.decoded {
				t.Errorf("Expected output: %v output recieved: %v", test.decoded, output)
			}
			if (err != nil) != (test.err != nil) {
				t.Errorf("Expected error: %v error recieved: %v", test.err, err)
			}
		})
	}
}

func TestBase64_Encode(t *testing.T) {
	tests := []struct {
		s       string
		encoded string
		err     error
		name    string
	}{
		{
			"Hello there!",
			"SGVsbG8gdGhlcmUh",
			nil,
			"Simple Text to hex",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b64 := Base64{}
			err := b64.Encode(test.s)
			if !strings.EqualFold(b64.S, test.encoded) {
				t.Errorf("Expected output: %v output recieved: %v", test.encoded, b64.S)
			}
			if (err != nil) != (test.err != nil) {
				t.Errorf("Expected error: %v error recieved: %v", test.err, err)
			}
		})
	}
}
