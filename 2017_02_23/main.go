// defer,panic,recover使用
package main

import "fmt"

func main() {
	// a
	// c
	// d
	// panic: 55

	// goroutine 1 [running]:
	// main.f()
	// 	/donnol/Project/Golang/src/jdscript.com/day_test/2017_02_23/main.go:40 +0xd1
	// main.panicDefer()
	// 	/donnol/Project/Golang/src/jdscript.com/day_test/2017_02_23/main.go:24 +0x3e
	// main.main()
	// 	/donnol/Project/Golang/src/jdscript.com/day_test/2017_02_23/main.go:7 +0x20
	// panicDefer() // 执行了defer，但是defer里面没有recover，所以最后还是panic了

	// a
	// c
	// 55
	// d
	// h
	panicDeferRecover() // 执行了defer，因为defer里面执行了recover方法，所以程序可以继续执行
	goOn()
}

func goOn() {
	fmt.Println("h")
}

func panicDefer() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		// if err:=recover();err!=nil{
		//     fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		// }
		fmt.Println("d")
	}()
	f()
}

func panicDeferRecover() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	f()
}

func f() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")
	fmt.Println("f")
}
