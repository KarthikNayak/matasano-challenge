package cipher

import (
	"reflect"

	"gitlab.com/karthiknayak/matasano/metrics/frequency"
	"gitlab.com/karthiknayak/matasano/types"
)

func DecodeSingleByteXOR(c types.Cipher, f frequency.Frequency) (string, float64, error) {
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

func RepeatingXorEncode(c types.Cipher, key string) (string, error) {
	cipherType := reflect.ValueOf(c)
	output := reflect.New(reflect.Indirect(cipherType).Type()).Interface().(types.Cipher)

	s, err := c.Decode()
	if err != nil {
		return "", err
	}

	outputBytes := make([]byte, len(s))
	keyLength := len(key)
	for i, val := range s {
		outputBytes[i] = byte(val) ^ key[int(i) % keyLength]
	}
	output.Encode(string(outputBytes))
	return output.Get(), nil
}