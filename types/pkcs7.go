package types

import (
	"github.com/pkg/errors"
)

type PKCS7 struct {
	B []byte
	size int
}

func (p *PKCS7) Set(b []byte) Cipher {
	p.B = b
	return p
}

func (p *PKCS7) Get() []byte {
	return p.B
}

func (p *PKCS7) SetBlockSize(size int) {
	p.size = size
}

func (p *PKCS7) Decode() ([]byte, error) {
	if len(p.B) == 0 {
		return []byte{}, nil
	}

	length := len(p.B)
	padding := p.B[length - 1]
	for i := 1; i < int(padding); i++ {
		if p.B[length - 1 - i] != padding {
			return p.B, nil
		}
	}
	return p.B[:length - int(padding)], nil
}

func (p *PKCS7) Encode(b []byte) (error) {
	if p.size == 0 {
		return errors.New("Block Size not set, use SetBlockSize() to set the block size")
	}

	srcLen := len(b)
	deltaDiff := p.size - srcLen
	if deltaDiff < 0 {
		return errors.New("Size is smaller than the given byte array length")
	}

	output := make([]byte, p.size)
	copy(output, b)
	for i := 0; i < deltaDiff; i++ {
		output[srcLen + i] = byte(deltaDiff)
	}
	p.B = output
	return nil
}

