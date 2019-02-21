package cipher

import (
	"matasano/types"
	"testing"
)

const (
	msg2 = "Hello, there! Why are you reading this sentence?"
)

func TestMT19937Cipher(t *testing.T) {
	var m1, m2 MT19937

	m1.Init([]byte{12, 34})
	m2.Init([]byte{12, 34})
	b, err := m1.Encode(&types.Text{T: []byte(msg2)})
	if err != nil {
		t.Error(err)
	}
	b2, err := m2.Decode(&types.Text{T: b})
	if err != nil {
		t.Error(err)
	}

	if string(b2) != msg2 {
		t.Errorf("expected: %s got %s", msg2, b2)
	}
}
