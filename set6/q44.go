package set6

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"matasano/keys"
	"math/big"
	"os"
)

type Q44Data struct {
	msg []byte
	s   *big.Int
	r   *big.Int
	m   *big.Int
}

func GetQ44Data() []Q44Data {
	fileName := "./q44_data.txt"
	file, _ := os.Open(fileName)
	defer file.Close()

	q := make([]Q44Data, 0)
	scanner := bufio.NewScanner(file)

	count := 0

	var n Q44Data
	for scanner.Scan() {
		t := scanner.Text()
		switch count {
		case 0:
			n.msg = []byte(t[5:])
		case 1:
			n.s, _ = new(big.Int).SetString(t[3:], 10)
		case 2:
			n.r, _ = new(big.Int).SetString(t[3:], 10)
		case 3:
			n.m, _ = new(big.Int).SetString(t[3:], 16)
			q = append(q, n)
		}
		count++
		count %= 4
	}

	return q
}

func GetIdenticalKs() (*Q44Data, *Q44Data) {
	d := GetQ44Data()
	l := len(d)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if d[i].r.Cmp(d[j].r) == 0 {
				return &d[i], &d[j]
			}
		}
	}
	return nil, nil
}

func Q44GetK(A, B *Q44Data, Q *big.Int) (k *big.Int) {
	k = new(big.Int)

	t := new(big.Int).Sub(A.s, B.s)
	t.Mod(t, Q)
	t.ModInverse(t, Q)

	k.Sub(A.m, B.m).Mod(k, Q).Mul(k, t).Mod(k, Q)

	return
}

func SolveQ44() bool {
	const (
		y = "2d026f4bf30195ede3a088da85e398ef869611d0f68f0713d51c9c1a3a26c95105d915e2d8cdf26d056b86b8a7b85519b1c23cc3ecdc6062650462e3063bd179c2a6581519f674a61f1d89a1fff27171ebc1b93d4dc57bceb7ae2430f98a6a4d83d8279ee65d71c1203d2c96d65ebbf7cce9d32971c3de5084cce04a2e147821"
	)

	A, B := GetIdenticalKs()

	_, Q, _ := keys.DSAGeneratePQG()
	// Y, _ := new(big.Int).SetString(y, 16)

	k := Q44GetK(A, B, Q)

	h := fmt.Sprintf("%x", sha1.Sum(A.msg))
	H, _ := new(big.Int).SetString(h, 16)
	RInv := new(big.Int).ModInverse(A.r, Q)

	x := new(big.Int)
	x.Mul(A.s, k).Mod(x, Q).Sub(x, H).Mul(x, RInv).Mod(x, Q)

	sh := fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%x", x.Bytes()))))
	if sh == "ca8f6f7c66fa362d40760d135b763eb8527d3d52" {
		return true
	}
	return false
}
