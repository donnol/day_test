package main

import (
	"fmt"
)

func main() {
	var n = 0
	i := &n
	// var i *int = new(int)
	fmt.Println("before", *i)
	add(i)
	fmt.Println("after", *i)
}

func add(i *int) {
	*i++
}
