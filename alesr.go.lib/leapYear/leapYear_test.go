package leapYear

import "testing"

var testCases = []struct {
	year     int
	expected bool
}{
	{1996, true},
	{1997, false},
	{1900, false},
	{2400, true},
}

func TestIsLeapYear(t *testing.T) {
	for _, test := range testCases {
		observed := IsLeapYear(test.year)
		if observed != test.expected {
			t.Errorf("For year %d expected %t, got %t.",
				test.year, test.expected, observed)
		}
	}
}

func BenchmarkIsLeapYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			IsLeapYear(test.year)
		}
	}
}
