package set4

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	go server()

	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{
			"bad request",
			"http://localhost:8000/",
			400,
		},
		{
			"good request",
			"http://localhost:8000/?file=foo&signature=goo",
			500,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := http.Get(test.url)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != test.statusCode {
				t.Fatalf("expected %v got %v", test.statusCode, resp.StatusCode)
			}
		})
	}
}

func TestInsecureCompare(t *testing.T) {
	tests := []struct {
		sha1       string
		sha2       string
		correctLen int
	}{
		{
			"fbdb1d1b18aa6c08324b7d64b71fb76370690e1d",
			"0bdb1d1b18aa6c08324b7d64b71fb76370690e1d",
			0,
		},
		{
			"fbdb1d1b18aa6c08324b7d64b71fb76370690e1d",
			"fb23412341234123412341231231231231231231",
			1,
		},
		{
			"fbdb1d1b18aa6c08324b7d64b71fb76370690e1d",
			"fbdb412341234123412341231231231231231231",
			2,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("test with %d correct values", test.correctLen), func(t *testing.T) {
			var tDur float64
			loop := 5
			for i := 0; i < loop; i++ {
				startTime := time.Now()
				insecureCompare(test.sha1, test.sha2)
				duration := time.Now().Sub(startTime)
				tDur += duration.Seconds() * 1000
			}
			//normalize
			tDur /= float64(loop)
			if tDur < float64(test.correctLen*50) || tDur > float64((test.correctLen+1)*50) {
				t.Fatalf("time took: %v", tDur)
			}
		})
	}
}

func TestSolve31(t *testing.T) {
	t.Skip("takes too long")

	go server()

	_, err := Solve31()
	if err != nil {
		t.Fatal(err)
	}
}
