package main

import "testing"

func BenchmarkLuhn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Luhn("49927398716")
	}
}

func BenchmarkLuhnParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Luhn("49927398716")
		}
	})
}

func TestLunh(t *testing.T) {
	if !Luhn("001230147647009683210024") {
		t.Errorf("Checksum failed for valid unput")
	}

	if Luhn("1234567812345678") {
		t.Errorf("Checksum passed for invalid unput")
	}
}
