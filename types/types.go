package types

type Type interface {
	Decode() ([]byte, error)
}
