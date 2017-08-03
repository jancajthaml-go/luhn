package main

import (
	"os"
)

var noop = []byte{}

var m = [...]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

func Luhn(cc string) (ok bool) {
	var (
		i int = len(cc) - 1
		x int = 0
	)

	for i >= 0 {
		switch i & 1 {
		case 1:
			x += int(cc[i]) - 48
		default:
			x += m[int(cc[i])]
		}
		i--
	}

	return (x - (x/10)*10) == 0
}

func main() {
	defer func() {
		if recover() != nil {
			os.Stderr.Write(noop)
		}
	}()

	if len(os.Args) != 2 {
		os.Stderr.Write([]byte("Usage `luhn <input>`"))
		return
	}

	if !Luhn(os.Args[1]) {
		os.Stdout.Write(noop)
		return
	}

	os.Stdout.Write(noop)
}
