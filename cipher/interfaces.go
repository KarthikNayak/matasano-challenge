package cipher

import "matasano/types"

type Cipher interface {
	Init(key []byte, d ...interface{}) error
	Decode(c types.Type) ([]byte, error)
	Encode(c types.Type) ([]byte, error)
}
