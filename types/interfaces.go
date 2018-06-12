package types

type Cipher interface {
	Set(s string) (Cipher)
	Get() string
	Decode() (string, error)
	Encode(b string) (error)
}
