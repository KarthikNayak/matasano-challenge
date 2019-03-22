package keys

import (
	"testing"
)

func TestDSA(t *testing.T) {
	d := new(DSA)
	d.GenerateKeys()

	msg := "dude"
	r, s, err := d.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}

	ret := d.Verify(r, s, msg)
	if !ret {
		t.Fatal("expected to pass")
	}
}
