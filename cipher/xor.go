package cipher

import (
	"gitlab.com/karthiknayak/matasano/metrics/frequency"
	"gitlab.com/karthiknayak/matasano/types"
)

func SingleByteXOR(c types.Cipher, f frequency.Frequency) (string, float64, error) {
	s, err := c.Decode()
	if err != nil {
		return "", 0.0, err
	}
	maxScore := 0.0
	FinalString := ""
	tmp := make([]byte, len(s))
	for i := 0; i < 255; i++ {
		for j, v := range s {
			tmp[j] = byte(v) ^ byte(i)
		}
		score := f.GetFrequency(tmp)
		if score > maxScore {
			maxScore = score
			FinalString = string(tmp)
		}
	}
	return FinalString, maxScore, nil
}
