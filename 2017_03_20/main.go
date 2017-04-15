package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 25379.0 - 10.0
	b := 25379.0
	c := a / b
	fmt.Println(c)

	c100 := c * 100

	rateStr := strconv.FormatFloat(c100, 'f', 1, 64) + "%"
	fmt.Println(rateStr)

	rateStr2 := fmt.Sprintf("%.1F%%\n", c100)
	fmt.Println(rateStr2)

	rateStr3 := fmt.Sprintf("%.2f%%\n", c100)
	fmt.Println(rateStr3)

	rateStr4 := fmt.Sprintf("%.1F%%\n", 99.89)
	fmt.Println(rateStr4)
}
