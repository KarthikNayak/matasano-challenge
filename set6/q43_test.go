package set6

import (
	"fmt"
	"strings"
	"testing"
)

const (
	msgHash = "d2d0714f014a9784047eaeccf956520045c45265"
)

func TestHashMsg(t *testing.T) {
	h := fmt.Sprintf("%x", msgSha())
	if strings.Compare(h, msgHash) != 0 {
		t.Fatalf("expected: %x got: %x", msgHash, h)
	}
}

func TestSolveQ43(t *testing.T) {
	SolveQ43()
}
