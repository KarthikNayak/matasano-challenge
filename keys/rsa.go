package keys

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

const (
	RSABitSize = 100
)

type RSA struct {
	// Public Key: [e, n]
	// Private Key: [d, n]
	E *big.Int
	d *big.Int
	N *big.Int
}

func (r *RSA) GenerateKeys() error {
	for r.d == nil {
		p, err := rand.Prime(rand.Reader, RSABitSize)
		if err != nil {
			return err
		}

		q, err := rand.Prime(rand.Reader, RSABitSize)
		if err != nil {
			return err
		}

		r.N = new(big.Int).Mul(p, q)

		one := new(big.Int).SetInt64(1)
		p1 := new(big.Int).Sub(p, one)
		q1 := new(big.Int).Sub(q, one)

		r.E = new(big.Int).SetInt64(3)
		et := new(big.Int).Mul(p1, q1)

		d := new(big.Int).ModInverse(r.E, et)
		if d != nil {
			r.d = d
		}
	}
	return nil
}

func (r *RSA) EncryptBigInt(m *big.Int) *big.Int {
	return new(big.Int).Exp(m, r.E, r.N)
}

func (r *RSA) DecryptBigInt(c *big.Int) *big.Int {
	return new(big.Int).Exp(c, r.d, r.N)
}

func (r *RSA) EncryptString(m string) *big.Int {
	// Encode to Hex
	h := make([]byte, hex.EncodedLen(len(m)))
	hex.Encode(h, []byte(m))

	// Hex to Big Int
	v, _ := new(big.Int).SetString(string(h), 16)
	return new(big.Int).Exp(v, r.E, r.N)
}

func (r *RSA) DecryptString(c *big.Int) string {
	// Decode it first
	d := new(big.Int).Exp(c, r.d, r.N)
	h := d.Text(16)

	// Convert back to string
	m := make([]byte, len(h)/2)
	hex.Decode(m, []byte(h))
	return string(m)
}
