package keys

import (
	"math/big"
	"testing"
)

func TestRSA(t *testing.T) {
	var r RSA
	err := r.GenerateKeys()
	if err != nil {
		t.Fatal(err)
	}

	m := new(big.Int).SetInt64(42)
	c := r.EncryptBigInt(m)
	newM := r.DecryptBigInt(c)

	if m.Cmp(newM) != 0 {
		t.Fatalf("expected: %d got: %d", m.Int64(), newM.Int64())
	}
}

func TestRSAString(t *testing.T) {
	var r RSA
	err := r.GenerateKeys()
	if err != nil {
		t.Fatal(err)
	}

	m := "foobarwhatthehell"
	c := r.EncryptString(m)
	newM := r.DecryptString(c)

	if m != newM {
		t.Fatalf("expected: %s got: %s", m, newM)
	}
}
