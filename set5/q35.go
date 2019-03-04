package set5

import (
	"crypto/sha1"
	"fmt"
	"matasano/cipher"
	"matasano/exchange"
	"matasano/types"
	"math/big"
)

type SetBPGModifier func(A, P, G big.Int) (big.Int, big.Int, big.Int)
type SetBModifier func(B big.Int) big.Int
type FigureOutHash func(d *DFMITM2)

type DFMITM2 struct {
	G, P  big.Int
	A, B  big.Int
	Key   big.Int
	Hash1 [20]byte
	Hash2 [20]byte

	m1 SetBPGModifier
	m2 SetBModifier
	f  FigureOutHash

	D1 exchange.DFClient
	D2 exchange.DFClient
}

func (d *DFMITM2) SetBPG(B, P, G big.Int) {
	d.A = B
	d.P = P
	d.G = G

	d.D2.SetBPG(d.m1(B, P, G))
}

func (d *DFMITM2) SetB(B big.Int) {
	d.B = B
	d.D1.SetB(d.m2(B))

	// We have A, B, p, G lets generate the hash
	d.f(d)
}

func (d *DFMITM2) SendPGA(d2 exchange.DFClient) error {
	return nil
}

func (d *DFMITM2) SendB(d2 exchange.DFClient) error {
	return nil
}

func (d *DFMITM2) GenHash() {

}

func (d *DFMITM2) GetHash() []byte {
	return []byte{}
}

func (d *DFMITM2) SendAESMsg(d2 exchange.DFClient) error {
	return nil
}

func (d *DFMITM2) EchoMsg(msg []byte, d2 exchange.DFClient) error {
	origIV := make([]byte, 16)
	origMsgEncoded := make([]byte, len(msg)-16)
	copy(origIV, msg[len(msg)-16:])
	copy(origMsgEncoded, msg[:len(msg)-16])

	var c cipher.CBC
	err := c.Init(d.Hash1[:16], 16*8, origIV)
	if err != nil {
		return err
	}

	decoded, err := c.Decode(&types.Text{T: origMsgEncoded})
	if err != nil {
		return err
	}
	fmt.Println("Intercepted message", string(decoded))

	if d.Hash2 != [20]byte{} {
		origIV := make([]byte, 16)
		origMsgEncoded := make([]byte, len(msg)-16)
		copy(origIV, msg[len(msg)-16:])
		copy(origMsgEncoded, msg[:len(msg)-16])

		var c2 cipher.CBC
		err := c2.Init(d.Hash2[:16], 16*8, origIV)
		fmt.Println()
		if err != nil {
			return err
		}

		decoded2, err := c2.Decode(&types.Text{T: origMsgEncoded})
		if err != nil {
			return err
		}

		fmt.Println("Intercepted message (hash 2)", string(decoded2))
	}

	d.D2.EchoMsg(msg, d)

	return nil
}

func (d *DFMITM2) ReceiveMsg(msg []byte) error {
	origIV := make([]byte, 16)
	origMsgEncoded := make([]byte, len(msg)-16)
	copy(origIV, msg[len(msg)-16:])
	copy(origMsgEncoded, msg[:len(msg)-16])

	var c cipher.CBC
	err := c.Init(d.Hash1[:16], 16*8, origIV)
	if err != nil {
		return err
	}

	decoded, err := c.Decode(&types.Text{T: origMsgEncoded})
	if err != nil {
		return err
	}
	fmt.Println("Intercepted echo", string(decoded))

	if d.Hash2 != [20]byte{} {
		origIV := make([]byte, 16)
		origMsgEncoded := make([]byte, len(msg)-16)
		copy(origIV, msg[len(msg)-16:])
		copy(origMsgEncoded, msg[:len(msg)-16])

		var c2 cipher.CBC
		err := c2.Init(d.Hash2[:16], 16*8, origIV)
		if err != nil {
			return err
		}

		decoded2, err := c2.Decode(&types.Text{T: origMsgEncoded})
		if err != nil {
			return err
		}

		fmt.Println("Intercepted echo(hash2)", string(decoded2))
	}

	d.D1.ReceiveMsg(msg)

	return nil
}

func SolveQ35Prob1() error {
	// g = 1

	var Alice, Bob, MrX exchange.DFClient

	Alice = new(exchange.DF)
	Bob = new(exchange.DF)
	MrX = new(DFMITM2)

	MrX.(*DFMITM2).D1 = Alice
	MrX.(*DFMITM2).D2 = Bob
	MrX.(*DFMITM2).m1 = func(A, P, G big.Int) (big.Int, big.Int, big.Int) {
		var One big.Int
		One.SetInt64(1)
		return A, P, One
	}
	MrX.(*DFMITM2).m2 = func(B big.Int) big.Int {
		return B
	}
	MrX.(*DFMITM2).f = func(d *DFMITM2) {
		var x big.Int
		x.SetInt64(1)

		d.Hash1 = sha1.Sum(x.Bytes())
	}

	err := Alice.SendPGA(MrX)
	if err != nil {
		return err
	}

	err = Bob.SendB(MrX)
	if err != nil {
		return err
	}

	Alice.GenHash()
	Bob.GenHash()

	Alice.SendAESMsg(MrX)

	return nil
}

func SolveQ35Prob2() error {
	// g = p

	var Alice, Bob, MrX exchange.DFClient

	Alice = new(exchange.DF)
	Bob = new(exchange.DF)
	MrX = new(DFMITM2)

	MrX.(*DFMITM2).D1 = Alice
	MrX.(*DFMITM2).D2 = Bob
	MrX.(*DFMITM2).m1 = func(A, P, G big.Int) (big.Int, big.Int, big.Int) {
		return A, P, P
	}
	MrX.(*DFMITM2).m2 = func(B big.Int) big.Int {
		return B
	}
	MrX.(*DFMITM2).f = func(d *DFMITM2) {
		var x big.Int
		x.SetInt64(0)

		d.Hash1 = sha1.Sum(x.Bytes())
	}

	err := Alice.SendPGA(MrX)
	if err != nil {
		return err
	}

	err = Bob.SendB(MrX)
	if err != nil {
		return err
	}

	Alice.GenHash()
	Bob.GenHash()

	Alice.SendAESMsg(MrX)

	return nil
}

func SolveQ35Prob3() error {
	// g = p - 1

	var Alice, Bob, MrX exchange.DFClient

	Alice = new(exchange.DF)
	Bob = new(exchange.DF)
	MrX = new(DFMITM2)

	MrX.(*DFMITM2).D1 = Alice
	MrX.(*DFMITM2).D2 = Bob
	MrX.(*DFMITM2).m1 = func(A, P, G big.Int) (big.Int, big.Int, big.Int) {
		var One, Z big.Int
		One.SetInt64(1)
		Z.Sub(&P, &One)
		return A, P, Z
	}
	MrX.(*DFMITM2).m2 = func(B big.Int) big.Int {
		return B
	}
	MrX.(*DFMITM2).f = func(d *DFMITM2) {
		var x1, x2 big.Int
		x1.SetInt64(1)
		x2.Set(&d.P)
		x2.Sub(&x2, &x1)

		d.Hash1 = sha1.Sum(x1.Bytes())
		d.Hash2 = sha1.Sum(x2.Bytes())
	}

	err := Alice.SendPGA(MrX)
	if err != nil {
		return err
	}

	err = Bob.SendB(MrX)
	if err != nil {
		return err
	}

	Alice.GenHash()
	Bob.GenHash()

	Alice.SendAESMsg(MrX)

	return nil
}
