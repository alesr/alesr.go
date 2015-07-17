package findAndSumAllMultiples

import "testing"

func TestFindAndSumAllMultiples(t *testing.T) {
	var r, expected uint32 = FindAndSumAllMultiples(1000), 233168
	if r != expected {
		t.Error("Expected ", expected, " got: ", r)
	}

}
