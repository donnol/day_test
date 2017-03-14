package main

import "fmt"

const (
	a = 1
	b = 6
	c
	d = 8
	// e int // const declaration cannot have type without expression
)

var (
	f = 1
	g = 6
	h int
	i = 8
	j int
)

func main() {
	fmt.Println(a, b, c, d)

	fmt.Println(f, g, h, i, j)
}
