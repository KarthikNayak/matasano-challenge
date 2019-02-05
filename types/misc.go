package types

import (
	"strings"
)

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
