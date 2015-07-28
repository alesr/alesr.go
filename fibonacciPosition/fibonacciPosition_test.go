package fibonacciPosition

import "testing"

var testCases = []struct {
	n, expected int
}{
	{1, 1},
	{5, 5},
	{7, 13},
	{10, 55},
	{19, 4181},
}

func TestFibonacciPosition(t *testing.T) {
	for _, test := range testCases {
		observed := FibonacciPosition(test.n)
		if observed != test.expected {
			t.Errorf("For n = %d, expected %t. Got %t.",
				test.n, test.expected, observed)
		}
	}
}

func BenchmarkFibonacciPosition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			FibonacciPosition(test.n)
		}
	}
}
