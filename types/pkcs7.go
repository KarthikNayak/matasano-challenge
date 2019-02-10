package types

import (
	"github.com/pkg/errors"
)

type PKCS7 struct {
	B         []byte
	blockSize int
}

func (p *PKCS7) Set(b []byte) Type {
	p.B = b
	return p
}

func (p *PKCS7) Get() []byte {
	return p.B
}

func (p *PKCS7) SetBlockSize(blockSize int) {
	p.blockSize = blockSize
}

func (p *PKCS7) Decode() ([]byte, error) {
	if len(p.B) == 0 {
		return []byte{}, nil
	}

	length := len(p.B)
	padding := p.B[length-1]
	if padding == 0 {
		return []byte{}, errors.New("improper padding for pkcs7")
	}
	for i := 1; i < int(padding); i++ {
		if p.B[length-1-i] != padding {
			return []byte{}, errors.New("improper padding for pkcs7")
		}
	}
	return p.B[:length-int(padding)], nil
}

func (p *PKCS7) Encode(b []byte) error {
	if p.blockSize == 0 {
		return errors.New("Block Size not set, use SetBlockSize() to set the block size")
	}

	srcLen := len(b)
	padding := 0

	if srcLen > p.blockSize {
		extra := srcLen % p.blockSize
		if extra > 0 {
			padding = p.blockSize - extra
		}
	} else {
		padding = p.blockSize - srcLen
	}

	output := make([]byte, srcLen+padding)
	copy(output, b)
	for i := 0; i < padding; i++ {
		output[srcLen+i] = byte(padding)
	}
	p.B = output
	return nil
}
