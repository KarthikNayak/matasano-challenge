package oracle

import (
	"cryptopals/cipher"
	"cryptopals/padding"
	"fmt"
	"math/rand"
	"time"
)

func EncryptionECBCBCOracle(input []byte, key []byte) []byte {
	blockSizeBytes := 16

	rand.Seed(time.Now().UTC().UnixNano())

	randomStartSize := rand.Intn(6) + 5
	randomEndSize := rand.Intn(6) + 5

	startBytes := make([]byte, randomStartSize)
	endBytes := make([]byte, randomEndSize)

	data := append(startBytes, input...)
	data = append(data, endBytes...)

	paddedData := padding.PKCS7(data, blockSizeBytes)

	r := rand.Intn(2)
	fmt.Println("rand number: ", r)

	if r == 1 {
		fmt.Println("ECB mode")
		var e cipher.ECB
		e.Init(key, blockSizeBytes*8)

		encoded, _ := e.Encode(paddedData)
		return encoded
	}

	fmt.Println("CBC mode")

	var c cipher.CBC

	IV := make([]byte, blockSizeBytes)
	rand.Read(IV)

	c.Init(key, blockSizeBytes*8, IV)

	encoded, _ := c.Encode(paddedData)
	return encoded
}
