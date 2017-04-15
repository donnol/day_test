package main

import "fmt"

func main() {
	testCase := []int{1, 2, 3, 4}
	// testLoopBreak(testCase)
	testLoopContinue(testCase)
}

func testLoopBreak(arr []int) {
	for _, single := range arr {
		fmt.Println(single)
		for {
			fmt.Println(single)
			if single != 4 {
				break
			}
		}
	}
}

func testLoopContinue(arr []int) {
	for _, single := range arr {
		fmt.Println(single)
		for i := 0; i < 2; i++ {
			fmt.Println(single)
			if single != 4 {
				continue
			}
			fmt.Println(single)
		}
	}
}
