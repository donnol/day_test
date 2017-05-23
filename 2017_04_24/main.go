package main

import "fmt"
import "os"
import "runtime/pprof"

func main() {
	f, err := os.Create("profile_file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	pprof.StopCPUProfile()

	goTest()
}

func goTest() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	work := make(chan int)
	sum := 0
	aLen := len(a)

	// 开启三个goroutine
	for i := 0; i < 3; i++ {
		go func() {
			for single := range work {
				fmt.Println(single)
				sum += single
				aLen--
			}
		}()
	}

	for _, value := range a {
		work <- value
	}
	for aLen > 0 {

	}
	fmt.Println(sum)
}
