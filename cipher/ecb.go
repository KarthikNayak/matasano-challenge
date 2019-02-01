package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"matasano/types"
)

type ECB struct {
	Block    cipher.Block
	SizeBits int
}

func (e *ECB) Init(key []byte, d ...interface{}) error {
	var err error

	if len(d) == 1 {
		e.SizeBits = d[0].(int)
		e.Block, err = aes.NewCipher(key)
		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("expected only 1 argument, but got: %v", len(d))
}

func (e *ECB) Decode(c types.Type) ([]byte, error) {
	data, err := c.Decode()
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	sizeBytes := e.SizeBits / 8

	for i := 0; i < len(data); i += sizeBytes {
		e.Block.Decrypt([]byte(data[i:i+sizeBytes]), []byte(data[i:i+sizeBytes]))
	}

	return data, nil
}

func (e *ECB) Encode(c types.Type) ([]byte, error) {
	data, err := c.Decode()
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	sizeBytes := e.SizeBits / 8

	for i := 0; i < len(data); i += sizeBytes {
		e.Block.Encrypt([]byte(data[i:i+sizeBytes]), []byte(data[i:i+sizeBytes]))
	}

	return data, nil
}
