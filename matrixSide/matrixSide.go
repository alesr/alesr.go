package matrixSide

import "fmt"

// MatrixSide defines side values of a matrix nxn
type MatrixSide struct {
	Leftvalues, Rightvalues []int
}

func (m *MatrixSide) get() {
	fmt.Printf("Left values: %v\nRight values: %v\n", m.Leftvalues, m.Rightvalues)
}

// SideValues prints out the left and right values of a matrix of size n
func SideValues(n int) {
	m := &MatrixSide{
		Leftvalues:  LeftSide(n),
		Rightvalues: RightSide(n),
	}
	m.get()
}

// LeftSide return the left collumn
func LeftSide(n int) []int {
	ls := make([]int, 0, n)

	for i := 0; i < n; i++ {
		ls = append(ls, i*n)
	}
	return ls
}

// RightSide return the right collumn
func RightSide(n int) []int {
	rs := make([]int, 0, n)

	for i := 1; i <= n; i++ {
		rs = append(rs, ((n * i) - 1))
	}
	return rs
}
