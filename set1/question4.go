package set1

import (
	"bufio"

	"matasano/cipher"
	"matasano/metrics/frequency"
	"matasano/types"
)

func SolveQ4(scanner *bufio.Scanner) (string, error) {
	maxScore := 0.0
	finalText := ""
	for scanner.Scan() {
		h := types.Hex{B: []byte(scanner.Text())}
		f := frequency.CharacterFrequency{}
		s, score, err := cipher.DecodeSingleByteXOR(&h, &f)
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
