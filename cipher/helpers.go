package cipher

import (
	"bytes"
	"fmt"
	"matasano/types"
	"math/rand"
)

type Oracle func([]byte, []byte) ([]byte, error)

func EncryptionOracle(input []byte, key []byte) ([]byte, error) {
	blockSizeBytes := 16

	randomStartSize := rand.Intn(6) + 5
	randomEndSize := rand.Intn(6) + 5

	startBytes := make([]byte, randomStartSize)
	endBytes := make([]byte, randomEndSize)

	data := append(startBytes, input...)
	data = append(data, endBytes...)

	var p types.PKCS7
	p.SetBlockSize(blockSizeBytes)
	err := p.Encode(data)
	if err != nil {
		return []byte{}, err
	}

	var t types.Text
	t.Set(p.B)

	if rand.Intn(2) == 1 {
		// ECB
		var e ECB
		fmt.Println("ECB Cipher")

		e.Init(key, blockSizeBytes*8)
		if err != nil {
			return []byte{}, err
		}
		encoded, err := e.Encode(&t)
		if err != nil {
			return []byte{}, err
		}

		return encoded, nil

	}
	// CBC
	var c CBC
	fmt.Println("CBC Cipher")

	IV := make([]byte, blockSizeBytes)
	rand.Read(IV)

	err = c.Init(key, blockSizeBytes*8, IV)
	if err != nil {
		return []byte{}, err
	}
	encoded, err := c.Encode(&t)
	if err != nil {
		return []byte{}, err
	}

	return encoded, nil
}

func DetectECB(oracle Oracle) (bool, error) {
	blockSize := 16

	input := make([]byte, 16*5)

	key := make([]byte, blockSize)
	rand.Read(key)
	decoded, err := oracle(input, key)
	if err != nil {
		return false, err
	}

	for i := 0; i < len(decoded); i += blockSize {
		for j := i + blockSize; j < len(decoded); j += blockSize {
			if bytes.Compare(decoded[i:i+blockSize], decoded[j:j+blockSize]) == 0 {
				return true, nil
			}
		}
	}
	return false, nil
}
