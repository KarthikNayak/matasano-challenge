package helpers

import (
	"fmt"
	"strings"
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

func JSONIsh(s string) map[string]string {
	m := make(map[string]string)

	vals := strings.Split(s, "&")
	for _, row := range vals {
		keyVal := strings.Split(row, "=")
		if len(keyVal) != 2 {
			continue
		}
		m[keyVal[0]] = keyVal[1]
	}
	return m
}
