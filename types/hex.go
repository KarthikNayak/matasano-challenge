package types

import (
	"errors"
	"fmt"
)

type Hex struct {
	B []byte
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

// parseByte is the opposite of parseHex, parseByte converts the given byte to its corresponding hex character.
// parseByte always return upper case characters.
func parseByte(b byte) (byte, error) {
	val := int(b)

	if val < 10 {
		return byte(val + 48), nil
	} else if val < 16 {
		return byte(val + 55), nil
	}
	return 0, errors.New(fmt.Sprintf("nyte cannot be parsed to hex: %v", b))
}

func (h *Hex) Set(b []byte) Type {
	h.B = b
	return h
}

func (h *Hex) Get() []byte {
	return h.B
}

func (h *Hex) Decode() ([]byte, error) {
	var val []byte

	for i := 0; i < len(h.B); i += 2 {
		a, err := parseHex(h.B[i])
		if err != nil {
			return val, err
		}
		if i + 1 >= len(h.B) {
			return val, errors.New("hex array length is odd")
		}
		b, err := parseHex(h.B[i + 1])
		if err != nil {
			return val, err
		}
		val = append(val, byte(a << 4 | b))
	}

	return val, nil
}

func (h *Hex) Encode(b []byte) (error) {
	for _, c := range b {
		firstHex, err := parseByte(byte(c) >> 4)
		if err != nil {
			return err
		}

		secondHex, err := parseByte(byte(c) & 0x0f)
		if err != nil {
			return err
		}
		h.B = append(h.B, firstHex, secondHex)
	}
	return nil
}

