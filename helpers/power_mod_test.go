package helpers

import (
	"fmt"
	"testing"
)

func TestPow(t *testing.T) {
	tests := []struct {
		a int
		b int
		c int
	}{
		{2, 2, 4},
		{2, -1, 1},
		{2, -2, 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d ^ %d", test.a, test.b), func(t *testing.T) {
			ans := Pow(test.a, test.b)
			if ans != test.c {
				t.Fatalf("expected: %d got: %d", test.c, ans)
			}
		})
	}
}

func TestPowMod(t *testing.T) {
	tests := []struct {
		a int
		b int
		m int
		v int
	}{
		{2, 4, 5, 1},
		{3, 3, 4, 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d ^ %d", test.a, test.b), func(t *testing.T) {
			ans := PowMod(test.a, test.b, test.m)
			if ans != test.v {
				t.Fatalf("expected: %d got: %d", test.v, ans)
			}
		})
	}
}
