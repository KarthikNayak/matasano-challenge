package set3

import (
	"matasano/cipher"
	"matasano/types"
	"strings"
	"testing"
)

func TestMT19937EncodeWithRandomPrefix(t *testing.T) {
	msg := "AAAAAAAAAAAAAAAAAAAAAAAA"
	enc, err := MT19937EncodeWithRandomPrefix([]byte{12, 34}, []byte(msg))
	if err != nil {
		t.Error(err)
	}

	var m cipher.MT19937
	m.Init([]byte{12, 34})

	dec, err := m.Decode(&types.Text{T: enc})
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(dec), msg) {
		t.Errorf("final string: %s doesn't contain: %s", dec, msg)
	}
}

func TestFind16BitKey(t *testing.T) {
	err := Find16BitKey()
	if err != nil {
		t.Error(err)
	}
}

func TestCheckGeneratedToken(t *testing.T) {
	token := Generate16BitToken()
	if !CheckGeneratedToken(token) {
		t.Errorf("not generated using current time")
	}
}
