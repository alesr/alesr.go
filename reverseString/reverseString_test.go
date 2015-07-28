package reverseString

import "testing"

func TestReverseString(t *testing.T) {
	var r string
	r = ReverseString("granny easter")
	if r != "retsae ynnarg" {
		t.Error("Expected \"retsae ynnarg\", got ", r)
	}
}
