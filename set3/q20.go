package set3

import (
	"bufio"
	"matasano/cipher"
	"matasano/types"
	"os"
)

func getQ20Data() ([][]byte, error) {
	var data [][]byte

	file, err := os.Open("q20_data")
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

func CTREncryptedDataQ20() ([][]byte, int, error) {
	var data [][]byte

	rawData, err := getQ19Data()
	if err != nil {
		return [][]byte{}, 0, err
	}
	smallest := 10000

	for _, row := range rawData {
		d, err := CTRFixedNonceOracle(row, []byte("YELLOW SUBMARINE"))
		if err != nil {
			return [][]byte{}, 0, err
		}
		if len(d) < smallest {
			smallest = len(d)
		}
		data = append(data, d)
	}
	return data, smallest, nil
}

func SolveQ20() (string, error) {
	data, smallest, err := CTREncryptedDataQ20()
	if err != nil {
		return "", err
	}

	smallest = 16

	var conc []byte
	for _, row := range data {
		conc = append(conc, row[:smallest]...)
	}

	return cipher.DecodeRepeatingXor(&types.Text{T: conc})
}
