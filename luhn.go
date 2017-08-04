package main

var m = [58]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

/**
 * Luhn Checksum algorithm
 *
 * @see https://en.wikipedia.org/wiki/Luhn_algorithm
 *
 * @author jan.cajthaml
 */
func Luhn(cc string) (ok bool) {
	var (
		i int  = len(cc) - 1
		x uint = 0
	)

scan:
	switch i & 1 {
	case 1:
		x += uint(cc[i]) - 48
	default:
		x += m[int(cc[i])]
	}
	if i == 0 {
		goto eos
	}
	i--
	goto scan

eos:
	return (x - (x/10)*10) == 0
}
