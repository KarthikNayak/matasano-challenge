package cipher

import (
	"cryptopals/helpers"
)

func EncodeRepeatingXor(data []byte, key []byte) ([]byte, error) {
	outputBytes := make([]byte, len(data))
	keyLength := len(key)
	for i, val := range data {
		outputBytes[i] = byte(val) ^ key[int(i)%keyLength]
	}
	return outputBytes, nil
}

func DecodeSingleByteXOR(b []byte) (float64, []byte) {
	type result struct {
		score float64
		key   rune
		val   []byte
	}

	maxScore := 0.0
	res := make([]byte, len(b))
	c := make(chan result)

	for i := 0; i < 255; i++ {
		go func(i int) {
			var r result

			r.val = make([]byte, len(b))
			for j, v := range b {
				r.val[j] = byte(v) ^ byte(i)
			}

			score := helpers.CharacterFrequency(r.val)
			r.score = score
			r.key = rune(i)
			c <- r
		}(i)
	}
	for i := 0; i < 255; i++ {
		r := <-c
		if r.score > maxScore {
			maxScore = r.score
			copy(res, r.val)
		}
	}
	close(c)

	return maxScore, res
}
