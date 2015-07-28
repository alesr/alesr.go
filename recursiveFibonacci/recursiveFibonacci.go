package recursiveFibonacci

func RecursiveFibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
	}
}
