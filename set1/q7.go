package set1

import (
	"crypto/aes"
	"cryptopals/types"
	"io/ioutil"
	"log"
)

func Q7(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	b64 := types.Base64(content)
	data, err := b64.Decode()
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(data); i += 16 {
		block.Decrypt([]byte(data[i:i+16]), []byte(data[i:i+16]))
	}

	return string(data), nil
}
