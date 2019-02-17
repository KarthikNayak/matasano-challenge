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
