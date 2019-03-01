package set4

import (
	"bytes"
	"encoding/binary"
	"errors"
	"matasano/hash"
	"matasano/types"
)

var (
	q29Msg    = []byte("comment1=cooking%20MCs;userdata=foo;comment2=%20like%20a%20pound%20of%20bacon")
	q29NewMsg = []byte(";admin=true")
)

func gluePadding(ml uint64) []byte {
	ml = uint64(ml) * 8
	// append the bit '1' to the message e.g. by adding 0x80 if message length is a multiple of 8 bits.
	data := make([]byte, ml/8)
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

	return data[ml/8:]
}

func byteToUint64(b []byte) uint64 {
	return uint64(b[0])<<24 | uint64(b[1])<<16 | uint64(b[2])<<8 | uint64(b[3])
}

func splitHex(h []byte) (a, b, c, d, e uint64) {
	a = byteToUint64(h[0:4])
	b = byteToUint64(h[4:8])
	c = byteToUint64(h[8:12])
	d = byteToUint64(h[12:16])
	e = byteToUint64(h[16:20])
	return
}

func writeCorrectLength(b []byte, l int) []byte {
	l = l * 8
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint64(l))
	return append(b[:len(b)-8], buf.Bytes()...)
}

func SolveQ29() error {
	oracle, hasher := hash.Sha1MACOracle()
	msgSha1 := hasher(q29Msg)

	for i := 1; i < 30; i++ {
		glue := gluePadding(uint64(i + len(q29Msg)))
		forge := append(append(q29Msg, glue...), q29NewMsg...)

		h := types.Hex{B: []byte(msgSha1)}
		hex, _ := h.Decode()

		h0, h1, h2, h3, h4 := splitHex(hex)

		pData := hash.PreProcess(q29NewMsg)
		pData = writeCorrectLength(pData, len(forge)+i)
		sha1 := hash.Sha1Base(pData, uint32(h0), uint32(h1), uint32(h2), uint32(h3), uint32(h4))

		if oracle(forge, sha1) {
			return nil
		}
	}
	return errors.New("couldn't find the key")
}
