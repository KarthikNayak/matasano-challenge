package types

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeHex(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want Hex
	}{
		{
			"Simple Text to hex",
			args{b: []byte("Hello there!")},
			Hex([]byte("48656c6c6f20746865726521")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeHex(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHex_Decode(t *testing.T) {
	tests := []struct {
		name    string
		h       Hex
		want    []byte
		wantErr bool
	}{
		{
			"Simple Text to hex",
			Hex([]byte("48656c6c6f20746865726521")),
			[]byte("Hello there!"),
			false,
		},
		{
			"odd hex array length",
			Hex([]byte("48656c6c6f2074686572652")),
			[]byte(""),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.Decode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Hex.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if bytes.Compare([]byte(got), []byte(tt.want)) != 0 {
				t.Errorf("Hex.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
