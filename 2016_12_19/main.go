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

func main() {
	data := c{
		name: "edward",
	}
	test(data)
}
