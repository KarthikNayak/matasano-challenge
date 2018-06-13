package types

type Cipher interface {
	Set(b []byte) (Cipher)
	Get() []byte
	Decode() ([]byte, error)
	Encode(b []byte) (error)
}
