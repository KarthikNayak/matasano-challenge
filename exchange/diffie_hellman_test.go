package exchange

import (
	"bytes"
	"testing"
)

func TestDiffieHellman(t *testing.T) {
	var Alice, Bob DFClient

	Alice = new(DF)
	Bob = new(DF)

	err := Alice.SendPGA(Bob)
	if err != nil {
		t.Fatal(err)
	}
	err = Bob.SendB(Alice)
	if err != nil {
		t.Fatal(err)
	}

	Alice.GenHash()
	Bob.GenHash()

	hash1, hash2 := Alice.GetHash(), Bob.GetHash()
	if bytes.Compare(hash1, hash2) != 0 {
		t.Fatal("expected the hashes to be the same")
	}
}
