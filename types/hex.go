package types

import (
	"errors"
	"fmt"
)

type Hex []byte

// parseByte is the opposite of parseHex, parseByte converts the given byte to its corresponding hex character.
// parseByte always return upper case characters.
func parseByte(b byte) byte {
	val := int(b)

	if val < 10 {
		return byte(val + 48)
	}
	return byte(val + 87)
}

func EncodeHex(b []byte) Hex {
	var h Hex

	for _, c := range b {
		firstHex := parseByte(byte(c) >> 4)
		secondHex := parseByte(byte(c) & 0x0f)

		h = append(h, firstHex, secondHex)
	}
	return h
}

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

	return 0, errors.New(fmt.Sprintf("invalid hex character received: %v", hex))
}

func (h Hex) Decode() ([]byte, error) {
	var b []byte

	if len(h)%2 != 0 {
		return b, errors.New("hex array length is odd")
	}

	for i := 0; i < len(h); i += 2 {
		p, err := parseHex(h[i])
		if err != nil {
			return b, err
		}
		q, err := parseHex(h[i+1])
		if err != nil {
			return b, err
		}
		b = append(b, byte(p<<4|q))
	}

	return b, nil
}
