package hack

import (
	"bytes"
	"cryptopals/oracle"
)

func DetectECB() bool {
	data := make([]byte, 16*200)
	for i := 0; i < len(data); i++ {
		data[i] = 'a'
	}

	encrypted, _ := oracle.EncryptionECBCBCOracle(data)

	for i := 0; i < len(encrypted); i += 16 {
		for j := i + 16; j < len(encrypted); j += 16 {
			if bytes.Compare(encrypted[i:i+16], encrypted[j:j+16]) == 0 {
				return true
			}
		}
	}
	return false
}
