package types

import (
	"reflect"
	"testing"
)

func TestJSONIsh(t *testing.T) {
	table := []struct {
		s string
		m map[string]string
	}{
		{
			s: "foo=bar&baz=qux&zap=zazzle",
			m: map[string]string{
				"foo": "bar",
				"baz": "qux",
				"zap": "zazzle",
			},
		},
		{
			s: "foo=",
			m: map[string]string{
				"foo": "",
			},
		},
		{
			s: "foo",
			m: map[string]string{},
		},
		{
			s: "=foo",
			m: map[string]string{
				"": "foo",
			},
		},
	}

	for _, test := range table {
		m := JSONIsh(test.s)
		if !reflect.DeepEqual(m, test.m) {
			t.Errorf("expected: %v got: %v", test.m, m)
		}
	}
}
