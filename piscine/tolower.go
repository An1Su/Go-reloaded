package piscine

func ToLower(s string) string {
	Lowmap := map[rune]rune{
		'A': 'a',
		'B': 'b',
		'C': 'c',
		'D': 'd',
		'E': 'e',
		'F': 'f',
		'G': 'g',
		'H': 'h',
		'I': 'i',
		'J': 'j',
		'K': 'k',
		'L': 'l',
		'M': 'm',
		'N': 'n',
		'O': 'o',
		'P': 'p',
		'Q': 'q',
		'R': 'r',
		'S': 's',
		'T': 't',
		'U': 'u',
		'V': 'v',
		'W': 'w',
		'X': 'x',
		'Y': 'y',
		'Z': 'z',
	}
	var box []rune
	for _, v := range s {
		low, ok := Lowmap[v]
		if ok {
			box = append(box, low)
		} else {
			box = append(box, v)
		}
	}
	return string(box)
}
