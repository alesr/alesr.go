package removeDuplicate

import "testing"

func TestRemoveDuplicate(t *testing.T) {
	var r []string
	r = RemoveDuplicate([]string{"A", "B", "C", "C", "D", "E", "E", "F"})
	expected := []string{"A", "B", "J", "D", "E", "F"}
	if len(r) != len(expected) {
		t.Error("Expected \"[A B C D E F]\", got ", r)
	}

	for i := range r {
		if r[i] != expected[i] {
			t.Error("Expexted ", expected[i], "got ", r[i])
		}
	}
}
