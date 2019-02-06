package set2

import (
	"fmt"
	"matasano/cipher"
	"matasano/helpers"
	"matasano/types"
	"math/rand"
)

func maliciousEmail() []byte {
	l := blockSizeBytes - len("admin")
	b := make([]byte, blockSizeBytes)
	for i, val := range "admin" {
		b[i] = byte(val)
	}
	for i := len("admin"); i < len(b); i++ {
		b[i] = byte(l)
	}

	p := make([]byte, blockSizeBytes-len("email="))
	for i := range p {
		p[i] = byte('x')
	}
	return append(p, b...)

}

func normalEmail() []byte {
	l := (len("email=") + len("&uid=10&role="))
	l = ((l/blockSizeBytes)+1)*blockSizeBytes - l

	p := make([]byte, l)
	for i := range p {
		p[i] = byte('x')
	}
	return p
}

func SolveQ13() error {
	key := make([]byte, blockSizeBytes)
	rand.Read(key)

	malEmail := maliciousEmail()
	malProfile := helpers.ProfileForEncoding(string(malEmail))
	malECB, _ := fixedECBOracle([]byte(malProfile), key)

	// Aligns to blocksize
	norEmail := normalEmail()
	norProfile := helpers.ProfileForEncoding(string(norEmail))
	norECB, _ := fixedECBOracle([]byte(norProfile), key)

	extractedAdmin := malECB[16:32]
	l := len(norECB) - blockSizeBytes
	finalECB := append(norECB[:l], extractedAdmin...)

	var t types.Text
	t.Set(finalECB)

	var e cipher.ECB
	e.Init(key, blockSizeBytes*8)
	decoded, _ := e.Decode(&t)

	fmt.Println(string(decoded))

	return nil
}
