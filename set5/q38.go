package set5

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"matasano/exchange"
	"math/big"
)

const (
	N = "ffffffffffffffffc90fdaa22168c234c4c6628b80dc1cd129024e088a67cc74020bbea63b139b22514a08798e3404ddef9519b3cd3a431b302b0a6df25f14374fe1356d6d51c245e485b576625e7ec6f44c42e9a637ed6b0bff5cb6f406b7edee386bfb5a899fa5ae9f24117c4b1fe649286651ece45b3dc2007cb8a163bf0598da48361c55d39a69163fa8fd24cf5f83655d23dca3ad961c62f356208552bb9ed529077096966d670c354e4abc9804f1746c08ca237327ffffffffffffffff"
	g = "2"
	k = "3"
)

type SRPServerFake struct {
	N    big.Int
	g    big.Int
	k    big.Int
	salt []byte
	v    big.Int
	b    *big.Int
	B    big.Int
	A    big.Int
	S    big.Int

	K    [32]byte
	I, P string
}

func (s *SRPServerFake) SendStartParams(c exchange.SRPClient) {
	s.N.SetString(N, 16)
	s.g.SetString(g, 16)
	s.k.SetString(k, 16)

	c.ReceiveParams(s.N, s.g, s.k)
}

func (s *SRPServerFake) ReceiveUser(email, password string) {
	s.I = email
	s.P = password
}

func (s *SRPServerFake) GenSalt() {
	rand.Read(s.salt)
	xH := sha256.Sum256(append(s.salt, []byte(s.P)...))
	x := exchange.Sha256ToBigInt(xH[:])

	s.v.Exp(&s.g, x, &s.N)
}

func (s *SRPServerFake) ReceiveIA(A big.Int, I string) {
	s.A = A
}

func (s *SRPServerFake) SendSaltB(c exchange.SRPClient) {
	// var tmp big.Int
	// s.B.Set(tmp.Mul(&s.k, &s.v))
	// s.b, _ = rand.Int(rand.Reader, &s.N)
	// tmp.Exp(&s.g, s.b, &s.N)
	// s.B.Add(&s.B, &tmp)
	var tmp big.Int
	s.b = &tmp

	c.ReceiveSaltB([]byte{}, tmp)
}

func (s *SRPServerFake) ComputeHSK() {
	uH := sha256.Sum256(append(s.A.Bytes(), s.B.Bytes()...))
	u := exchange.Sha256ToBigInt(uH[:])

	var tmp big.Int
	tmp.Exp(&s.v, u, &s.N)
	tmp.Mul(&s.A, &tmp)

	s.S.Exp(&tmp, s.b, &s.N)
	s.K = sha256.Sum256(s.S.Bytes())
}

func (s *SRPServerFake) CheckHMAC(HMAC []byte) bool {

	uH := sha256.Sum256(append(s.A.Bytes(), s.B.Bytes()...))
	u := exchange.Sha256ToBigInt(uH[:])
	pass := []byte("password1234")
	xH := sha256.Sum256(append(s.salt, pass...))
	x := exchange.Sha256ToBigInt(xH[:])

	var tmp big.Int
	tmp.Mul(u, x)
	tmp.Add(&s.A, &tmp)
	tmp.Mod(&tmp, &s.N)

	s.S.Exp(&s.B, &tmp, &s.N)
	s.K = sha256.Sum256(s.S.Bytes())

	HMAC2 := sha256.Sum256(append(s.K[:], s.salt...))

	fmt.Println(HMAC)
	fmt.Println(HMAC2)

	fmt.Println("password:", string(pass))

	return bytes.Compare(HMAC, HMAC2[:]) == 0
}

func SolveQ38() bool {
	var s SRPServerFake
	var c SRPClient

	// Signup
	s.SendStartParams(&c)
	c.SendUser(&s)

	s.GenSalt()
	c.SendIA(&s)
	s.SendSaltB(&c)
	s.ComputeHSK()
	c.ComputeHSK(true)

	return c.SendHMAC(&s)
}
