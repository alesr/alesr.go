package fibonacciGen

// FibonacciGen - output a sequence of Fibonacci numbers
func FibonacciGen() chan int {
	c := make(chan int)

	go func() {
		for i, j := 0, 1; ; i, j = i+j, i {
			c <- i
		}
	}()
	return c
}
