package helpers

import (
	"reflect"

	"matasano/types"
)

func Xor(a, b types.Type) (types.Type, error) {
	cipherType := reflect.ValueOf(a)
	output := reflect.New(reflect.Indirect(cipherType).Type()).Interface().(types.Type)

	aDecoded, err := a.Decode()
	if err != nil {
		return output, err
	}
	bDecoded, err := b.Decode()
	if err != nil {
		return output, err
	}

	l := len(aDecoded)
	if l > len(bDecoded) {
		l = len(bDecoded)
	}

	outputBytes := make([]byte, l)

	for i := 0; i < l; i++ {
		outputBytes[i] = aDecoded[i] ^ bDecoded[i]
	}
	output.Encode(outputBytes)
	return output, nil
}
