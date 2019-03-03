package exchange

import (
	"crypto/rand"
	"crypto/sha1"
	"math/big"
)

type DF struct {
	a    *big.Int
	B    big.Int
	P    big.Int
	G    big.Int
	Key  big.Int
	Hash [20]byte
}

type DFClient interface {
	SetBPG(B, P, G big.Int)
	SetB(B big.Int)

	SendPGA(d2 DFClient) error
	SendB(d2 DFClient) error

	GenHash()
	GetHash() []byte
}

const (
	p = "ffffffffffffffffc90fdaa22168c234c4c6628b80dc1cd129024e088a67cc74020bbea63b139b22514a08798e3404ddef9519b3cd3a431b302b0a6df25f14374fe1356d6d51c245e485b576625e7ec6f44c42e9a637ed6b0bff5cb6f406b7edee386bfb5a899fa5ae9f24117c4b1fe649286651ece45b3dc2007cb8a163bf0598da48361c55d39a69163fa8fd24cf5f83655d23dca3ad961c62f356208552bb9ed529077096966d670c354e4abc9804f1746c08ca237327ffffffffffffffff"
	g = "2"
)

func (d *DF) SendPGA(d2 DFClient) error {
	d.P.SetString(p, 16)
	d.G.SetString(g, 16)

	var err error
	d.a, err = rand.Int(rand.Reader, &d.P)
	if err != nil {
		return err
	}

	var A big.Int
	A.Exp(&d.G, d.a, &d.P)
	d2.SetBPG(A, d.P, d.G)

	return nil
}

func (d *DF) SendB(d2 DFClient) error {
	var err error
	d.a, err = rand.Int(rand.Reader, &d.P)
	if err != nil {
		return err
	}

	var A big.Int
	A.Exp(&d.G, d.a, &d.P)
	d2.SetB(A)

	return nil
}

func (d *DF) SetB(B big.Int) {
	d.B = B
}

func (d *DF) SetBPG(B, P, G big.Int) {
	d.B = B
	d.P = P
	d.G = G
}

func (d *DF) GenHash() {
	d.Key.Exp(&d.B, d.a, &d.P)
	d.Hash = sha1.Sum(d.Key.Bytes())
}

func (d *DF) GetHash() []byte {
	return d.Hash[:]
}
