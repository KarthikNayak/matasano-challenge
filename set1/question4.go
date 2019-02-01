package set1

import (
	"bufio"

	"matasano/cipher"
	"matasano/types"
)

func SolveQ4(scanner *bufio.Scanner) (string, error) {
	maxScore := 0.0
	finalText := ""
	for scanner.Scan() {
		h := types.Hex{B: []byte(scanner.Text())}
		s, score, err := cipher.DecodeSingleByteXOR(&h)
		if err != nil {
			return "", err
		}
		if score > maxScore {
			maxScore = score
			finalText = s
		}
	}
	return finalText, nil
}
