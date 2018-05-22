package hex

import (
	"bytes"
	stdhex "encoding/hex"
	"testing"
)

func TestHexToByte(t *testing.T) {
	table := []string{
		"AA",
		"12345678",
		"1a2b3c4d",
		"1",
	}

	for _, input := range table {
		ReqOutput, ReqErr := stdhex.DecodeString(input)
		ObtOutput, ObtErr := HexToByte(input)
		if !bytes.Equal(ReqOutput, ObtOutput) {
			t.Errorf("Required Output: %v Obtained Output: %v", ReqOutput, ObtOutput)
		}
		if (ObtErr == nil) != (ReqErr == nil) {
			t.Errorf("Required Error: %v Obtained Error: %v", ReqErr, ObtErr)
		}
	}
}

func TestByteToHex(t *testing.T) {
	table := []string{
		"AA",
		"12345678",
		"1A2B3C4D",
	}

	for _, output := range table {
		input, err := stdhex.DecodeString(output)
		if err != nil {
			t.Errorf("Required creating input from hex: %v", output)
		}
		ObtOutput, _ := ByteToHex(input)
		if ObtOutput != output {
			t.Errorf("Required Output: %v Obtained Output: %v", output, ObtOutput)
		}
	}
}
