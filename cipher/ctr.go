package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"matasano/helpers"
	"matasano/types"
)

type CTR struct {
	Block    cipher.Block
	SizeBits int
	Nonce    []byte
}

func (c *CTR) Init(key []byte, d ...interface{}) error {
	var err error

	if len(d) == 2 {
		c.SizeBits = d[0].(int)
		c.Block, err = aes.NewCipher(key)
		if err != nil {
			return err
		}

		c.Nonce = helpers.IntToLittleEndianBytes(d[1].(uint64))
		return nil
	}

	return fmt.Errorf("expected only 2 argument, but got: %v", len(d))
}

func (c *CTR) Decode(t types.Type) ([]byte, error) {
	return c.Encode(t)
}

func (c *CTR) Encode(t types.Type) ([]byte, error) {
	sizeBytes := c.SizeBits / 8

	data, err := t.Decode()
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	xor := make([]byte, len(data))
	count := 0

	for i := 0; i < len(data); i += sizeBytes {
		counter := helpers.IntToLittleEndianBytes(uint64(count))

		keystream := make([]byte, sizeBytes)
		tmp := append(c.Nonce, counter...)

		c.Block.Encrypt(keystream, tmp)

		copy(xor[i:], keystream)
		count++
	}

	val, _ := helpers.Xor(&types.Text{T: data}, &types.Text{T: xor})
	return val.Get(), nil
}
