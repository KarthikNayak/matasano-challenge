package types

import (
	"bytes"
	"testing"
)

func TestPKCS7_Encode(t *testing.T) {
	tests := []struct {
		name string
		input, output []byte
		size int
		err error
	}{
		{
			"Test 1",
			[]byte("YELLOW SUBMARINE"),
			[]byte{89, 69, 76, 76, 79, 87, 32, 83, 85, 66, 77, 65, 82, 73, 78, 69, 4, 4, 4, 4},
			20,
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := PKCS7{}
			p.SetBlockSize(test.size)
			err := p.Encode(test.input)
			if bytes.Compare(p.B, test.output) != 0 {
				t.Errorf("Expected output: %v, received output: %v", test.output, p.B)
			}
			if err != test.err {
				t.Errorf("Expected error: %v, received error: %v", test.err, err)
			}
		})
	}
}

func TestPKCS7_Decode(t *testing.T) {
	tests := []struct {
		name string
		input, output []byte
		err error
	}{
		{
			"Test 1",
			[]byte{89, 69, 76, 76, 79, 87, 32, 83, 85, 66, 77, 65, 82, 73, 78, 69, 4, 4, 4, 4},
			[]byte("YELLOW SUBMARINE"),
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := PKCS7{}
			p.Set(test.input)
			output, err := p.Decode()
			if bytes.Compare(output, test.output) != 0 {
				t.Errorf("Expected output: %v, received output: %v", test.output, output)
			}
			if err != test.err {
				t.Errorf("Expected error: %v, received error: %v", test.err, err)
			}
		})
	}
}
