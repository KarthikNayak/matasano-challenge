package set2

import (
	"bytes"
	"matasano/cipher"
	"matasano/types"
	"math/rand"
	"testing"
)

func TestCbcEncoder(t *testing.T) {
	table := []struct {
		input    string
		unparsed []byte
	}{
		{
			"hello=world",
			[]byte("comment1\"=\"cooking%20MCs\";\"userdata\"=\"hello\"=\"world\";\"comment2\"=\"%20like%20a%20pound%20of%20bacon"),
		},
	}

	key := make([]byte, blockSizeBytes)
	rand.Read(key)

	for _, test := range table {
		s, _ := cbcEncoder([]byte(test.input), key)

		var txt types.Text
		txt.Set(s)
		IV := make([]byte, blockSizeBytes)

		var c cipher.CBC
		c.Init(key, blockSizeBytes*8, IV)
		s, _ = c.Decode(&txt)

		var p types.PKCS7
		p.SetBlockSize(blockSizeBytes)
		p.B = s
		s, _ = p.Decode()

		if !bytes.Equal(s, test.unparsed) {
			t.Errorf("expected: %s got: %s", test.unparsed, s)
		}
	}
}

func TestDecryptSearchAdmin(t *testing.T) {
	table := []struct {
		data   string
		result bool
	}{
		{
			";admin=true",
			true,
		},
		{
			"foo=true",
			false,
		},
	}

	key := make([]byte, blockSizeBytes)
	rand.Read(key)

	for _, test := range table {
		var p types.PKCS7
		p.SetBlockSize(blockSizeBytes)
		p.Encode([]byte(test.data))

		var txt types.Text
		txt.Set(p.B)

		IV := make([]byte, blockSizeBytes)

		var c cipher.CBC
		c.Init(key, blockSizeBytes*8, IV)
		data, _ := c.Encode(&txt)

		if test.result != decryptSearchAdmin(data, key) {
			t.Errorf("expected: %v got: %v", test.result, !test.result)
		}
	}
}

func TestQ16(t *testing.T) {
	if SolveQ16() != true {
		t.Errorf("expected true")
	}
}
