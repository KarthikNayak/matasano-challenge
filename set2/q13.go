package set2

import (
	"cryptopals/cipher"
	"cryptopals/helpers"
	"cryptopals/oracle"
	"math/rand"
)

func helperEmail() []byte {
	// The email should be such that when the formatted string is created
	// "role=" should end at block size bytes
	l := (len("email=") + len("&uid=10&role="))
	l = (oracle.BlockSizeBytes - (l % oracle.BlockSizeBytes))

	m := make([]byte, l)
	for i := range m {
		m[i] = 'h'
	}
	return m
}

func maliciousEmail() []byte {
	// We create a fake block in the middle which can be then appended to any
	// block which ends with "role=" and it'll work
	padLen := oracle.BlockSizeBytes - len("admin")
	m := make([]byte, oracle.BlockSizeBytes)

	for i, val := range "admin" {
		m[i] = byte(val)
	}
	for i := len("admin"); i < oracle.BlockSizeBytes; i++ {
		m[i] = byte(padLen)
	}

	p := make([]byte, oracle.BlockSizeBytes-len("email="))
	for i := range p {
		p[i] = byte('m')
	}

	return append(p, m...)
}

func Q13() string {
	key := make([]byte, oracle.BlockSizeBytes)
	rand.Read(key)

	hEmail := helperEmail()
	hData := helpers.ProfileForEncoding(string(hEmail))
	hECB := oracle.FixedECBOracle([]byte(hData), key)

	mEmail := maliciousEmail()
	mData := helpers.ProfileForEncoding(string(mEmail))
	mECB := oracle.FixedECBOracle([]byte(mData), key)

	extractedAdmin := mECB[16:32]
	l := len(hECB) - oracle.BlockSizeBytes
	finalECB := append(hECB[:l], extractedAdmin...)

	var e cipher.ECB
	e.Init(key, 8*oracle.BlockSizeBytes)
	val, _ := e.Decode(finalECB)

	return string(val)
}
