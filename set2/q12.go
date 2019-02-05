package set2

import (
	"errors"
	"fmt"
	"matasano/cipher"
	"matasano/types"
	"math/rand"
)

const (
	blockSizeBytes   = 16
	unkownTextString = `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK`
)

// Copied from cipher/helpers/EncryptionOracle
func fixedECBOracle(input []byte, key []byte) ([]byte, error) {
	var b types.Base64
	b.Set([]byte(unkownTextString))
	unkownText, err := b.Decode()
	if err != nil {
		return []byte{}, err
	}

	data := append(input, unkownText...)

	var p types.PKCS7
	p.SetBlockSize(blockSizeBytes)
	err = p.Encode(data)
	if err != nil {
		return []byte{}, err
	}

	var t types.Text
	t.Set(p.B)

	var e cipher.ECB
	e.Init(key, blockSizeBytes*8)
	if err != nil {
		return []byte{}, err
	}
	encoded, err := e.Encode(&t)
	if err != nil {
		return []byte{}, err
	}

	return encoded, nil
}

func discoverBlockSize(key []byte) int {
	prevLen := 0
	bSize := 0
	for i := 1; i < 35; i++ {
		b := make([]byte, i)
		for _, j := range b {
			b[j] = 'A'
		}
		op, _ := fixedECBOracle(b, key)
		if prevLen > 0 {
			curLen := len(op)
			if curLen > prevLen {
				diff := curLen - prevLen
				if diff < bSize || bSize == 0 {
					bSize = diff
				}
			}
		}
		prevLen = len(op)
	}
	return bSize
}

func SolveQ12() error {
	key := make([]byte, blockSizeBytes)
	rand.Read(key)
	bSize := discoverBlockSize(key)

	check, err := cipher.DetectECB(fixedECBOracle)
	if err != nil {
		return err
	} else if check == false {
		return errors.New("not in ECB mode")
	}

	d, err := cipher.BreakECB(fixedECBOracle, bSize)
	if err != nil {
		return err
	}
	fmt.Println(string(d))

	return nil
}