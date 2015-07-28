package isMatrixSide

import "testing"

var testCase = []struct {
	value, size                 int
	expectedLeft, expectedRight bool
}{
	{3, 3, true, false},
	{7, 4, false, true},
	{2, 2, true, false},
	{3, 2, false, true},
	{1, 3, false, false},
	{8, 4, true, false},
	{24, 5, false, true},
	{10, 5, true, false},
}

func TestIsSideValue(t *testing.T) {
	for _, test := range testCase {
		leftResult := IsLeftSide(test.value, test.size)
		rightResult := IsRightSide(test.value, test.size)

		if test.expectedLeft != leftResult || test.expectedRight != rightResult {
			t.Errorf("error")
		}
	}
}

func BenchmarkIsMatrixSide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCase {
			IsLeftSide(test.value, test.size)
			IsRightSide(test.value, test.size)
		}
	}
}
