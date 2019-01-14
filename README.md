## luhn check-digit mod10 algorithm

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/luhn)](https://goreportcard.com/report/jancajthaml-go/luhn)

Algorithm will detect any single-digit error, as well as almost all transpositions of adjacent digits. It will not, however, detect transposition of the two-digit sequence 09 to 90 (or vice versa).

It is not intended to be a cryptographically secure hash function. It is mostly used for pre-flight credit card number validation as specified in [ISO/IEC 7812-1:2015](http://www.iso.org/iso/catalogue_detail?csnumber=66011)

### Usage ###

```
import "github.com/jancajthaml-go/luhn"

ok := luhn.Validate("00123014764700968325")

digit, error := luhn.Digit("x")

checksum := luhn.Generate("1")
```

### Performance ###

```
BenchmarkLuhnSmall-4          200000000  13.9 ns/op  0 B/op  0 allocs/op
BenchmarkLuhnLarge-4          50000000   31.8 ns/op  0 B/op  0 allocs/op
BenchmarkLuhnSmallParallel-4  300000000  4.06 ns/op  0 B/op  0 allocs/op
BenchmarkLuhnLargeParallel-4  100000000  15.0 ns/op  0 B/op  0 allocs/op
```

verify your performance by running `make benchmark`

### Resources ###

* [Credit Card Validation - Check Digits](https://web.eecs.umich.edu/~bartlett/credit_card_number.html)
* [Wikipedia - Luhn algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm)
