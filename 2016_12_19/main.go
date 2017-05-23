package main

import (
	"fmt"
)

type card struct {
	name string
}

type c card

func test(data card) {
	fmt.Println(data)
}

type s []card

func test2(data []card) {
	fmt.Println(data)
}

func main() {
	// data := c{}
	// test(data)

	var data2 s
	test2(data2)
}
