package oracle

import (
	"cryptopals/cipher"
	"cryptopals/padding"
	"cryptopals/types"
)

const (
	BlockSizeBytes   = 16
	unkownTextString = `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK`
)

func FixedECBOracle(data []byte, key []byte) []byte {
	data = padding.PKCS7(data, BlockSizeBytes)

	var e cipher.ECB
	e.Init(key, BlockSizeBytes*8)
	encoded, _ := e.Encode(data)
	return encoded
}

func FixedECBWithUnkownText(input []byte, key []byte) []byte {
	b := types.Base64(unkownTextString)

	unkownText, _ := b.Decode()

	data := append(input, unkownText...)
	return FixedECBOracle(data, key)
}
