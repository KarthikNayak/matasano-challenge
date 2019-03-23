package set6

import (
	"crypto/sha1"
	"fmt"
	"matasano/keys"
	"math/big"
)

const (
	msg = `For those that envy a MC it can be hazardous to your health
So be friendly, a matter of life and death, just like a etch-a-sketch
`
	y           = "84ad4719d044495496a3201c8ff484feb45b962e7302e56a392aee4abab3e4bdebf2955b4736012f21a08084056b19bcd7fee56048e004e44984e2f411788efdc837a0d2e5abb7b555039fd243ac01f0fb2ed1dec568280ce678e931868d23eb095fde9d3779191b8c0299d6e07bbb283e6633451e535c45513b2d33c99ea17"
	r           = "548099063082341131477253921760299949438196259240"
	s           = "857042759984254168557880549501802188789837994940"
	expectedSha = "0954edd5e0afe5542a4adf012611a91912a3ec16"
)

var (
	n1 = big.NewInt(1)
)

func msgSha() []byte {
	s := sha1.Sum([]byte(msg))
	return s[:]
}

func SolveQ43() bool {
	_, Q, _ := keys.DSAGeneratePQG()
	// Y, _ := new(big.Int).SetString(y, 16)

	h := fmt.Sprintf("%x", msgSha())
	H, _ := new(big.Int).SetString(h, 16)

	R, _ := new(big.Int).SetString(r, 10)
	S, _ := new(big.Int).SetString(s, 10)

	RInv := new(big.Int).ModInverse(R, Q)

	k, x := big.NewInt(1), new(big.Int)
	for k.BitLen() <= 16 {
		x.Mul(S, k).Mod(x, Q).Sub(x, H).Mul(x, RInv).Mod(x, Q)

		sh := fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%x", x.Bytes()))))
		if sh == expectedSha {
			fmt.Println(x)
			return true
		}

		k.Add(k, n1)
	}
	return false
}
