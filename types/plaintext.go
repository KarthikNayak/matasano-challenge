package types

type Text struct {
	T []byte
}

func (t *Text) Set(b []byte) Type {
	t.T = b
	return t
}

func (t *Text) Get() []byte {
	return t.T
}

func (t *Text) Decode() ([]byte, error) {
	return t.T, nil
}

func (t *Text) Encode(b []byte) error {
	t.T = b
	return nil
}
