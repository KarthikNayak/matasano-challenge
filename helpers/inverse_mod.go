package helpers

import (
	"math/big"
)

func InverseMod(x, y *big.Int) *big.Int {
	z := new(big.Int)

	val := new(big.Int).SetInt64(1)

	for z.Int64() == 0 {
		tmp := new(big.Int).Mul(y, val)
		tmp.Add(tmp, new(big.Int).SetInt64(1))
		mod := new(big.Int).Mod(tmp, x)

		if mod.Int64() == 0 {
			return z.Div(tmp, x)
		}

		val.Add(val, new(big.Int).SetInt64(1))
		if val.Int64() == x.Int64() {
			break
		}
	}
	return z
}
