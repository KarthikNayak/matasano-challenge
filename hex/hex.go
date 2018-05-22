package hex

import (
	"fmt"

	"github.com/pkg/errors"
)

// 0 -> 48, 1 -> 49 ... 9 -> 57
// a -> 97, b -> 98, c -> 99, d -> 100, e -> 101, f -> 102
// A -> 65, B -> 66, C -> 67, D -> 68, E -> 69, F -> 70


// parseHex returns the absolute value given a hex character.
// parseHex works with upper and lower case characters (a-fA-F).
func parseHex(hex byte) (int, error) {
	val := int(hex)
	if val >= 48 && val <= 57 {
		return val - 48, nil
	} else if val >= 97 && val <= 102 {
		return val - 87, nil
	} else if val >= 65 && val <= 70 {
		return val - 55, nil
	}

	return 0, errors.New(fmt.Sprintf("Invalid hex character received: %v", hex))
}

// parseByte is the opposite of parseHex, parseByte converts the given byte to its corresponding hex character.
// parseByte always return upper case characters.
func parseByte(b byte) (byte, error) {
	val := int(b)

	if val < 10 {
		return byte(val + 48), nil
	} else if val < 16 {
		return byte(val + 55), nil
	}
	return 0, errors.New(fmt.Sprintf("Byte cannot be parsed to hex: %v", b))
}

// HexToByte converts a given hex string to a byte list.
func HexToByte(hex string) ([]byte, error) {
	val := make([]byte, len(hex)/2)

	for i := 0; i < len(hex); i += 2 {
		a, err := parseHex(byte(hex[i]))
		if err != nil {
			return nil, err
		}
		if i + 1 >= len(hex) {
			return nil, errors.New("Hex array length is odd")
		}
		b, err := parseHex(byte(hex[i + 1]))
		if err != nil {
			return nil, err
		}
		val[i/2] = byte(a << 4 | b)
	}

	return val, nil
}

// ByteToHex converts a given byte list to the corresponding hex string.
func ByteToHex(val []byte) (string, error) {
	var hex string
	for _, h := range val {
		firstHex, err := parseByte(h >> 4)
		if err != nil {
			return "", err
		}
		secondHex, err := parseByte(h & 0x0f)
		if err != nil {
			return "", err
		}
		hex = hex + string(firstHex) + string(secondHex)
	}
	return hex, nil
}