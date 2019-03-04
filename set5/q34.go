package set5

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"matasano/cipher"
	"matasano/exchange"
	"matasano/types"
	"math/big"
)

type DFMITM struct {
	G, P big.Int
	A, B big.Int
	Key  big.Int
	Hash [20]byte
	D1   exchange.DFClient
	D2   exchange.DFClient
}

func (d *DFMITM) SetBPG(B, P, G big.Int) {
	d.A = B
	d.P = P
	d.G = G
	d.D2.SetBPG(P, P, G)
}

func (d *DFMITM) SetB(B big.Int) {
	d.B = B
	d.D1.SetB(d.P)

	// We have A, B, p, G lets generate the hash
	var a *big.Int
	a, _ = rand.Int(rand.Reader, &d.P)

	d.Key.Exp(&d.P, a, &d.P)
	d.Hash = sha1.Sum(d.Key.Bytes())
}

func (d *DFMITM) SendPGA(d2 exchange.DFClient) error {
	return nil
}

func (d *DFMITM) SendB(d2 exchange.DFClient) error {
	return nil
}

func (d *DFMITM) GenHash() {

}

func (d *DFMITM) GetHash() []byte {
	return []byte{}
}

func (d *DFMITM) SendAESMsg(d2 exchange.DFClient) error {
	return nil
}

func (d *DFMITM) EchoMsg(msg []byte, d2 exchange.DFClient) error {
	origMsg := make([]byte, len(msg))
	copy(origMsg, msg)

	origIV := msg[len(msg)-16:]
	origMsgEncoded := msg[:len(msg)-16]

	var c cipher.CBC
	err := c.Init(d.Hash[:16], 16*8, origIV)
	if err != nil {
		return err
	}

	decoded, err := c.Decode(&types.Text{T: origMsgEncoded})
	if err != nil {
		return err
	}
	fmt.Println("Intercepted message", string(decoded))

	d.D2.EchoMsg(origMsg, d)

	return nil
}

func (d *DFMITM) ReceiveMsg(msg []byte) error {
	origMsg := make([]byte, len(msg))
	copy(origMsg, msg)

	origIV := msg[len(msg)-16:]
	origMsgEncoded := msg[:len(msg)-16]

	var c cipher.CBC
	err := c.Init(d.Hash[:16], 16*8, origIV)
	if err != nil {
		return err
	}

	decoded, err := c.Decode(&types.Text{T: origMsgEncoded})
	if err != nil {
		return err
	}
	fmt.Println("Intercepted echo", string(decoded))

	d.D1.ReceiveMsg(origMsg)

	return nil
}

func SolveQ34() error {
	var Alice, Bob, MrX exchange.DFClient

	Alice = new(exchange.DF)
	Bob = new(exchange.DF)
	MrX = new(DFMITM)

	MrX.(*DFMITM).D1 = Alice
	MrX.(*DFMITM).D2 = Bob

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
