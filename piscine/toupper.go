package piscine

func ToUpper(s string) string {
	bigMap := map[rune]rune{
		'a': 'A',
		'b': 'B',
		'c': 'C',
		'd': 'D',
		'e': 'E',
		'f': 'F',
		'g': 'G',
		'h': 'H',
		'i': 'I',
		'j': 'J',
		'k': 'K',
		'l': 'L',
		'm': 'M',
		'n': 'N',
		'o': 'O',
		'p': 'P',
		'q': 'Q',
		'r': 'R',
		's': 'S',
		't': 'T',
		'u': 'U',
		'v': 'V',
		'w': 'W',
		'x': 'X',
		'y': 'Y',
		'z': 'Z',
	}
	var result []rune
	for _, r := range s {
		big, ok := bigMap[r]
		if ok {
			result = append(result, big)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
