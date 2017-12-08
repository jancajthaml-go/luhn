package main

import "testing"

func BenchmarkLuhnSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Digit("123")
	}
}

func BenchmarkLuhnLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Digit("00123014764700968325")
	}
}

func BenchmarkLuhnSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Digit("123")
		}
	})
}

func BenchmarkLuhnLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Digit("00123014764700968325")
		}
	})
}

func TestLuhn(t *testing.T) {
	expectSignature := func(a string, b string) {
		c, err := Generate(a)
		if err != nil {
			t.Errorf("%v", err)
		} else if b != c {
			t.Errorf("for input " + a + " expected signature " + b + " but got " + c + " instead")
		}

		if !Validate(c) {
			t.Errorf("Unable to validate signature that was generated")
		}
	}

	if !Validate("001230147647009683210024") {
		t.Errorf("Checksum failed for valid input")
	}

	if Validate("1234567812345678") {
		t.Errorf("Checksum passed for invalid input")
	}

	if Validate("xy-1") {
		t.Errorf("Checksum passed for invalid alphabet")
	}

	expectSignature("00123014764700968321002", "001230147647009683210024")

}
