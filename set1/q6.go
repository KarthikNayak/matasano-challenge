package set1

import (
	"cryptopals/cipher"
	"cryptopals/types"
	"io/ioutil"
	"log"
)

func Q6(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := types.Base64(content).Decode()
	return cipher.DecodeRepeatingXor(data)
}
