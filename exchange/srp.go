package exchange

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const (
	N = "ffffffffffffffffc90fdaa22168c234c4c6628b80dc1cd129024e088a67cc74020bbea63b139b22514a08798e3404ddef9519b3cd3a431b302b0a6df25f14374fe1356d6d51c245e485b576625e7ec6f44c42e9a637ed6b0bff5cb6f406b7edee386bfb5a899fa5ae9f24117c4b1fe649286651ece45b3dc2007cb8a163bf0598da48361c55d39a69163fa8fd24cf5f83655d23dca3ad961c62f356208552bb9ed529077096966d670c354e4abc9804f1746c08ca237327ffffffffffffffff"
	// g = "2" declared in deffie-hellman
	k = "3"
)

func Sha256ToBigInt(sha []byte) *big.Int {
	x := new(big.Int)
	x.SetString(fmt.Sprintf("%x", sha), 16)
	return x
}

type SRPClient interface {
	ReceiveParams(N, g, k big.Int)
	SendUser(s SRPServer)
	SendIA(s SRPServer)
	ReceiveSaltB(salt []byte, B big.Int)
	ComputeHSK(simple bool)
	SendHMAC(s SRPServer) bool
}

type SRPServer interface {
	SendStartParams(c SRPClient)
	ReceiveUser(email, password string)
	GenSalt()
	ReceiveIA(A big.Int, I string)
	SendSaltB(c SRPClient)
	ComputeHSK()
	CheckHMAC(HMAC []byte) bool
}

type SRP struct {
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

func (s *SRP) SendStartParams(c SRPClient) {
	s.N.SetString(N, 16)
	s.g.SetString(g, 16)
	s.k.SetString(k, 16)

	c.ReceiveParams(s.N, s.g, s.k)
}

func (s *SRP) ReceiveUser(email, password string) {
	s.I = email
	s.P = password
}

func (s *SRP) GenSalt() {
	rand.Read(s.salt)
	xH := sha256.Sum256(append(s.salt, []byte(s.P)...))
	x := Sha256ToBigInt(xH[:])

	s.v.Exp(&s.g, x, &s.N)
}

func (s *SRP) ReceiveIA(A big.Int, I string) {
	s.A = A
}

func (s *SRP) SendSaltB(c SRPClient) {
	var tmp big.Int
	s.B.Set(tmp.Mul(&s.k, &s.v))
	s.b, _ = rand.Int(rand.Reader, &s.N)
	tmp.Exp(&s.g, s.b, &s.N)
	s.B.Add(&s.B, &tmp)

	c.ReceiveSaltB(s.salt, s.B)
}

func (s *SRP) ComputeHSK() {
	uH := sha256.Sum256(append(s.A.Bytes(), s.B.Bytes()...))
	u := Sha256ToBigInt(uH[:])

	var tmp big.Int
	tmp.Exp(&s.v, u, &s.N)
	tmp.Mul(&s.A, &tmp)

	s.S.Exp(&tmp, s.b, &s.N)
	s.K = sha256.Sum256(s.S.Bytes())
}

func (s *SRP) CheckHMAC(HMAC []byte) bool {
	HMAC2 := sha256.Sum256(append(s.K[:], s.salt...))
	return bytes.Compare(HMAC, HMAC2[:]) == 0
}
