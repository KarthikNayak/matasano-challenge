package set6

import (
	"math/big"
	"testing"
)

func TestSolveQ41(t *testing.T) {
	var val int64 = 45
	got := SolveQ41(val)
	if val != got {
		t.Fatalf("expected: %v got %v", val, got)
	}
}

func TestRSAOracleDecrypt(t *testing.T) {
	r, f := rsaOracleDecrypt()

	c := r.EncryptBigInt(new(big.Int).SetInt64(25))
	m := f(c)
	if m.Int64() != 25 {
		t.Fatalf("expected %v got %v", 25, m.Int64())
	}

	n := f(c)
	if n != nil {
		t.Fatalf("expected %v got %v", nil, n)
	}
}
