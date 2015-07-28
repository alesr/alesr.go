package recursiveFibonacci

import "testing"

var testCases = []struct {
	n, expected int
}{
	{1, 1},
	{7, 13},
	{25, 75025},
	{13, 233},
	{11, 89},
}

func TestRecursiveFibonacci(t *testing.T) {
	for _, test := range testCases {
		observed := RecursiveFibonacci(test.n)
		if observed != test.expected {
			t.Errorf("For n = %d, expected %t. Got %t.",
				test.n, test.expected, observed)
		}
	}
}

func BenchmarkRecursiveFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			RecursiveFibonacci(test.n)
		}
	}
}
