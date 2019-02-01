package set2

import (
	"fmt"
	"testing"
)

func TestQ11(t *testing.T) {
	for i := 0; i < 100; i++ {
		b, err := SolveQ11()
		if err != nil {
			t.Error(err)
		}
		if b {
			fmt.Println("Detected ECB Cipher")
		} else {
			fmt.Println("Detected CBC Cipher")
		}
	}
}
