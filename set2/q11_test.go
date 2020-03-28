package set2

import (
	"cryptopals/oracle"
	"fmt"
	"testing"
)

func TestDetectECB(t *testing.T) {
	fmt.Println(DetectECB(oracle.EncryptionECBCBCOracle))
}
