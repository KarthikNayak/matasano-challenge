package random

import (
	"fmt"
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
	fmt.Println((ReverseBitFlip(1797302575)))
}
