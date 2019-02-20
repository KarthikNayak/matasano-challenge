package random

import "fmt"

// Found at https://en.wikipedia.org/wiki/Mersenne_Twister
// w: word size (in number of bits)
// n: degree of recurrence
// m: middle word, an offset used in the recurrence relation defining the series x, 1 ≤ m < n
// r: separation point of one word, or the number of bits of the lower bitmask, 0 ≤ r ≤ w - 1
// a: coefficients of the rational normal form twist matrix
// b, c: TGFSR(R) tempering bitmasks
// s, t: TGFSR(R) tempering bit shifts
// u, d, l: additional Mersenne Twister tempering bit shifts/masks
const (
	_W, _N, _M, _R = 32, 624, 397, 31
	_A             = 0x9908B0DF
	_U, _D         = 11, 0xFFFFFFFF
	_S, _B         = 7, 0x9D2C5680
	_T, _C         = 15, 0xEFC60000
	_L             = 18
	_F             = 1812433253
)

type MT19937 struct {
	data      []uint32
	index     uint32
	lowerMask uint32
	upperMask uint32
}

func (m *MT19937) init() {
	m.data = make([]uint32, _N)
	m.index = _N + 1
	// 0x7fffffff (That is, the binary number of r 1's)
	m.lowerMask = (1 << _R) - 1
	// 0x80000000
	m.upperMask = (^m.lowerMask)
}

func (m *MT19937) Seed(seed uint32) {
	if len(m.data) == 0 {
		m.init()
	}
	m.index = _N
	m.data[0] = seed
	for i := 1; i < _N-1; i++ {
		// lowest w bits of (f * (MT[i-1] xor (MT[i-1] >> (w-2))) + i)
		m.data[i] = (uint32(_F) * (m.data[i-1] ^ (m.data[i-1] >> (_W - 2)))) + uint32(i)
	}
}

func (m *MT19937) twist() {
	for i := 0; i < _N-1; i++ {
		x := (m.data[i] & m.upperMask) + (m.data[(i+1)%_N] & m.lowerMask)
		xA := x >> 1
		if (x % 2) != 0 {
			xA = xA ^ _A
		}
		m.data[i] = m.data[(i+_M)%_N] ^ xA
	}
	m.index = 0
}

func (m *MT19937) Uint32() uint32 {
	if len(m.data) == 0 {
		m.init()
	}

	if m.index >= _N {
		if m.index > _N {
			m.Seed(5489)
		}
		m.twist()
	}

	y := m.data[m.index]
	fmt.Println("before:", y)
	y = y ^ ((y >> _U) & _D)
	fmt.Println("1:", y)
	y = y ^ ((y << _S) & _B)
	fmt.Println("2:", y)
	y = y ^ ((y << _T) & _C)
	fmt.Println("3:", y)
	y = y ^ (y >> _L)
	fmt.Println("after:", y)

	m.index = m.index + 1
	return y
}
