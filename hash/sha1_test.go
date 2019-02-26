package hash

import (
	"crypto/rand"
	"fmt"
	"strings"
	"testing"
)

func TestPreProcess(t *testing.T) {
	table := []struct {
		len int
	}{
		{1},
		{63},
		{64},
		{66},
	}

	for _, test := range table {
		data := make([]byte, test.len)
		rand.Read(data)

		fmt.Println(data)
		data = PreProcess(data)
		fmt.Println(data)

		if len(data)%64 != 0 {
			t.Errorf("expected the final data to be congruent to 64")
		}

		for i := test.len + 1; i < len(data)-8; i++ {
			if data[i] != 0 {
				t.Errorf("expected zero at position: %v", i)
			}
		}
	}
}

func TestSha1(t *testing.T) {
	tests := []struct {
		data []byte
		sha1 string
	}{
		{
			[]byte("The quick brown fox jumps over the lazy dog"),
			"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12",
		},
		{
			[]byte("The quick brown fox jumps over the lazy cog"),
			"de9f2c7fd25e1b3afad3e85a0bd17d9b100db4b3",
		},
		{
			[]byte(""),
			"da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
	}

	for _, test := range tests {
		sha1 := Sha1(test.data)
		if strings.Compare(sha1, test.sha1) != 0 {
			t.Errorf("expected: %s got: %s\n", test.sha1, sha1)
		}
	}
}
