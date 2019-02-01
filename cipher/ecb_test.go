package cipher

import (
	"bytes"
	"io/ioutil"
	"matasano/types"
	"testing"
)

func TestECBDecode(t *testing.T) {
	content, err := ioutil.ReadFile("ecb_test_data")
	if err != nil {
		t.Errorf(err.Error())
	}

	b64 := types.Base64{B: content}

	var e ECB
	err = e.Init([]byte("YELLOW SUBMARINE"), 128)
	if err != nil {
		t.Errorf(err.Error())
	}
	_, err = e.Decode(&b64)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestECBEncode(t *testing.T) {
	content, err := ioutil.ReadFile("ecb_test_data")
	if err != nil {
		t.Errorf(err.Error())
	}

	b64 := types.Base64{B: content}
	original, err := b64.Decode()
	if err != nil {
		t.Errorf(err.Error())
	}

	var e ECB
	err = e.Init([]byte("YELLOW SUBMARINE"), 128)
	if err != nil {
		t.Errorf(err.Error())
	}
	data, err := e.Decode(&b64)
	if err != nil {
		t.Errorf(err.Error())
	}

	text := types.Text{T: data}
	encoded, err := e.Encode(&text)
	if !bytes.Equal(original, encoded) {
		t.Errorf("expected the encoded value to match the existing data")
	}
}
