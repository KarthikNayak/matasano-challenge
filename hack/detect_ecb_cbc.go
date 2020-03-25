package hack

import (
	"bytes"
	"math/rand"
)

type oracleFunc func([]byte, []byte) []byte

func random16BitKey() []byte {
	b := make([]byte, 16)
	nR := rand.New(rand.NewSource(13))
	nR.Read(b)

	return b
}

func DetectECB(o oracleFunc) bool {
	data := make([]byte, 16*200)
	for i := 0; i < len(data); i++ {
		data[i] = 'a'
	}

	encrypted := o(data, random16BitKey())

	for i := 0; i < len(encrypted); i += 16 {
		for j := i + 16; j < len(encrypted); j += 16 {
			if bytes.Compare(encrypted[i:i+16], encrypted[j:j+16]) == 0 {
				return true
			}
		}
	}
	return false
}
