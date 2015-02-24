package fibonacciPosition

// FibonacciPosition - find the Fibonacci value for input position
func FibonacciPosition(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		var f2, f1, i int = 0, 1, 2
		fn := f2 + f1
		for i < n {
			f2 = f1
			f1 = fn
			fn = f2 + f1
			i++
		}
		return fn
	}
}
