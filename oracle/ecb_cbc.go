package oracle

import (
	"cryptopals/cipher"
	"cryptopals/padding"
	"fmt"
	"math/rand"
	"time"
)

func random16BitKey() []byte {
	b := make([]byte, 16)
	nR := rand.New(rand.NewSource(13))
	nR.Read(b)

	return b
}

func EncryptionECBCBCOracle(input []byte) ([]byte, error) {
	blockSizeBytes := 16

	rand.Seed(time.Now().UTC().UnixNano())

	randomStartSize := rand.Intn(6) + 5
	randomEndSize := rand.Intn(6) + 5

	startBytes := make([]byte, randomStartSize)
	endBytes := make([]byte, randomEndSize)

	data := append(startBytes, input...)
	data = append(data, endBytes...)

	paddedData := padding.PKCS7(data, blockSizeBytes)

	key := random16BitKey()

	r := rand.Intn(2)
	fmt.Println("rand number: ", r)

	if r == 1 {
		fmt.Println("ECB mode")
		var e cipher.ECB
		e.Init(key, blockSizeBytes*8)

		encoded, err := e.Encode(paddedData)
		if err != nil {
			return []byte{}, err
		}

		return encoded, nil
	}

	fmt.Println("CBC mode")

	var c cipher.CBC

	IV := make([]byte, blockSizeBytes)
	rand.Read(IV)

	c.Init(key, blockSizeBytes*8, IV)

	encoded, err := c.Encode(paddedData)
	if err != nil {
		return []byte{}, err
	}

	return encoded, nil
}
