package set5

import (
	"matasano/keys"
	"math/big"
)

func SolveQ40() bool {
	var r1, r2, r3 keys.RSA

	r1.GenerateKeys()
	r2.GenerateKeys()
	r3.GenerateKeys()

	secretMsg := new(big.Int).SetInt64(1223342)

	c1 := r1.EncryptBigInt(secretMsg)
	c2 := r2.EncryptBigInt(secretMsg)
	c3 := r3.EncryptBigInt(secretMsg)

	n1, n2, n3 := r1.N, r2.N, r3.N
	n := new(big.Int).Mul(n1, n2)
	n.Mul(n, n3)

	ms1 := new(big.Int).Mul(n2, n3)
	ms2 := new(big.Int).Mul(n1, n3)
	ms3 := new(big.Int).Mul(n1, n2)

	t1 := new(big.Int).ModInverse(ms1, n1)
	t1.Mul(t1, ms1)
	t1.Mul(t1, c1)

	t2 := new(big.Int).ModInverse(ms2, n2)
	t2.Mul(t2, ms2)
	t2.Mul(t2, c2)

	t3 := new(big.Int).ModInverse(ms3, n3)
	t3.Mul(t3, ms3)
	t3.Mul(t3, c3)

	t := new(big.Int).Add(t1, t2)
	t.Add(t, t3)
	t.Mod(t, n)

	return t.Cmp(secretMsg.Exp(secretMsg, new(big.Int).SetInt64(3), nil)) == 0
}
