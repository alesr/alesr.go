package pigLatin

import "testing"

var testCases = []struct {
	word, expected string
}{
	{"pig", "igpay"},
	{"banana", "ananabay"},
	{"duck", "uckday"},
}

func TestPigTranslator(t *testing.T) {
	for _, test := range testCases {
		observed := PigTranslator(test.word)
		if observed != test.expected {
			t.Errorf("For word %s, expected %s. Got %s",
				test.word, test.expected, observed)
		}
	}
}

func BenchmarkPigTranslator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			PigTranslator(test.word)
		}
	}
}
