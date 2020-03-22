package padding

import (
	"strings"
	"testing"
)

func TestPKCS7(t *testing.T) {
	f := PKCS7([]byte("YELLOW SUBMARINE"), 20)
	if !strings.Contains(string(f), "YELLOW SUBMARINE") {
		t.Fatalf("missing original string")
	}
	for i := 16; i < 20; i++ {
		if f[i] != byte(4) {
			t.Fatalf("missing the padding at the end")
		}
	}
}
