package cipher

import (
	"bytes"
	"matasano/types"
	"testing"
)

const (
	msg = "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ=="
)

func TestCTRDecode(t *testing.T) {
	b64 := types.Base64{B: []byte(msg)}

	var c CTR
	err := c.Init([]byte("YELLOW SUBMARINE"), 128, uint64(0))
	if err != nil {
		t.Errorf(err.Error())
	}
	_, err = c.Decode(&b64)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestCTREncode(t *testing.T) {
	b64 := types.Base64{B: []byte(msg)}

	var c CTR
	err := c.Init([]byte("YELLOW SUBMARINE"), 128, uint64(0))
	if err != nil {
		t.Errorf(err.Error())
	}
	data, err := c.Decode(&b64)
	if err != nil {
		t.Errorf(err.Error())
	}

	encoded, err := c.Encode(&types.Text{T: data})
	if err != nil {
		t.Errorf(err.Error())
	}

	b64Decoded, _ := b64.Decode()

	if bytes.Compare(b64Decoded, encoded) != 0 {
		t.Errorf("encoded string not matching expected value")
	}
}
