package cipher

import (
	"bytes"
	"cryptopals/types"
	"testing"
)

func TestEncodeRepeatingXor(t *testing.T) {
	type args struct {
		data []byte
		key  []byte
	}
	tests := []struct {
		name    string
		args    args
		want    types.Hex
		wantErr bool
	}{
		{
			"test success",
			args{data: []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`), key: []byte("ICE")},
			types.Hex("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeRepeatingXor(tt.args.data, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeRepeatingXor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			wantBytes, _ := tt.want.Decode()
			if bytes.Compare(got, wantBytes) != 0 {
				t.Errorf("EncodeRepeatingXor() = %v, want %v", got, wantBytes)
			}
		})
	}
}
