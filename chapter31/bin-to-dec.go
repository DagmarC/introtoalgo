package chapter31

func BinToDec(bin string) int {
	return convert(bin, 0, len(bin)-1)
}

// convert binary to decimal - rereq: bin string contains only 1 and 0
func convert(bin string, p, r int) int {
	if p > r {
		return -1000
	}
	if p == r {
		// from the end of the string
		if bin[len(bin)-p-1] == '0' {
			return 0
		} else {
			return 1 << p // shift
		}
	}
	q := (p + r) / 2
	return convert(bin, p, q) + convert(bin, q+1, r)
}
