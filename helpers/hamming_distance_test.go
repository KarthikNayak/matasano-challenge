package helpers

import "testing"

func TestHammingDistance(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"test hamming distance",
			args{a: []byte("this is a test"), b: []byte("wokka wokka!!!")},
			37,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HammingDistance(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("HammingDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
