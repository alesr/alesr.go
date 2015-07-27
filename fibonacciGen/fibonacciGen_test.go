package fibonacciGen

import "testing"

func TestFibonacciGen(t *testing.T) {
	var result int
	var c = FibonacciGen()
	for n := 0; n < 12; n++ {
		result = <-c
	}
	if result != 89 {
		t.Error("For n = 12, expected 89. Got %t.", result)
	}
}

/*
main call example
c := fibonacciGen.FibonacciGen()
for n := 0; n < 12; n++ {
fmt.Println(<-c)
}
*/
