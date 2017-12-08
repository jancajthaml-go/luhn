// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Luhn_algorithm
//
package main

import "errors"

var m = [10]uint{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// Digit returns luhn digit for given numeric string
func Digit(cc string) (int, error) {
	var (
		i int = len(cc) - 1
		x uint
		d uint
	)

loop:
	if i < 0 {
		return int(x - (x/10)*10), nil
	}

	d = uint(cc[i]) - 48
	if d > 9 {
		return 1, errors.New("string must contain only digits")
	}

	switch i & 1 {
	case 1:
		x += d
	default:
		x += m[d]
	}
	i--
	goto loop
}

// Validate returns true if numeric string is signed with valid luhn digit
func Validate(cc string) (ok bool) {
	digit, err := Digit(cc)
	return err == nil && digit == 0
}
