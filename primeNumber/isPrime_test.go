package primes

import "testing"

var testCase = []struct {
	n        int
	expected bool
}{
	{1, false},
	{2, true},
	{4, false},
	{13, true},
	{34, false},
	{3, true},
	{52, false},
	{101, true},
	{421, true},
	{502, false},
	{338, false},
	{1871, true},
	{601, true},
	{9781, true},
	{5843, true},
	{8888, false},
}

func TestIsPrime(t *testing.T) {
	for _, test := range testCase {
		observed := IsPrime(test.n)
		if observed != test.expected {
			t.Errorf("For n = %d, expected %t. Got %t.",
				test.n, test.expected, observed)
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCase {
			IsPrime(test.n)
		}
	}
}
