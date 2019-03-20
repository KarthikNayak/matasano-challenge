package set5

import (
	"crypto/rand"
	"crypto/sha256"
	"matasano/exchange"
	"math/big"
)

type SRPClient struct {
	N    big.Int
	g    big.Int
	k    big.Int
	a    *big.Int
	A    big.Int
	salt []byte
	B    big.Int
	S    big.Int

	K        [32]byte
	email    string
	password string
}

func (c *SRPClient) ReceiveParams(N, g, k big.Int) {
	c.N, c.g, c.k = N, g, k
}

func (c *SRPClient) SendUser(s exchange.SRPServer) {
	c.email = "foo@boo.com"
	c.password = "password1234"

	s.ReceiveUser(c.email, c.password)
}

func (c *SRPClient) SendIA(s exchange.SRPServer) {
	c.a, _ = rand.Int(rand.Reader, &c.N)
	c.A.Exp(&c.g, c.a, &c.N)

	s.ReceiveIA(c.A, c.email)
}

func (c *SRPClient) ReceiveSaltB(salt []byte, B big.Int) {
	c.salt = salt
	c.B = B
}

func (c *SRPClient) ComputeHSK(simple bool) {
	uH := sha256.Sum256(append(c.A.Bytes(), c.B.Bytes()...))
	u := exchange.Sha256ToBigInt(uH[:])
	xH := sha256.Sum256(append(c.salt, c.password...))
	x := exchange.Sha256ToBigInt(xH[:])

	var tmp2 big.Int
	tmp2.Mul(u, x)
	tmp2.Add(c.a, &tmp2)
	tmp2.Mod(&tmp2, &c.N)

	if !simple {
		var tmp big.Int
		tmp.Exp(&c.g, x, &c.N)
		tmp.Mul(&c.k, &tmp)
		tmp.Sub(&c.B, &tmp)

		c.S.Exp(&tmp, &tmp2, &c.N)
		c.K = sha256.Sum256(c.S.Bytes())
	} else {
		c.S.Exp(&c.B, &tmp2, &c.N)
		c.K = sha256.Sum256(c.S.Bytes())
	}
}

func (c *SRPClient) SendHMAC(s exchange.SRPServer) bool {
	HMAC := sha256.Sum256(append(c.K[:], c.salt...))
	return s.CheckHMAC(HMAC[:])
}

func SolveQ36() bool {
	var s exchange.SRP
	var c SRPClient

	s.SendStartParams(&c)
	c.SendUser(&s)

	s.GenSalt()
	c.SendIA(&s)
	s.SendSaltB(&c)

	s.ComputeHSK()
	c.ComputeHSK(false)

	return c.SendHMAC(&s)
}
