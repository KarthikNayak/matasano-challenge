package set6

import (
	"math/big"
	"testing"
)

func TestGetQ44Data(t *testing.T) {
	d := GetQ44Data()
	if string(d[0].msg) != "Listen for me, you better listen for me now. " {
		t.Fail()
	}

	v, _ := new(big.Int).SetString("29097472083055673620219739525237952924429516683", 10)
	if d[1].s.Cmp(v) != 0 {
		t.Fail()
	}

	if len(d) != 11 {
		t.Fail()
	}
}

func TestGetIdenticalKsQ44(t *testing.T) {
	a, b := GetIdenticalKs()
	if a == nil || b == nil {
		t.Fail()
	}
	if a.r.Cmp(b.r) != 0 {
		t.Fail()
	}
}

func TestSolveQ44(t *testing.T) {
	if !SolveQ44() {
		t.Fail()
	}
}
