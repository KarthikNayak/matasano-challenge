package set1

import (
	"strings"
	"testing"
)

func TestQ6(t *testing.T) {
	ans, err := Q6("q6_data")

	if !strings.Contains(ans, "I'm back and I'm ringin' the bell") {
		t.Fatalf("doesn't contain the require answer")
	}

	if err != nil {
		t.Fatalf(err.Error())
	}
}
