package cipher

import (
	"cryptopals/helpers"
	"sort"
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

const KeySizeMin = 2
const KeySizeMax = 40
const Blocks = 10

type kv struct {
	keysize   int
	hdistance float64
}

func getKeySizesSorted(b []byte) []kv {
	var data []kv

	c := make(chan kv)

	for keysize := KeySizeMin; keysize <= KeySizeMax; keysize++ {
		go func(keysize int, c chan<- kv) {
			blocks := make([][]byte, Blocks)
			for i := range blocks {
				blocks[i] = make([]byte, keysize)
				copy(blocks[i], b[i*keysize:i*keysize+keysize])
			}
			hDist := 0.0
			for i := 0; i < Blocks-1; i++ {
				dist, _ := helpers.HammingDistance(blocks[i], blocks[i+1])
				hDist += float64(dist)
			}
			hDist = float64(hDist) / (float64(keysize) * float64(Blocks-1))
			c <- kv{keysize, hDist}
		}(keysize, c)
	}

	for i := KeySizeMin; i <= KeySizeMax; i++ {
		data = append(data, <-c)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].hdistance < data[j].hdistance
	})
	return data
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
		_, tmp := DecodeSingleByteXOR(val)
		s[i] = string(tmp)
	}
	return s
}

func DecodeRepeatingXor(data []byte) (string, error) {
	distances := getKeySizesSorted(data)
	bestDistance := distances[0].keysize

	buckets := createBuckets(data, bestDistance)
	strings := solvedStrings(buckets, bestDistance)

	finalString := ""
	for i := 0; i < len(data); i++ {
		finalString += string(strings[i%bestDistance][i/bestDistance])
	}

	return finalString, nil
}
