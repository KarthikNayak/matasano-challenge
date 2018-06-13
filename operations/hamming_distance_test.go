package operations

import (
	"testing"

	"gitlab.com/karthiknayak/matasano/types"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		distance int
	}{
		{
			"Test 1",
			"this is a test",
			"wokka wokka!!!",
			37,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h1 := types.Hex{}
			h1.Encode([]byte(test.a))
			h2 := types.Hex{}
			h2.Encode([]byte(test.b))

			distance, err := HammingDistance(&h1, &h2)
			if err != nil {
				t.Errorf("Got an unexpected error %v", err)
			}

			if distance != 37 {
				t.Errorf("Expected hamming distance: 37 got: %v", distance)
			}
		})
	}
}

