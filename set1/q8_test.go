package set1

import "testing"

func TestQ8(t *testing.T) {
	if !Q8("q8_data") {
		t.Fatalf("couldn't find ECB encrypted data")
	}
}
