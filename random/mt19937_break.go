package random

import (
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
	y = y ^ (y >> _L)
	y = y ^ ((y << _T) & _C)
	mask := 0x7f
	for i := 0; i < 4; i++ {
		b := _B & uint32(mask<<uint32(7*(uint32(i)+1)))
		y = y ^ ((y << _S) & b)
	}
	for i := 0; i < 3; i++ {
		y = y ^ (y >> _U)
	}
	return y
}

func CloneMT19937(original MT19937) MT19937 {
	var clone MT19937
	clone.init()

	for i := 0; i < 624; i++ {
		rVal := original.Uint32()
		revVal := ReverseBitFlip(rVal)
		clone.data[i] = revVal
	}
	clone.index = 624
	return clone
}
