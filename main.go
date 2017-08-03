package main

import "os"

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
