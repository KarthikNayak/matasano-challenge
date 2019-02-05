package cipher

import (
	"bytes"
	"math/rand"
)

func getLength(oracle Oracle, key []byte) int {
	b, _ := oracle([]byte{}, key)
	return len(b)
}

func BreakECB(oracle Oracle, bSize int) ([]byte, error) {
	key := make([]byte, bSize)
	rand.Read(key)

	// get the length we would finally need
	length := getLength(oracle, key)

	// final output
	var output []byte

	// We increment block and reset curSize as it hits the last element each
	// time within a particular block
	block := 1
	curSize := bSize - 1

	for i := 0; i < length; i++ {
		// Create pseudo input
		input := make([]byte, block*bSize-1)
		for i := range input {
			input[i] = 'A'
		}

		// Leave out gaps to figure out character
		b, _ := oracle(input[:curSize], key)
		for j := 0; j < len(output); j++ {
			input[len(input)-len(output)+j] = output[j]
		}

		// Now loop through all 128 ASCII values and find that character
		for val := 0; val < 128; val++ {
			test := append(input, byte(val))
			c, _ := oracle(test, key)

			if bytes.Compare(b[:(block)*bSize], c[:(block)*bSize]) == 0 {
				output = append(output, byte(val))
				break
			}
		}

		// If curSize becomes 0, that means we move to the next block
		if curSize%bSize == 0 {
			block++
			curSize = bSize - 1
		} else {
			curSize--
		}

	}

	return output, nil
}
