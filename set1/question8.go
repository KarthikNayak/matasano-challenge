package set1

import (
	"bufio"
	"bytes"
	"log"
	"os"

	"matasano/types"
)

func SolveQ8(fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		count++
		h := types.Hex{B: []byte(scanner.Text())}
		decoded, err := h.Decode()
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(decoded); i += 16 {
			for j := i + 16; j < len(decoded); j += 16 {
				if bytes.Compare(decoded[i:i+16], decoded[j:j+16]) == 0 {
					return true
				}
			}
		}
	}
	return false
}
