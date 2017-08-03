package main

import "os"

var m = [...]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

/**
 * Luhn Checksum algorithm
 *
 * @see https://en.wikipedia.org/wiki/Luhn_algorithm
 *
 * @author jan.cajthaml
 */
func Luhn(cc string) (ok bool) {
	var (
		i int = len(cc) - 1
		x int = 0
	)

	for i >= 0 {
		switch i & 1 {
		case 1:
			// [0 => 0]
			x += int(cc[i]) - 48
		default:
			// lookup [0..4 => x<<1], [5..8 => (x<<1)+1]
			x += m[int(cc[i])]
		}
		i--
	}

	// int32 mod 10
	return (x - (x/10)*10) == 0
}

func main() {
	defer func() {
		if recover() != nil {
			os.Exit(1)
		}
	}()

	if len(os.Args) != 2 {
		os.Stderr.Write([]byte("Usage           : ./luhn <input>\nValid Example   : ./luhn 123; echo \"$?\"\nInvalid Example : ./luhn 12; echo \"$?\"\n"))
		return
	}

	if !Luhn(os.Args[1]) {
		os.Exit(1)
		return
	}

	os.Exit(0)
}
