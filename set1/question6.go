package set1

import (
	"io/ioutil"
	"log"

	"gitlab.com/karthiknayak/matasano/cipher"
	"gitlab.com/karthiknayak/matasano/types"
)

func SolveQ6(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	b := types.Base64{B: content}
	return cipher.DecodeRepeatingXor(&b)
}
