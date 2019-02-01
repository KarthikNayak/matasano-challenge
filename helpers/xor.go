package helpers

import (
	"reflect"

	"errors"
	"matasano/types"
)

func Xor(a, b types.Type) (types.Type, error) {
	cipherType := reflect.ValueOf(a)
	output := reflect.New(reflect.Indirect(cipherType).Type()).Interface().(types.Type)

	if len(a.Get()) != len(b.Get()) {
		return output, errors.New("The two ciphers have different length")
	}

	aDecoded, err := a.Decode()
	if err != nil {
		return output, err
	}
	bDecoded, err := b.Decode()
	if err != nil {
		return output, err
	}

	outputBytes := make([]byte, len(aDecoded))
	for i := range aDecoded {
		outputBytes[i] = aDecoded[i] ^ bDecoded[i]
	}
	output.Encode(outputBytes)
	return output, nil
}
