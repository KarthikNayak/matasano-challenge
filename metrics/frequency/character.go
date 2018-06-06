package frequency

const correction = 5

var frequencyMap = map[rune]float64 {
	'a': 8.167, 'A': 8.167,
	'b': 1.492, 'B': 1.492,
	'c': 2.782, 'C': 2.782,
	'd': 4.253, 'D': 4.253,
	'e': 12.702, 'E': 12.702,
	'f': 2.228, 'F': 2.228,
	'g': 2.015, 'G': 2.015,
	'h': 6.094, 'H': 6.094,
	'i': 6.966, 'I': 6.966,
	'j': 0.153, 'J': 0.153,
	'k': 0.772, 'K': 0.772,
	'l': 4.025, 'L': 4.025,
	'm': 2.406, 'M': 2.406,
	'n': 6.749, 'N': 6.749,
	'o': 7.507, 'O': 7.507,
	'p': 1.929, 'P': 1.929,
	'q': 0.095, 'Q': 0.095,
	'r': 5.987, 'R': 5.987,
	's': 6.327, 'S': 6.327,
	't': 9.056, 'T': 9.056,
	'u': 2.758, 'U': 2.758,
	'v': 0.978, 'V': 0.978,
	'w': 2.360, 'W': 2.360,
	'x': 0.150, 'X': 0.150,
	'y': 1.974, 'Y': 1.974,
	'z': 0.074, 'Z': 0.074,
	' ': 0.0,
}

type CharacterFrequency struct {
}

func (c *CharacterFrequency) GetFrequency(b []byte) float64 {
	sum := 0.0
	for _, r := range b {
		if val, ok := frequencyMap[rune(r)]; ok {
			sum += val
		} else {
			sum -= correction
		}
	}
	return sum
}
