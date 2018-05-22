package base64

import (
	"strings"
	"testing"
)

func TestFromHex(t *testing.T) {
	tables := []struct {
		input  string
		output string
		err    error
	}{
		{
			"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
			nil,
		},
	}

	for _, table := range tables {
		encoding, err := FromHex(table.input)
		if err != table.err {
			t.Errorf("Received error: %v not equal to the expected output: %s", err, table.err)
		}
		if strings.Compare(encoding, table.output) != 0 {
			t.Errorf("Encoded string: %s not equal to the expected output: %s", encoding, table.output)
		}
	}
}

func TestToHex(t *testing.T) {
	tables := []struct {
		input  string
		output string
		err    error
	}{
		{
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
			"49276D206B696C6C696E6720796F757220627261696E206C696B65206120706F69736F6E6F7573206D757368726F6F6D",
			nil,
		},
	}

	for _, table := range tables {
		encoding, err := ToHex(table.input)
		if err != table.err {
			t.Errorf("Received error: %v not equal to the expected output: %s", err, table.err)
		}
		if strings.Compare(encoding, table.output) != 0 {
			t.Errorf("Encoded string: %s not equal to the expected output: %s", encoding, table.output)
		}
	}
}