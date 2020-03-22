package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

type CBC struct {
	Block    cipher.Block
	SizeBits int
	IV       []byte
}

func (c *CBC) Init(key []byte, d ...interface{}) error {
	var err error
	if len(d) == 2 {
		c.SizeBits = d[0].(int)
		c.Block, err = aes.NewCipher(key)
		if err != nil {
			return err
		}

		c.IV = d[1].([]byte)
		if len(c.IV) != c.SizeBits/8 {
			return errors.New("iv size and sizebits don't match")
		}
		return nil
	}
	return fmt.Errorf("expected only 1 argument, but got: %v", len(d))
}

func (c *CBC) Decode(b []byte) ([]byte, error) {
	sizeBytes := c.SizeBits / 8

	data := make([]byte, len(b))

	prevBlock := make([]byte, sizeBytes)
	copy(prevBlock, c.IV)

	for i := 0; i < len(data); i += sizeBytes {
		tmp := make([]byte, sizeBytes)
		copy(tmp, b[i:i+sizeBytes])

		c.Block.Decrypt([]byte(data[i:i+sizeBytes]), b[i:i+sizeBytes])
		for j, val := range data[i : i+sizeBytes] {
			data[i+j] = val ^ prevBlock[j]
		}

		copy(prevBlock, tmp)
	}
	return data, nil
}

func (c *CBC) Encode(b []byte) ([]byte, error) {
	sizeBytes := c.SizeBits / 8
	data := make([]byte, len(b))

	prevBlock := make([]byte, sizeBytes)
	copy(prevBlock, c.IV)

	for i := 0; i < len(b); i += sizeBytes {
		for j, val := range b[i : i+sizeBytes] {
			data[i+j] = val ^ prevBlock[j]
		}

		c.Block.Encrypt([]byte(data[i:i+sizeBytes]), []byte(data[i:i+sizeBytes]))
		copy(prevBlock, data[i:i+sizeBytes])
	}
	return data, nil
}
