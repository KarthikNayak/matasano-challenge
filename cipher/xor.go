package cipher

import (
	"reflect"
	"sort"

	"matasano/helpers"
	"matasano/types"
)

const KeySizeMin = 2
const KeySizeMax = 40
const Blocks = 10

type kv struct {
	keysize   int
	hdistance float64
}

func getKeySizesSorted(b []byte) ([]kv, error) {
	var data []kv

	for keysize := KeySizeMin; keysize <= KeySizeMax; keysize++ {
		blocks := make([][]byte, Blocks)
		for i := range blocks {
			blocks[i] = make([]byte, keysize)
			copy(blocks[i], b[i*keysize:i*keysize+keysize])
		}
		hDist := 0.0
		for i := 0; i < Blocks-1; i++ {
			dist, err := helpers.HammingDistance(types.Bytes(blocks[i]), types.Bytes(blocks[i+1]))
			if err != nil {
				return data, err
			}
			hDist += float64(dist)
		}
		hDist = float64(hDist) / (float64(keysize) * float64(Blocks-1))
		data = append(data, kv{keysize, hDist})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].hdistance < data[j].hdistance
	})
	return data, nil
}

func createBuckets(b []byte, size int) [][]byte {
	buckets := make([][]byte, size)
	for i := 0; i < size; i++ {
		buckets[i] = make([]byte, len(b)/size+1)
	}
	for i, val := range b {
		buckets[i%size][i/size] = val
	}
	return buckets
}

func solvedStrings(buckets [][]byte, size int) []string {
	s := make([]string, size)
	for i, val := range buckets {
		s[i], _, _ = DecodeSingleByteXOR(types.Bytes(val))
	}
	return s
}

func DecodeRepeatingXor(c types.Type) (string, error) {
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

func DecodeSingleByteXOR(c types.Type) (string, float64, error) {
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
		score := helpers.CharacterFrequency(tmp)
		if score > maxScore {
			maxScore = score
			FinalString = string(tmp)
		}
	}
	return FinalString, maxScore, nil
}

func EncodeRepeatingXor(c types.Type, key []byte) ([]byte, error) {
	cipherType := reflect.ValueOf(c)
	output := reflect.New(reflect.Indirect(cipherType).Type()).Interface().(types.Type)

	s, err := c.Decode()
	if err != nil {
		return []byte{}, err
	}

	outputBytes := make([]byte, len(s))
	keyLength := len(key)
	for i, val := range s {
		outputBytes[i] = byte(val) ^ key[int(i)%keyLength]
	}
	output.Encode(outputBytes)
	return output.Get(), nil
}
