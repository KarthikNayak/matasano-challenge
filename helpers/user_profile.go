package helpers

import (
	"fmt"
)

var metaChar = [2]byte{'&', '='}

func ProfileForEncoding(email string) string {
	m := make(map[string]string)

	var finalEmail string
	for _, val := range email {
		isMC := false
		for _, mc := range metaChar {
			if rune(mc) == val {
				isMC = true
				break
			}
		}

		if !isMC {
			finalEmail += string(val)
		}
	}
	m["email"] = finalEmail
	m["uid"] = "10"
	m["role"] = "user"

	var output string
	for key, val := range m {
		if len(output) > 0 {
			output += "&"
		}
		output += fmt.Sprintf("%s=%s", key, val)

	}
	return output
}
