package types

type Type interface {
	Set(b []byte) (Type)
	Get() []byte
	Decode() ([]byte, error)
	Encode(b []byte) (error)
}
