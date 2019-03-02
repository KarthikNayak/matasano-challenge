package set1

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			"Line 1",
			`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`,
			"0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			xorKey := "ICE"
			s, err := SolveQ5([]byte(test.input), []byte(xorKey))
			if string(s) != test.output {
				t.Errorf("Expected output: %v obtained output: %v", test.output, s)
			}
			if err != nil {
				t.Errorf("Got an unexpected error %v", err)
			}
		})
	}
}
