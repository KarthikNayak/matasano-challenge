package cipher

import (
	"bytes"
	"cryptopals/types"
	"io/ioutil"
	"testing"
)

func TestECBDecode(t *testing.T) {
	content, err := ioutil.ReadFile("ecb_test_data")
	if err != nil {
		t.Errorf(err.Error())
	}

	b64 := types.Base64(content)
	b64Decoded, _ := b64.Decode()

	var e ECB
	err = e.Init([]byte("YELLOW SUBMARINE"), 128)
	if err != nil {
		t.Errorf(err.Error())
	}
	_, err = e.Decode(b64Decoded)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestECBEncode(t *testing.T) {
	content, err := ioutil.ReadFile("ecb_test_data")
	if err != nil {
		t.Errorf(err.Error())
	}

	b64 := types.Base64(content)
	original, _ := b64.Decode()

	var e ECB
	err = e.Init([]byte("YELLOW SUBMARINE"), 128)
	if err != nil {
		t.Errorf(err.Error())
	}

	data, err := e.Decode(original)
	if err != nil {
		t.Errorf(err.Error())
	}

	encoded, err := e.Encode(data)
	if !bytes.Equal(original, encoded) {
		t.Errorf("expected the encoded value to match the existing data")
	}
}
