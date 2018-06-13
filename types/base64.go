package types

import (
	"fmt"
	"strconv"
)

type Base64 struct {
	B []byte
}

var base64Map = map[int64]string{
	0:  "A", 16: "Q", 32: "g", 48: "w",
	1:  "B", 17: "R", 33: "h", 49: "x",
	2:  "C", 18: "S", 34: "i", 50: "y",
	3:  "D", 19: "T", 35: "j", 51: "z",
	4:  "E", 20: "U", 36: "k", 52: "0",
	5:  "F", 21: "V", 37: "l", 53: "1",
	6:  "G", 22: "W", 38: "m", 54: "2",
	7:  "H", 23: "X", 39: "n", 55: "3",
	8:  "I", 24: "Y", 40: "o", 56: "4",
	9:  "J", 25: "Z", 41: "p", 57: "5",
	10: "K", 26: "a", 42: "q", 58: "6",
	11: "L", 27: "b", 43: "r", 59: "7",
	12: "M", 28: "c", 44: "s", 60: "8",
	13: "N", 29: "d", 45: "t", 61: "9",
	14: "O", 30: "e", 46: "u", 62: "+",
	15: "P", 31: "f", 47: "v", 63: "/",
}

var base64RevMap = map[string]int64{
	"A": 0, "Q": 16, "g": 32, "w": 48,
	"B": 1, "R": 17, "h": 33, "x": 49,
	"C": 2, "S": 18, "i": 34, "y": 50,
	"D": 3, "T": 19, "j": 35, "z": 51,
	"E": 4, "U": 20, "k": 36, "0": 52,
	"F": 5, "V": 21, "l": 37, "1": 53,
	"G": 6, "W": 22, "m": 38, "2": 54,
	"H": 7, "X": 23, "n": 39, "3": 55,
	"I": 8, "Y": 24, "o": 40, "4": 56,
	"J": 9, "Z": 25, "p": 41, "5": 57,
	"K": 10, "a": 26, "q": 42, "6": 58,
	"L": 11, "b": 27, "r": 43, "7": 59,
	"M": 12, "c": 28, "s": 44, "8": 60,
	"N": 13, "d": 29, "t": 45, "9": 61,
	"O": 14, "e": 30, "u": 46, "+": 62,
	"P": 15, "f": 31, "v": 47, "/": 63,
}

func (b64 *Base64) Set(b []byte) Cipher {
	b64.B = b
	return b64
}

func (b64 *Base64) Get() []byte {
	return b64.B
}

func (b64 *Base64) Decode() ([]byte, error) {
	var binary string
	for _, c := range b64.B {
		if c == '=' {
			break
		}
		binary = binary + fmt.Sprintf("%.8b", base64RevMap[string(c)])[2:]
	}

	var b []byte
	for i := 0; i < len(binary); i = i + 8 {
		if len(binary) - i < 8 {
			break
		}
		val, _ := strconv.ParseInt(binary[i: i + 8], 2, 64)
		b = append(b, byte(val))
	}
	return b, nil
}

func (b64 *Base64) Encode(b []byte) (error) {
	var binary string
	for _, c := range b {
		binary = fmt.Sprintf("%s%.8b", binary, c)
	}

	for i := 0; i < len(binary); i = i + 6 {
		base6 := binary[i : i+6]
		mappedValue, _ := strconv.ParseInt(base6, 2, 64)
		for _, elem := range base64Map[mappedValue] {
			b64.B = append(b64.B, byte(elem))
		}

	}
	return nil
}