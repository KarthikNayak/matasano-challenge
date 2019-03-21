package exchange

import (
	"crypto/sha256"
	"testing"
)

func TestSRP(t *testing.T) {
	// pass
}

func TestSha256ToBigInt(t *testing.T) {
	sha256 := sha256.Sum256([]byte("F0E4C2F76C58916EC258F246851BEA091D14D4247A2FC3E18694461B1816E13B"))
	x1 := Sha256ToBigInt(sha256[:])
	x2 := Sha256ToBigInt(sha256[:])

	if x1.Cmp(x2) != 0 {
		t.Fatalf("the outputs were supposed to be the same")
	}
}