package set4

import (
	"bytes"
	"errors"
	"fmt"
	"matasano/cipher"
	"matasano/helpers"
	"matasano/types"
	"math/rand"
)

func cbcEncoderNoIV(input, key []byte) ([]byte, error) {
	var p types.PKCS7
	p.SetBlockSize(bSize)
	err := p.Encode(input)
	if err != nil {
		return []byte{}, err
	}

	var c cipher.CBC
	c.Init(key, bSize*8, key)
	if err != nil {
		return []byte{}, err
	}
	return c.Encode(&types.Text{T: p.B})
}

func TestHighAscii(data, key []byte) ([]byte, error) {
	var c cipher.CBC
	c.Init(key, bSize*8, key)
	s, _ := c.Decode(&types.Text{T: data})
	fmt.Println(s)

	for _, val := range s {
		if val > 127 {
			return s, errors.New("high ASCII")
		}
	}

	return []byte{}, nil
}

func SolveQ27() error {
	key := make([]byte, bSize)
	rand.Read(key)

	data := make([]byte, 3*bSize)
	for i := range data {
		data[i] = 128
	}

	s, _ := cbcEncoderNoIV(data, key)
	for i := bSize; i < 2*bSize; i++ {
		s[i] = 0
	}
	for i := 0; i < bSize; i++ {
		s[2*bSize+i] = s[i]
	}

	b, err := TestHighAscii(s, key)
	if err == nil {
		return errors.New("was expecting error")
	}

	val, _ := helpers.Xor(&types.Text{T: b[:bSize]}, &types.Text{T: b[2*bSize:]})

	if bytes.Compare(val.Get(), key) != 0 {
		return errors.New("incorrect key")
	}

	return nil
}
