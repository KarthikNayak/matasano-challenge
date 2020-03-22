package helpers

import (
	"errors"
)

func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Length of given strings don't match")
	}

	hDistance := 0

	for i := range a {
		val := a[i] ^ b[i]
		for val > 0 {
			hDistance += int(val) % 2
			val = val / 2
		}
	}
	return hDistance, nil
}
