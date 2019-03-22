package set6

import (
	"crypto/sha256"
	"matasano/keys"
	"math/big"
)

func rsaOracleDecrypt() (*keys.RSA, func(c *big.Int) *big.Int) {
	r := new(keys.RSA)
	r.GenerateKeys()

	m := make(map[[32]byte]bool)
	return r, func(c *big.Int) *big.Int {
		s := sha256.Sum256(c.Bytes())
		if _, ok := m[s]; ok {
			return nil
		}

		m[s] = true
		return r.DecryptBigInt(c)
	}
}

func SolveQ41(val int64) int64 {
	r, f := rsaOracleDecrypt()
	c := r.EncryptBigInt(new(big.Int).SetInt64(val))

	s := new(big.Int).SetInt64(1234)

	c1 := new(big.Int).Exp(s, r.E, r.N)
	c1.Mul(c1, c)
	c1.Mod(c1, r.N)

	p1 := f(c1)

	s1 := new(big.Int).ModInverse(s, r.N)
	s1.Mul(s1, p1)
	s1.Mod(s1, r.N)

	return s1.Int64()
}
