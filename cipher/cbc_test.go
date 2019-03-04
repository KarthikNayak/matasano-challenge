package cipher

import (
	"bytes"
	"io/ioutil"
	"matasano/types"
	"testing"
)

func TestCBCDecode(t *testing.T) {
	content, err := ioutil.ReadFile("cbc_test_data")
	if err != nil {
		t.Errorf(err.Error())
	}

	b64 := types.Base64{B: content}

	var c CBC
	b := make([]byte, 16)
	err = c.Init([]byte("YELLOW SUBMARINE"), 128, b)
	if err != nil {
		t.Errorf(err.Error())
	}
	data, err := c.Decode(&b64)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(string(data))
}

func TestCBCEncode(t *testing.T) {
	content, err := ioutil.ReadFile("cbc_test_data")
	if err != nil {
		t.Errorf(err.Error())
	}

	b64 := types.Base64{B: content}
	original, err := b64.Decode()
	if err != nil {
		t.Errorf(err.Error())
	}

	var c CBC
	IV := make([]byte, 16)
	err = c.Init([]byte("YELLOW SUBMARINE"), 128, IV)
	if err != nil {
		t.Errorf(err.Error())
	}

	data, err := c.Decode(&b64)
	if err != nil {
		t.Errorf(err.Error())
	}

	text := types.Text{T: data}
	encoded, err := c.Encode(&text)
	if !bytes.Equal(original, encoded) {
		t.Errorf("expected the encoded value to match the existing data")
	}
}
