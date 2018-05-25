package set1

import (
	"gitlab.com/karthiknayak/matasano/types"
)

func SolveQ1(h string) (string, error){
	hex := types.Hex{S: h}
	b, err := hex.Decode()
	if err != nil {
		return "", err
	}

	b64 := types.Base64{}
	b64.Encode(b)
	return b64.S, nil
}