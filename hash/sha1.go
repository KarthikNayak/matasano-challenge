package hash

import (
	"bytes"
	"encoding/binary"
	"matasano/types"
	"strings"
)

const ()

func preProcess(data []byte) []byte {
	ml := uint64(len(data)) * 8

	// append the bit '1' to the message e.g. by adding 0x80 if message length is a multiple of 8 bits.
	data = append(data, 0x80)

	// 64 byte padding
	padding := ((len(data)/64)+1)*64 - len(data)
	for i := 0; i < padding; i++ {
		data = append(data, byte(0))
	}

	// append ml, the original message length, as a 64-bit big-endian integer.
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, ml)
	data = append(data[:len(data)-8], buf.Bytes()...)

	return data
}

func LeftRotate(x uint32, n uint32) uint32 {
	return (x << n) | (x >> (32 - n))
}

func Sha1(data []byte) string {
	pData := preProcess(data)

	var h0 uint32 = 0x67452301
	var h1 uint32 = 0xEFCDAB89
	var h2 uint32 = 0x98BADCFE
	var h3 uint32 = 0x10325476
	var h4 uint32 = 0xC3D2E1F0

	// Process the message in successive 512-bit chunks
	for i := 0; i < len(pData); i += 64 {
		// break chunk into sixteen 32-bit big-endian words
		w := make([]uint32, 80)
		for j := 0; j < 16; j++ {
			w[j] = binary.BigEndian.Uint32(pData[i+j*4 : i+j*4+4])
		}

		//Extend the sixteen 32-bit words into eighty 32-bit words
		for j := 16; j < 80; j++ {
			w[j] = LeftRotate(w[j-3]^w[j-8]^w[j-14]^w[j-16], 1)
		}

		// Initialize hash value for this chunk
		a := h0
		b := h1
		c := h2
		d := h3
		e := h4

		//Main loop
		for j := 0; j < 80; j++ {
			var f, k uint32
			if j <= 19 {
				f = (b & c) | ((^b) & d)
				k = 0x5A827999
			} else if j <= 39 {
				f = b ^ c ^ d
				k = 0x6ED9EBA1
			} else if j <= 59 {
				f = (b & c) | (b & d) | (c & d)
				k = 0x8F1BBCDC
			} else {
				f = b ^ c ^ d
				k = 0xCA62C1D6
			}
			var temp = uint32(LeftRotate(a, 5) + f + e + k + w[j])
			e = d
			d = c
			c = LeftRotate(b, 30)
			b = a
			a = temp
		}
		h0 = h0 + a
		h1 = h1 + b
		h2 = h2 + c
		h3 = h3 + d
		h4 = h4 + e
	}

	hash := make([]byte, 20)

	hash[0], hash[1], hash[2], hash[3] = byte(h0>>24), byte(h0>>16), byte(h0>>8), byte(h0)
	hash[4], hash[5], hash[6], hash[7] = byte(h1>>24), byte(h1>>16), byte(h1>>8), byte(h1)
	hash[8], hash[9], hash[10], hash[11] = byte(h2>>24), byte(h2>>16), byte(h2>>8), byte(h2)
	hash[12], hash[13], hash[14], hash[15] = byte(h3>>24), byte(h3>>16), byte(h3>>8), byte(h3)
	hash[16], hash[17], hash[18], hash[19] = byte(h4>>24), byte(h4>>16), byte(h4>>8), byte(h4)

	hex := types.Hex{}
	hex.Encode(hash)
	return strings.ToLower(string(hex.Get()))
}
