package set3

import (
	"testing"
)

func TestGetQ19Data(t *testing.T) {
	data, err := getQ19Data()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 40 {
		t.Errorf("expected length: 20 got: %v", len(data))
	}
}

func TestSolveQ19(t *testing.T) {
	_, err := SolveQ19()
	if err != nil {
		t.Error(err)
	}
}
