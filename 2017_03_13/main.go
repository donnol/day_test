package main

import (
	"fmt"
	// "errors"
)

func main() {
	// var err = errors.New("abc")
	var err error

	fmt.Println("begin")

    // 与if不同的地方在于，如果条件成立，大括号里面的语句会一直运行，直到条件不成立或break、return
	for err == nil {
		fmt.Println("wait")
	}

	fmt.Println("hello")
}