package cipher

import (
	"encoding/binary"
	"fmt"
	"log"
	"matasano/helpers"
	"matasano/random"
	"matasano/types"
)

type MT19937 struct {
	m random.MT19937
}

func (m *MT19937) Init(key []byte, d ...interface{}) error {
	if len(key) != 2 {
		return fmt.Errorf("expected length of key: 16bits got: %v", 8*len(key))
	}

	if len(d) != 0 {
		return fmt.Errorf("expected only 2 argument, but got: %v", len(d))
	}

	m.m.Seed(uint32(key[1])<<8 | uint32(key[0]))
	return nil
}

func (m *MT19937) Decode(t types.Type) ([]byte, error) {
	return m.Encode(t)
}

func (m *MT19937) Encode(t types.Type) ([]byte, error) {
	data, err := t.Decode()
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	keystream := make([]byte, len(data))

	for i := 0; i < len(data); i += 4 {
		data := m.m.Uint32()
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, data)
		copy(keystream[i:], bs)
	}

	val, _ := helpers.Xor(&types.Text{T: data}, &types.Text{T: keystream})
	return val.Get(), nil
}
