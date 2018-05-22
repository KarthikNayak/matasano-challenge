package hex

import (
	"github.com/pkg/errors"
)

func Xor(a, b string) (string, error) {
	if len(a) != len(b) {
		return "", errors.New("Cannot XOR when the strings are of different length")
	}
	byteA, err := HexToByte(a)
	if err != nil {
		return "", err
	}
	byteB, err := HexToByte(b)
	if err != nil {
		return "", err
	}

	outputBytes := make([]byte, len(byteA))
	for i := range outputBytes {
		outputBytes[i] = byteA[i] ^ byteB[i]
	}
	return ByteToHex(outputBytes)
}
