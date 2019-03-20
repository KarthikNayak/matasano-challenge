package helpers

import (
	"fmt"
	"math/big"
	"testing"
)

func TestInverseMod(t *testing.T) {
	tests := []struct {
		x, y *big.Int
		z    *big.Int
	}{
		{
			new(big.Int).SetInt64(17),
			new(big.Int).SetInt64(3120),
			new(big.Int).SetInt64(2753),
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("mod(%d,%d)", test.x, test.y), func(t *testing.T) {
			z := InverseMod(test.x, test.y)
			if z.Cmp(test.z) != 0 {
				t.Errorf("compare failed expected: %d got: %d", test.z, z)
			}
		})
	}
}
