package random

import (
	"fmt"
	"math/rand"
	"time"
)

func randomTimeForSleep() time.Duration {
	return time.Duration(rand.Intn(960)+40) * time.Second
}

func FindSeed(val uint32) uint32 {
	t := time.Now().Unix()
	for i := t - 3600; i <= t; i++ {
		var m MT19937
		//		time.Sleep(randomTimeForSleep())
		m.Seed(uint32(i))

		//		time.Sleep(randomTimeForSleep())
		x := m.Uint32()
		if x == val {
			return uint32(i)
		}
	}
	return 0
}

func ReverseBitFlip(y uint32) uint32 {
	fmt.Println("before:", y)
	y = y ^ (y >> _L)
	fmt.Println("1:", y)
	y = y ^ ((y << _T) & _C)

	for i := 0; i < 8; i++ {
		maskB := _B & 0b11111
		y = y ^ ((y << _S) & _B)
	}
	fmt.Println("3:", y)
	for i := 0; i < 3; i++ {
		y = y ^ (y >> _U)
	}
	fmt.Println("after:", y)
	return y
}
