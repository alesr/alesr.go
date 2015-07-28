package isMatrixSide

// IsLeftSide return true if i is on the left collumn
func IsLeftSide(i, n int) bool {
	flag := false
	for j := 0; j < n; j++ {
		if n*j == i {
			flag = true
			break
		}
	}
	return flag
}

// IsRightSide return true if i is on the right collumn
func IsRightSide(i, n int) bool {
	flag := false
	for j := 1; j <= n; j++ {
		if (n*j)-1 == i {
			flag = true
			break
		}
	}
	return flag
}
