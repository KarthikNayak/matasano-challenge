package set6

import (
	"errors"
	"matasano/keys"
	"math/big"
)

func DSAGeneratorGZero() (*big.Int, *big.Int, *big.Int) {
	P, _ := new(big.Int).SetString(keys.DefaultP, 16)
	Q, _ := new(big.Int).SetString(keys.DefaultQ, 16)
	G := new(big.Int).SetInt64(0)

	return P, Q, G
}

func DSAGeneratorGPPlusOne() (*big.Int, *big.Int, *big.Int) {
	n1 := new(big.Int).SetInt64(1)
	P, _ := new(big.Int).SetString(keys.DefaultP, 16)
	Q, _ := new(big.Int).SetString(keys.DefaultQ, 16)
	G := new(big.Int).Add(P, n1)

	return P, Q, G
}

func SolveQ45Part1() error {
	var d keys.DSA
	d.GenerateKeys(DSAGeneratorGZero)

	m := "Hey, Sup!"
	// r -> 0 as G -> 0 :(
	r, s, err := d.Sign(m)
	if err != nil {
		return err
	}

	if !d.Verify(r, s, m) {
		return errors.New("expected true got false")
	}

	// Notice how it verifies anything?
	if !d.Verify(r, s, "wtf") {
		return errors.New("expected true got false")
	}

	return nil
}

func SolveQ45Part2() error {
	var d keys.DSA
	d.GenerateKeys(DSAGeneratorGPPlusOne)

	z := new(big.Int).SetInt64(1)

	r, s := new(big.Int), new(big.Int)

	r.Exp(d.Y, z, d.P).Mod(z, d.Q)
	s.ModInverse(z, d.Q).Mul(s, r).Mod(s, d.Q)

	if !d.Verify(r, s, "wtf") {
		return errors.New("expected true got false")
	}

	return nil
}
