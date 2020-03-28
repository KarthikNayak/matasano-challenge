package set2

import (
	"strings"
	"testing"
)

func TestQ13(t *testing.T) {
	result := Q13()
	if !strings.Contains(result, "role=admin") {
		t.FailNow()
	}
}
