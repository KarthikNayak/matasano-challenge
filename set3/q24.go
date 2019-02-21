package set3

import (
	"bytes"
	"errors"
	"matasano/cipher"
	"matasano/random"
	"matasano/types"
	"math/rand"
	"strings"
	"time"
)

func MT19937EncodeWithRandomPrefix(key, data []byte) ([]byte, error) {
	var m1 cipher.MT19937

	err := m1.Init(key)
	if err != nil {
		return []byte{}, err
	}

	randSize := rand.Intn(5) + 5
	prefix := make([]byte, randSize)
	rand.Read(prefix)

	data = append(prefix, data...)
	return m1.Encode(&types.Text{T: data})
}

func Find16BitKey() error {
	data := "AAAAAAAAAAAAAA"
	keyToSolve := []byte{12, 34}
	enc, err := MT19937EncodeWithRandomPrefix(keyToSolve, []byte(data))
	if err != nil {
		return err
	}

	for i := 0; i < 0xff; i++ {
		for j := 0; j < 0xff; j++ {
			var m cipher.MT19937
			err := m.Init([]byte{byte(i), byte(j)})
			if err != nil {
				return err
			}

			dec, err := m.Decode(&types.Text{T: enc})
			if err != nil {
				return err
			}
			if strings.Contains(string(dec), data) {
				return nil
			}
		}
	}

	return errors.New("couldn't find the key")
}

func Generate16BitToken() []byte {
	var m random.MT19937
	m.Seed(uint32(time.Now().Unix()))

	token := make([]byte, 16)
	for i := 0; i < len(token); i += 4 {
		x := m.Uint32()
		token[i] = byte(x)
		token[i+1] = byte(x >> 8)
		token[i+2] = byte(x >> 16)
		token[i+3] = byte(x >> 24)
	}
	return token
}

func CheckGeneratedToken(ptoken []byte) bool {
	for delta := 0; delta < 60*60*24; delta++ {
		var m random.MT19937
		m.Seed(uint32(time.Now().Unix()))

		token := make([]byte, 16)
		for i := 0; i < len(token); i += 4 {
			x := m.Uint32()
			token[i] = byte(x)
			token[i+1] = byte(x >> 8)
			token[i+2] = byte(x >> 16)
			token[i+3] = byte(x >> 24)
		}
		if bytes.Compare(token, ptoken) == 0 {
			return true
		}
	}
	return false
}
