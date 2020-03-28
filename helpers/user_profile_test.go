package helpers

import (
	"reflect"
	"testing"
)

func TestProfileForEncoding(t *testing.T) {
	table := []struct {
		email string
		m     map[string]string
	}{
		{
			email: "foo@bar.com",
			m: map[string]string{
				"email": "foo@bar.com",
				"uid":   "10",
				"role":  "user",
			},
		},
		{
			email: "foo@bar.com&role=admin",
			m: map[string]string{
				"email": "foo@bar.comroleadmin",
				"uid":   "10",
				"role":  "user",
			},
		},
	}

	for _, test := range table {
		output := ProfileForEncoding(test.email)
		m := JSONIsh(output)

		if !reflect.DeepEqual(m, test.m) {
			t.Errorf("expected: %v got: %v", test.m, m)
		}
	}
}
