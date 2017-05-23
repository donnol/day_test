package main

import "fmt"

func main() {
	data := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(data) //

	data2 := data[:3]
	data3 := data[4:]
	data = append(data2, data3...)
	fmt.Println(data) // a, b, c, d, e, f

	left := data[:3]
	right := data[3:]
	fmt.Println(left)  // a, b, c
	fmt.Println(right) // d, e, f

	data = append(left, "d")
	fmt.Println(left)  // a, b, c, d
	fmt.Println(right) // d, e, f
}
