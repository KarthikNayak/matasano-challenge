package types

type Bytes []byte

func (b Bytes) Set(s string) (Cipher) {
	b = []byte(s)
	return b
}

func (b Bytes) Get() string {
	return string(b)
}

func (b Bytes) Decode() (string, error) {
	return b.Get(), nil
}

func (b Bytes) Encode(s string) (error) {
	b.Set(string(b))
	return nil
}