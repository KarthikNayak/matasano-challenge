package set4

import "testing"

func TestSolve32(t *testing.T) {
	t.Skip("too big")
	go server2()

	_, err := Solve32()
	if err != nil {
		t.Fatal(err)
	}
}
