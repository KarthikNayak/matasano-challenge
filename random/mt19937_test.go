package random

import (
	"testing"
)

func TestMT19937Init(t *testing.T) {
	var m MT19937
	m.init()
	if len(m.data) != _N-1 {
		t.Error("length of data incorrect")
	}
	if m.index != _N+1 {
		t.Error("incorrect index")
	}

	if m.lowerMask != 0x7fffffff {
		t.Errorf("incorrect lower_mask expected: %x, got: %x",
			0x7fffffff, m.lowerMask)
	}

	if m.upperMask != 0x80000000 {
		t.Errorf("incorrect upper_mask expected: %x, got: %x",
			0x80000000, m.upperMask)
	}
}

func TestMT19937Uint32(t *testing.T) {
	table := []uint32{
		1055721139,
		3422054626,
		2561641375,
		1376353668,
		1540998321,
		825546192,
		1627406507,
		1797302575,
		105017825,
		1514647480,
	}

	var m MT19937
	m.Seed(12345678)

	for _, val := range table {
		x := m.Uint32()
		if val != x {
			t.Fatalf("expected value: %x got: %x", val, x)
		}
	}

}
