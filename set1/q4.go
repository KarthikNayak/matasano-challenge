package set1

import (
	"bufio"
	"cryptopals/cipher"
	"cryptopals/types"
)

func Q4(scanner *bufio.Scanner) (string, error) {
	maxScore := 0.0
	finalText := ""

	type result struct {
		text  string
		score float64
	}

	c := make(chan result)
	count := 0

	for scanner.Scan() {
		h := types.Hex([]byte(scanner.Text()))
		count += 1

		go func(h types.Hex) {
			b, _ := h.Decode()
			score, ans := cipher.DecodeSingleByteXOR(b)
			c <- result{text: string(ans), score: score}
		}(h)
	}

	for i := 0; i < count; i++ {
		r := <-c
		if r.score > maxScore {
			maxScore = r.score
			finalText = string(r.text)
		}
	}
	return finalText, nil
}
