package types

type Cipher interface {
	Set(s string)
	Get() string
	Decode() (string, error)
	Encode(b string) (error)
}