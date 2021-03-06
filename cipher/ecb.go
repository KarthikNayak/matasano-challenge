package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
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

func (e *ECB) Decode(data []byte) ([]byte, error) {
	sizeBytes := e.SizeBits / 8

	for i := 0; i < len(data); i += sizeBytes {
		e.Block.Decrypt(data[i:i+sizeBytes], data[i:i+sizeBytes])
	}

	return data, nil
}

func (e *ECB) Encode(data []byte) ([]byte, error) {
	sizeBytes := e.SizeBits / 8

	for i := 0; i < len(data); i += sizeBytes {
		e.Block.Encrypt(data[i:i+sizeBytes], data[i:i+sizeBytes])
	}

	return data, nil
}
