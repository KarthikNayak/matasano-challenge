package set3

import (
	"bufio"
	"matasano/cipher"
	"matasano/helpers"
	"matasano/types"
	"os"
)

func getQ19Data() ([][]byte, error) {
	var data [][]byte

	file, err := os.Open("q19_data")
	if err != nil {
		return [][]byte{}, err
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		b64 := types.Base64{B: []byte(sc.Text())}
		b, err := b64.Decode()
		if err != nil {
			return [][]byte{}, err
		}
		data = append(data, b)
	}

	return data, nil
}

func CTRFixedNonceOracle(data, key []byte) ([]byte, error) {
	var p types.PKCS7
	p.SetBlockSize(bSize)
	err := p.Encode(data)
	if err != nil {
		return []byte{}, err
	}

	var c cipher.CTR
	c.Init(key, bSize*8, uint64(0))
	return c.Encode(&types.Text{T: p.Get()})
}

func CTREncryptedData() ([][]byte, int, error) {
	var data [][]byte

	rawData, err := getQ19Data()
	if err != nil {
		return [][]byte{}, 0, err
	}
	biggest := 0

	for _, row := range rawData {
		d, err := CTRFixedNonceOracle(row, []byte("YELLOW SUBMARINE"))
		if err != nil {
			return [][]byte{}, 0, err
		}
		if len(d) > biggest {
			biggest = len(d)
		}
		data = append(data, d)
	}
	return data, biggest, nil
}

func SolveQ19() ([]byte, error) {
	data, size, err := CTREncryptedData()
	if err != nil {
		return []byte{}, err
	}

	keystream := make([]byte, size)
	for i := 0; i < size; i++ {
		var maxScore float64
		var val byte
		for j := 1; j < 256; j++ {
			score := 0.0
			for _, row := range data {
				if len(row) > i {
					score += helpers.CanonicalFreq[row[i]^byte(j)]
				}
			}
			if score >= maxScore {
				maxScore = score
				val = byte(j)
			}
		}
		keystream[i] = val
	}

	for _, row := range data {
		val, _ := helpers.Xor(&types.Text{T: row[:22]}, &types.Text{T: keystream[:22]})
		s := string(val.Get())
		// output string
		_ = s
	}

	return []byte{}, nil
}
