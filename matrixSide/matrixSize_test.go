package matrixSide

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	n                           int
	expectedLeft, expectedRight []int
}{
	{5, []int{0, 5, 10, 15, 20}, []int{4, 9, 14, 19, 24}},
	{2, []int{0, 2}, []int{1, 3}},
	{3, []int{0, 3, 6}, []int{2, 5, 8}},
}

func TestSideValues(t *testing.T) {

	for _, test := range testCases {
		ls := LeftSide(test.n)
		rs := RightSide(test.n)

		fmt.Printf("n = %d\nFor left side got: %v\nExpecting: %v\nFor right, got: %v\nExpected: %v\n\n",
			test.n, ls, test.expectedLeft, rs, test.expectedRight)

	}
}
