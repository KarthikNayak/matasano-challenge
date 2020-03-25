package hack

import (
	"bytes"
	"log"
)

func DiscoverBlockSize(o oracleFunc, key []byte) int {
	cur := 0

	for i := 0; ; i++ {
		b := make([]byte, i)
		for i := range b {
			b[i] = 'A'
		}
		op := o(b, key)

		if cur == 0 {
			cur = len(op)
		}

		if len(op) != cur {
			return len(op) - cur
		}
	}
}

func BreakECB(o oracleFunc, bSize int, key []byte) ([]byte, error) {
	// get the length we would finally need
	length := len(o([]byte{}, key))

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
		b := o(input[:curSize], key)
		for j := 0; j < len(output); j++ {
			input[len(input)-len(output)+j] = output[j]
		}

		// Now loop through all 128 ASCII values and find that character
		for val := 0; val < 128; val++ {
			test := append(input, byte(val))
			c := o(test, key)

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

func BreakIfECB(o oracleFunc) []byte {
	key := random16BitKey()

	bSize := DiscoverBlockSize(o, key)
	log.Println("bSize:", bSize)
	if !DetectECB(o) {
		log.Fatal("not ECB")
	}
	log.Println("Ok it's ECB")

	output, err := BreakECB(o, bSize, key)
	if err != nil {
		log.Fatal(err)
	}
	return output
}
