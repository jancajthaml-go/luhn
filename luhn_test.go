package main

import "testing"

func BenchmarkLuhn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LuhnDigit("49927398716")
	}
}

func BenchmarkLuhnDigitParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			LuhnDigit("49927398716")
		}
	})
}

func TestLuhnDigit(t *testing.T) {
	if LuhnDigit("001230147647009683210024") != 0 {
		t.Errorf("Checksum failed for valid unput")
	}

	if LuhnDigit("1234567812345678") == 0 {
		t.Errorf("Checksum passed for invalid unput")
	}
}
