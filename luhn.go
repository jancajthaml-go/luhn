package main

var m = [58]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

/**
 * Luhn check-digit algorithm
 *
 * @see https://en.wikipedia.org/wiki/Luhn_algorithm
 *
 * @author jan.cajthaml
 */
func LuhnDigit(cc string) int {
	var (
		i int  = len(cc) - 1
		x uint = 0
	)

loop:
	if i == 0 {
		return int(x - (x/10)*10)
	}

	switch i & 1 {
	case 1:
		x += uint(cc[i]) - 48
	default:
		x += m[int(cc[i])]
	}
	i--
	goto loop
}
