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

/*
  main call example
  c := fibonacciGen.FibonacciGen()
  for n := 0; n < 12; n++ {
    fmt.Println(<-c)
  }
*/
