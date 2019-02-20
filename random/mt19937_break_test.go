package random

import (
	"testing"
	"time"
)

func TestFindSeed(t *testing.T) {
	seed := uint32(time.Now().Unix())
	var m MT19937
	m.Seed(seed)
	val := FindSeed(m.Uint32())
	if val != seed {
		t.Errorf("expected value: %x got: %x", seed, val)
	}
}

func TestReverseBitFlip(t *testing.T) {
	if ReverseBitFlip(1797302575) != 3019169788 {
		t.Errorf("expected value :%v got something else", 3019169788)
	}
}

func TestCloneMT19937(t *testing.T) {
	var m MT19937
	m.Seed(12342)

	clone := CloneMT19937(m)
	for i := 0; i < 1000; i++ {
		if clone.Uint32() != m.Uint32() {
			t.Errorf("non-matching numbers generated")
		}
	}
}
