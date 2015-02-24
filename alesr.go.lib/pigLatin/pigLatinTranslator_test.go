package pigLatin

import "testing"

var testCases = []struct {
	word, expected string
}{
	{"pig", "igpay"},
	{"banana", "ananabay"},
	{"duck", "uckday"},
}

func TestPigLatinTranslator(t *testing.T) {
	for _, test := range testCases {
		observed := PigLatinTranslator(test.word)
		if observed != test.expected {
			t.Errorf("For word %s, expected %s. Got %s",
				test.word, test.expected, observed)
		}
	}
}

func BenchmarkPigLatinTranslator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			PigLatinTranslator(test.word)
		}
	}
}
