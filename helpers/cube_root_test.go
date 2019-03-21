package helpers

import (
	"math/big"
	"testing"
)

func TestCubeRoot(t *testing.T) {
	test := new(big.Int).SetInt64(125)
	val, rem := CubeRoot(test)
	if val.Int64() != 5 || rem.Int64() != 0 {
		t.Fatalf("expected: (%v, %v) got: (%v, %v)", 25, 5, val.Int64(), rem.Int64())
	}
}
