package set4

import (
	"fmt"
	"matasano/cipher"
	"matasano/types"
	"math/rand"
	"strings"
)

const (
	bSize = 16
)

func cbcEncoder(input, key []byte) ([]byte, error) {
	data := append([]byte("comment1=cooking%20MCs;userdata="), input...)
	data = append(data, []byte(";comment2=%20like%20a%20pound%20of%20bacon")...)

	data = []byte(strings.Replace(string(data), "=", "\"=\"", -1))
	data = []byte(strings.Replace(string(data), ";", "\";\"", -1))

	var p types.PKCS7
	p.SetBlockSize(bSize)
	err := p.Encode(data)
	if err != nil {
		return []byte{}, err
	}

	var t types.Text
	t.Set(p.B)

	var c cipher.CTR
	c.Init(key, bSize*8, uint64(0))
	if err != nil {
		return []byte{}, err
	}
	return c.Encode(&t)
}

func decryptSearchAdmin(data, key []byte) bool {
	var c cipher.CTR
	c.Init(key, bSize*8, uint64(0))
	s, _ := c.Decode(&types.Text{T: data})

	fmt.Println(string(s))
	return strings.Contains(string(s), ";admin=true")
}

func SolveQ26() bool {
	key := make([]byte, bSize)
	rand.Read(key)

	s, _ := cbcEncoder([]byte("?admin?true"), key)
	s[38] = s[38] ^ '?' ^ ';'
	s[44] = s[44] ^ '?' ^ '='
	return decryptSearchAdmin(s, key)

}
