package keys

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

const (
	DefaultP = "800000000000000089e1855218a0e7dac38136ffafa72eda7859f2171e25e65eac698c1702578b07dc2a1076da241c76c62d374d8389ea5aeffd3226a0530cc565f3bf6b50929139ebeac04f48c3c84afb796d61e5a4f9a8fda812ab59494232c7d2b4deb50aa18ee9e132bfa85ac4374d7f9091abc3d015efc871a584471bb1"
	DefaultQ = "f4f47f05794b256174bba6e9b396a7707e563c5b"
	DefaultG = "5958c9d3898b224b12672c0b98e06c60df923cb8bc999d119458fef538b8fa4046c8db53039db620c094c9fa077ef389b5322a559946a71903f990f1f7e0e025e2d7f7cf494aff1a0470f5b64c36b625a097f1651fe775323556fe00b3608c887892878480e99041be601a62166ca6894bdd41a7054ec89f756ba9fc95302291"
)

func DSAGeneratePQG() (*big.Int, *big.Int, *big.Int) {
	P, _ := new(big.Int).SetString(DefaultP, 16)
	Q, _ := new(big.Int).SetString(DefaultQ, 16)
	G, _ := new(big.Int).SetString(DefaultG, 16)

	return P, Q, G
}

type CustomDSAGenerator func() (*big.Int, *big.Int, *big.Int)

type DSA struct {
	P, Q, G *big.Int
	// private key
	x *big.Int
	// public key
	Y *big.Int
}

func (d *DSA) GenerateKeys(f CustomDSAGenerator) error {
	var err error
	if f == nil {
		d.P, d.Q, d.G = DSAGeneratePQG()
	} else {
		d.P, d.Q, d.G = f()
	}

	d.x, err = rand.Int(rand.Reader, d.Q)
	if err != nil {
		return err
	}
	d.Y = new(big.Int).Exp(d.G, d.x, d.P)

	return nil
}

func DSAPrivateKeyFromK(Q, H, k, s, r *big.Int) (x *big.Int) {
	x = new(big.Int)

	x.Mul(s, k)
	x.Sub(x, H)
	x.Mul(x, new(big.Int).ModInverse(r, Q)).Mod(x, Q)

	return
}

func (d *DSA) Sign(m string) (*big.Int, *big.Int, error) {
	// calcK:
	k, err := rand.Int(rand.Reader, d.Q)
	if err != nil {
		return nil, nil, err
	}

	r := new(big.Int).Exp(d.G, k, d.P)
	r.Mod(r, d.Q)

	// if r.Int64() == 0 {
	//	goto calcK
	// }

	h := sha256.Sum256([]byte(m))

	kInv := new(big.Int).ModInverse(k, d.Q)
	s := new(big.Int).Mul(d.x, r)
	s.Add(s, new(big.Int).SetBytes(h[:])).Mul(s, kInv).Mod(s, d.Q)

	// if s.Int64() == 0 {
	//	goto calcK
	// }

	return r, s, nil
}

func (d *DSA) Verify(r, s *big.Int, m string) bool {
	// n0 := new(big.Int).SetInt64(0)

	// if r.Cmp(n0) <= 0 || r.Cmp(d.Q) >= 0 {
	//	return false
	// }
	// if s.Cmp(n0) <= 0 || s.Cmp(d.Q) >= 0 {
	//	return false
	// }

	w := new(big.Int).ModInverse(s, d.Q)
	h := sha256.Sum256([]byte(m))

	u1 := new(big.Int).SetBytes(h[:])
	u1.Mul(u1, w).Mod(u1, d.Q)

	u2 := new(big.Int).Mul(r, w)
	u2.Mod(u2, d.Q)

	v := new(big.Int).Exp(d.G, u1, d.P)
	tmp := new(big.Int).Exp(d.Y, u2, d.P)
	v.Mul(v, tmp).Mod(v, d.P).Mod(v, d.Q)

	if v.Cmp(r) == 0 {
		return true
	}

	return false
}
