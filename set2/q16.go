package set2

import (
	"matasano/cipher"
	"matasano/types"
	"math/rand"
	"strings"
)

func cbcEncoder(input, key []byte) ([]byte, error) {
	data := append([]byte("comment1=cooking%20MCs;userdata="), input...)
	data = append(data, []byte(";comment2=%20like%20a%20pound%20of%20bacon")...)

	data = []byte(strings.Replace(string(data), "=", "\"=\"", -1))
	data = []byte(strings.Replace(string(data), ";", "\";\"", -1))

	var p types.PKCS7
	p.SetBlockSize(blockSizeBytes)
	err := p.Encode(data)
	if err != nil {
		return []byte{}, err
	}

	var t types.Text
	t.Set(p.B)

	IV := make([]byte, blockSizeBytes)

	var c cipher.CBC
	c.Init(key, blockSizeBytes*8, IV)
	if err != nil {
		return []byte{}, err
	}
	return c.Encode(&t)
}

func decryptSearchAdmin(data, key []byte) bool {
	var t types.Text
	t.Set(data)
	IV := make([]byte, blockSizeBytes)

	var c cipher.CBC
	c.Init(key, blockSizeBytes*8, IV)
	s, _ := c.Decode(&t)

	return strings.Contains(string(s), ";admin=true")
}

func SolveQ16() bool {
	key := make([]byte, blockSizeBytes)
	rand.Read(key)

	s, _ := cbcEncoder([]byte("?admin?true"), key)

	s[38-blockSizeBytes] = s[38-blockSizeBytes] ^ '?' ^ ';'
	s[44-blockSizeBytes] = s[44-blockSizeBytes] ^ '?' ^ '='
	return decryptSearchAdmin(s, key)

}
