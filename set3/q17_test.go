package set3

import (
	"math/rand"
	"testing"
)

func TestCBCRandomDataEncoder(t *testing.T) {
	key := make([]byte, bSize)
	rand.Read(key)

	IV := make([]byte, bSize)
	rand.Read(IV)

	x, err := cbcRandomDataEncoder(key, IV)

	if len(x) == 0 {
		t.Errorf("got length 0")
	}
	if err != nil {
		t.Errorf("got error: %v", err)
	}
}

func TestCheckEncodedPadding(t *testing.T) {
	key := make([]byte, bSize)
	rand.Read(key)

	IV := make([]byte, bSize)
	rand.Read(IV)

	x, _ := cbcRandomDataEncoder(key, IV)

	if !checkEncodedPadding(x, key, IV) {
		t.Errorf("expected to be true, got false")
	}
}

func TestSolve17(t *testing.T) {
	err := Solve17()
	if err != nil {
		t.Errorf("got error: %v", err)
	}
}
