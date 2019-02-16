package helpers

const correction = 5

var frequencyMap = map[rune]float64{
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

var CanonicalFreq = map[byte]float64{
	' ': 0.00189169, '!': 0.0306942,
	'"': 0.000183067, '#': 0.00854313,
	'$': 0.00970255, '%': 0.00170863,
	'&': 0.00134249, '\'': 0.000122045,
	'(': 0.000427156, ')': 0.00115942,
	'*': 0.0241648, '+': 0.00231885,
	',': 0.00323418, '-': 0.0197712,
	'.': 0.0316706, '/': 0.00311214,
	'0': 2.74381, '1': 4.35053,
	'2': 3.12312, '3': 2.43339,
	'4': 1.94265, '5': 1.88577,
	'6': 1.75647, '7': 1.621,
	'8': 1.66225, '9': 1.79558,
	':': 0.000549201, ';': 0.00207476,
	'<': 0.000427156, '=': 0.00140351,
	'>': 0.000183067, '?': 0.00207476,
	'@': 0.0238597, 'A': 0.130466,
	'B': 0.0806715, 'C': 0.0660872,
	'D': 0.0698096, 'E': 0.0970865,
	'F': 0.0417393, 'G': 0.0497332,
	'H': 0.0544319, 'I': 0.070908,
	'J': 0.0363083, 'K': 0.0460719,
	'L': 0.0775594, 'M': 0.0782306,
	'N': 0.0748134, 'O': 0.0729217,
	'P': 0.073715, 'Q': 0.0147064,
	'R': 0.08476, 'S': 0.108132,
	'T': 0.0801223, 'U': 0.0350268,
	'V': 0.0235546, 'W': 0.0320367,
	'X': 0.0142182, 'Y': 0.0255073,
	'Z': 0.0170252, '[': 0.0010984,
	'\\': 0.00115942, ']': 0.0010984,
	'^': 0.00195272, '_': 0.0122655,
	'`': 0.00115942, 'a': 7.52766,
	'b': 2.29145, 'c': 2.57276,
	'd': 2.76401, 'e': 7.0925,
	'f': 1.2476, 'g': 1.85331,
	'h': 2.41319, 'i': 4.69732,
	'j': 0.836677, 'k': 1.96828,
	'l': 3.77728, 'm': 2.99913,
	'n': 4.56899, 'o': 5.17,
	'p': 2.45578, 'q': 0.346119,
	'r': 4.96032, 's': 4.61079,
	't': 3.87388, 'u': 2.10191,
	'v': 0.833626, 'w': 1.24492,
	'x': 0.573305, 'y': 1.52483,
	'z': 0.632558, '{': 0.000122045,
	'|': 0.000122045, '}': 6.10223e-0,
	'~': 0.00152556, 'ä': 6.10223e-05,
	'æ': 0.000183067, 'ö': 6.10223e-05,
	'ü': 0.000122045,
}

func CharacterFrequency(b []byte) float64 {
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
