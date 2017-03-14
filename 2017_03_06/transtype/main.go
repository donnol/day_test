package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 0
	var str = string(i)
	fmt.Println("'" + str + "'") // 空字符串 ''

	str2 := strconv.Itoa(i)
	fmt.Println("'" + str2 + "'") // '0'

	var j int = 1
	var str3 = string(j)
	fmt.Println("'" + str3 + "'") // 空字符串 ''

	str4 := strconv.Itoa(j)
	fmt.Println("'" + str4 + "'") // '1'
}
