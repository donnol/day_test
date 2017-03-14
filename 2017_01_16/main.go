package main

import (
	"fmt"
	"time"

	. "jdscript.com/day_test/2017_01_16/happenbefore"
)

func main() {
	time.Sleep(time.Second)
}

func init() {
	go func() {
		fmt.Println("goroutine: ", A)
	}()

	// go func(a int) {
	// 	fmt.Println(a)
	// }(a)

	func() {
		fmt.Println("normal: ", A)
	}()
}

// func test() {
// 	go func() {
// 		fmt.Println("goroutine: ", a) // 1
// 	}()

// 	// go func(a int) {
// 	// 	fmt.Println(a)
// 	// }(a)

// 	func() {
// 		fmt.Println("normal: ", a) // 0
// 	}()

// 	a = 1
// }
