package set2

import (
	"errors"
	"fmt"
	"matasano/cipher"
	"matasano/types"
	"math/rand"
)

func wrapperOracleRandomPrefix() cipher.Oracle {
	randomStartSize := rand.Intn(blockSizeBytes)
	startBytes := make([]byte, randomStartSize)
	rand.Read(startBytes)

	return func(input []byte, key []byte) ([]byte, error) {
		var b types.Base64
		b.Set([]byte(unkownTextString))
		unkownText, err := b.Decode()
		if err != nil {
			return []byte{}, err
		}

		data := append(input, unkownText...)
		data = append(startBytes, data...)
		return fixedECBOracle(data, key)
	}
}

func SolveQ14() error {
	key := make([]byte, blockSizeBytes)
	rand.Read(key)

	oracle := wrapperOracleRandomPrefix()

	bSize := cipher.DiscoverBlockSize(oracle, key)

	check, err := cipher.DetectECB(oracle)
	if err != nil {
		return err
	} else if check == false {
		return errors.New("not in ECB mode")
	}

	value, err := cipher.BreakECBRandomPrefix(oracle, bSize)
	if err != nil {
		return err
	}

	fmt.Println(string(value))

	return nil
}
