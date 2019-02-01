package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"log"
	"matasano/types"
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

func (c *CBC) Decode(t types.Type) ([]byte, error) {
	sizeBytes := c.SizeBits / 8

	data, err := t.Decode()
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	prevBlock := make([]byte, sizeBytes)
	copy(prevBlock, c.IV)

	for i := 0; i < len(data); i += sizeBytes {
		tmp := make([]byte, sizeBytes)
		copy(tmp, data[i:i+sizeBytes])

		c.Block.Decrypt([]byte(data[i:i+sizeBytes]), []byte(data[i:i+sizeBytes]))
		for j, val := range data[i : i+sizeBytes] {
			data[i+j] = val ^ prevBlock[j]
		}

		copy(prevBlock, tmp)
	}
	return data, nil
}

func (c *CBC) Encode(t types.Type) ([]byte, error) {
	sizeBytes := c.SizeBits / 8

	data, err := t.Decode()
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	prevBlock := make([]byte, sizeBytes)
	copy(prevBlock, c.IV)

	for i := 0; i < len(data); i += sizeBytes {
		for j, val := range data[i : i+sizeBytes] {
			data[i+j] = val ^ prevBlock[j]
		}

		c.Block.Encrypt([]byte(data[i:i+sizeBytes]), []byte(data[i:i+sizeBytes]))
		copy(prevBlock, data[i:i+sizeBytes])
	}
	return data, nil
}
