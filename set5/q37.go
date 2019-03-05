package set5

import (
	"crypto/rand"
	"crypto/sha256"
	"matasano/exchange"
	"math/big"
)

type SRPClientZero struct {
	N    big.Int
	g    big.Int
	k    big.Int
	a    *big.Int
	A    big.Int
	salt big.Int
	B    big.Int
	S    big.Int

	K        [32]byte
	email    string
	password string
}

func (c *SRPClientZero) ReceiveParams(N, g, k big.Int) {
	c.N, c.g, c.k = N, g, k
}

func (c *SRPClientZero) SendUser(s *exchange.SRP) {
	c.email = "foo@boo.com"
	c.password = "password1234"

	s.ReceiveUser(c.email, c.password)
}

func (c *SRPClientZero) SendIA(s *exchange.SRP) {
	c.a, _ = rand.Int(rand.Reader, &c.N)

	// Set A to 0 lol
	c.A.SetInt64(0)
	s.ReceiveIA(c.A, c.email)
}

func (c *SRPClientZero) ReceiveSaltB(salt, B big.Int) {
	c.salt = salt
	c.B = B
}

func (c *SRPClientZero) ComputeHSK() {
	uH := sha256.Sum256(append(c.A.Bytes(), c.B.Bytes()...))
	u := exchange.Sha256ToBigInt(uH[:])
	xH := sha256.Sum256(append(c.salt.Bytes(), c.password...))
	x := exchange.Sha256ToBigInt(xH[:])

	var tmp big.Int
	tmp.Exp(&c.g, &x, nil)
	tmp.Mul(&c.k, &tmp)
	tmp.Sub(&c.B, &tmp)

	var tmp2 big.Int
	tmp2.Mul(&u, &x)
	tmp2.Add(c.a, &tmp2)

	// Password set secret to 0
	c.S.SetInt64(0)

	c.K = sha256.Sum256(c.S.Bytes())
}

func (c *SRPClientZero) SendHMAC(s *exchange.SRP) bool {
	HMAC := sha256.Sum256(append(c.K[:], c.salt.Bytes()...))
	return s.CheckHMAC(HMAC[:])
}

func SolveQ37() bool {
	var s exchange.SRP
	var c SRPClientZero

	// Signup
	s.SendStartParams(&c)
	c.SendUser(&s)

	// Login starts here
	s.GenSalt()
	c.SendIA(&s)
	s.SendSaltB(&c)
	s.ComputeHSK()
	c.ComputeHSK()

	return c.SendHMAC(&s)
}
