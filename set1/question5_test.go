package set1

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		input string
		output string
	}{
		{
			"Line 1",
			`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`,
			"0B3637272A2B2E63622C2E69692A23693A2A3C6324202D623D63343C2A26226324272765272A282B2F20430A652E2C652A3124333A653E2B2027630C692B20283165286326302E27282F",
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