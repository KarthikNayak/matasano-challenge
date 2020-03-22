package set1

import (
	"strings"
	"testing"
)

func TestQ7(t *testing.T) {
	str, _ := Q7("q7_data")
	if !strings.Contains(str, "I'm back and I'm ringin' the bell") {
		t.Fatalf("couldn't decrypt correctly")
	}
}
