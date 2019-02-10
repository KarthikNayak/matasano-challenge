package set3

import (
	"fmt"
	"matasano/cipher"
	"matasano/types"
	"math/rand"
	"time"
)

const (
	bSize = 16
)

var randomData = []string{
	"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
	"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
	"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
	"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
	"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
	"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
	"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
	"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
	"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
	"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
}

func cbcRandomDataEncoder(key, IV []byte) ([]byte, error) {
	_, _, seconds := time.Now().Clock()
	rand.Seed(int64(seconds))
	s := randomData[rand.Intn(len(randomData))]

	b64 := types.Base64{B: []byte(s)}
	b64Decoded, err := b64.Decode()
	if err != nil {
		return []byte{}, err
	}
	fmt.Println("expected:", string(b64Decoded))

	var p types.PKCS7
	p.SetBlockSize(bSize)
	p.Encode(b64Decoded)

	var c cipher.CBC
	c.Init(key, bSize*8, IV)
	return c.Encode(&types.Text{T: p.B})
}

func checkEncodedPadding(data, key, IV []byte) bool {
	var t types.Text
	t.Set(data)

	var c cipher.CBC
	c.Init(key, bSize*8, IV)
	data, _ = c.Decode(&t)

	var p types.PKCS7
	p.SetBlockSize(bSize)
	p.B = data
	_, err := p.Decode()
	if err != nil {
		return false
	}
	return true
}

func wrapCheckPadding(key []byte) func(data, IV []byte) bool {
	return func(data, IV []byte) bool {
		return checkEncodedPadding(data, key, IV)
	}
}

func Solve17() error {
	key := make([]byte, bSize)
	rand.Read(key)

	IV := make([]byte, bSize)
	rand.Read(IV)

	x, err := cbcRandomDataEncoder(key, IV)
	if err != nil {
		return err
	}

	f := wrapCheckPadding(key)
	data, err := cipher.PaddingOracleBreak(bSize, x, IV, f)
	fmt.Println(string(data))
	return err
}
