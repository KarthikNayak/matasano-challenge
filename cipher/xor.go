package cipher

import (
	"reflect"
	"sort"

	"gitlab.com/karthiknayak/matasano/metrics/frequency"
	"gitlab.com/karthiknayak/matasano/operations"
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

func EncodeRepeatingXor(c types.Cipher, key string) (string, error) {
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


const KEYSIZE_MIN = 2
const KEYSIZE_MAX = 40
const BLOCKS = 10

type kv struct {
	keysize int
	hdistance float64
}

func getKeySizesSorted(s string) ([]kv, error) {
	var data []kv

	for keysize := KEYSIZE_MIN; keysize <= KEYSIZE_MAX; keysize++ {
		blocks := make([][]byte, BLOCKS)
		for i := range blocks {
			blocks[i] = make([]byte, keysize)
			copy(blocks[i], s[i*keysize:i*keysize+keysize])
		}
		hDist := 0.0
		for i := 0; i < BLOCKS - 1; i += 1 {
			dist, err := operations.HammingDistance(types.Bytes(blocks[i]), types.Bytes(blocks[i + 1]))
			if err != nil {
				return data, err
			}
			hDist += float64(dist)
		}
		hDist = float64(hDist) / (float64(keysize) * float64(BLOCKS - 1))
		data = append(data, kv{keysize, hDist})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].hdistance < data[j].hdistance
	})
	return data, nil
}

func createBuckets(s string, size int) ([][]byte) {
	buckets := make([][]byte, size)
	for i := 0; i < size; i++ {
		buckets[i] = make([]byte, len(s)/size + 1)
	}
	for i, val := range s {
		buckets[i%size][i/size] = byte(val)
	}
	return buckets
}

func solvedStrings(buckets [][]byte, size int) []string {
	s := make([]string, size)
	f := frequency.CharacterFrequency{}
	for i, val := range buckets {
		s[i], _, _ = DecodeSingleByteXOR(types.Bytes(val), &f)
	}
	return s
}

func DecodeRepeatingXor(c types.Cipher) (string, error) {
	s, err := c.Decode()
	if err != nil {
		return "", err
	}

	distances, err := getKeySizesSorted(s)
	if err != nil {
		return "", err
	}
	bestDistance := distances[0].keysize

	buckets := createBuckets(s, bestDistance)
	strings := solvedStrings(buckets, bestDistance)

	finalString := ""
	for i := 0; i < len(s); i++ {
		finalString += string(strings[i%bestDistance][i/bestDistance])
	}

	return finalString, nil
}