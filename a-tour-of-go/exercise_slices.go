package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	foo := make([][]uint8, dy)
	for x := range foo {
		foo[x] = make([]uint8, dx)
		for y := range foo[x] {
			foo[x][y] = uint8((x ^ y) * (x ^ y))
		}
	}
	return foo
}

func main() {
	pic.Show(Pic)
}
