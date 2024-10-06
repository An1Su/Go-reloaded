package piscine

func Capitalize(s string) string {
	var b []rune
	isFirst := true

	for _, a := range s {
		if a >= 'a' && a <= 'z' || a >= 'A' && a <= 'Z' || a >= '0' && a <= '9' {
			if a >= 'a' && a <= 'z' && isFirst {
				b = append(b, a-32)
				isFirst = false
			} else if a >= 'a' && a <= 'z' && !isFirst {
				b = append(b, a)
			} else if a >= 'A' && a <= 'Z' && isFirst {
				b = append(b, a)
				isFirst = false
			} else if a >= 'A' && a <= 'Z' && !isFirst {
				b = append(b, a+32)
			} else {
				b = append(b, a)
				isFirst = false
			}
		} else {
			isFirst = true
			b = append(b, a)
		}
	}
	return string(b)
}
