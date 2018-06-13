package types

type Bytes []byte

func (b Bytes) Set(bytes []byte) (Cipher) {
	return b
}

func (b Bytes) Get() []byte {
	return b
}

func (b Bytes) Decode() ([]byte, error) {
	return b.Get(), nil
}

func (b Bytes) Encode(bytes []byte) (error) {
	b.Set(bytes)
	return nil
}