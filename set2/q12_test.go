package set2

import (
	"cryptopals/oracle"
	"strings"
	"testing"
)

func TestBreakECB(t *testing.T) {
	output := BreakIfECB(oracle.FixedECBWithUnkownText)

	if !strings.Contains(string(output), "Rollin' in my 5.0") {
		t.FailNow()
	}
}
