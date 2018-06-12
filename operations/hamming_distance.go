package operations

import (
	"github.com/pkg/errors"
	"gitlab.com/karthiknayak/matasano/types"
)

func HammingDistance(a, b types.Cipher) (int, error) {
	s1, err := a.Decode()
	if err != nil {
		return 0, err
	}
	s2, err := b.Decode()
	if err != nil {
		return 0, err
	}

	if len(s1) != len(s2) {
		return 0, errors.New("Length of given strings don't match")
	}

	hDistance := 0

	for i := range s1 {
		val := s1[i] ^ s2[i]
		for val > 0 {
			hDistance += int(val) % 2
			val = val / 2
		}
	}
	return hDistance, nil
}

