package primeFactor

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	p        int
	expected []int
}{
	{15, []int{3, 5}},
	{26, []int{2, 13}},
	{37, []int{37}},
	{42, []int{2, 3, 7}},
}

func TestPrimeFactor(t *testing.T) {
	for _, test := range testCases {
		observed := PrimeFactor(test.p)
		// if len(observed) != len(test.expected) {
		// 	t.Error("For p = %d, expected %t. Got %t.",
		// 		test.p, test.expected, observed)
		// }
		if !reflect.DeepEqual(observed, test.expected) {
			t.Error("For p = %d, expected %t. Got %t.",
				test.p, test.expected, observed)
		}
	}
}
