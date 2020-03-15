package types

import (
	"reflect"
	"testing"
)

func TestBase64_Decode(t *testing.T) {
	tests := []struct {
		name    string
		b64     Base64
		want    []byte
		wantErr bool
	}{
		{
			"Test 1",
			Base64([]byte("SGVsbG8gdGhlcmUh")),
			[]byte("Hello there!"),
			false,
		},
		{
			"Test 2",
			Base64([]byte("YW55IGNhcm5hbCBwbGVhcw==")),
			[]byte("any carnal pleas"),
			false,
		},
		{
			"Test 3",
			Base64([]byte("YW55IGNhcm5hbCBwbGVhc3U=")),
			[]byte("any carnal pleasu"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b64.Decode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base64.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeBase64(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want Base64
	}{
		{
			"Simple Text to hex",
			args{b: []byte("Hello there!")},
			Base64([]byte("SGVsbG8gdGhlcmUh")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeBase64(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
