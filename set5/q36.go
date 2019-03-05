package set5

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"matasano/exchange"
	"math/big"
)

type SRPClient struct {
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

func (c *SRPClient) ReceiveParams(N, g, k big.Int) {
	c.N, c.g, c.k = N, g, k
}

func (c *SRPClient) SendUser(s *exchange.SRP) {
	c.email = "foo@boo.com"
	c.password = "password1234"

	s.ReceiveUser(c.email, c.password)
}

func (c *SRPClient) SendIA(s *exchange.SRP) {
	c.a, _ = rand.Int(rand.Reader, &c.N)
	c.A.Exp(&c.g, c.a, &c.N)

	s.ReceiveIA(c.A, c.email)
}

func (c *SRPClient) ReceiveSaltB(salt, B big.Int) {
	c.salt = salt
	c.B = B
}

func (c *SRPClient) ComputeHSK() {
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

	c.S.Exp(&tmp, &tmp2, &c.N)
	c.K = sha256.Sum256(c.S.Bytes())
}

func (c *SRPClient) SendHMAC(s *exchange.SRP) bool {
	HMAC := sha256.Sum256(append(c.K[:], c.salt.Bytes()...))
	return s.CheckHMAC(HMAC[:])
}

func SolveQ36() {
	var s exchange.SRP
	var c SRPClient

	s.SendStartParams(&c)
	c.SendUser(&s)

	s.GenSalt()

	c.SendIA(&s)

	s.SendSaltB(&c)

	s.ComputeHSK()
	c.ComputeHSK()

	fmt.Println(c.SendHMAC(&s))
}
