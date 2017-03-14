// 写出下面程序的输出及简要解释
package main

import (
	"fmt"
)

func main() {
	test1()
	test2()
	test3()
	test4()
	fmt.Println("test5: ", test5())
	fmt.Println("test6: ", test6())
	fmt.Println("test7: ", test7())
	fmt.Println("test8: ", test8())
}

func test1() {
	defer a()
	defer b()
	fmt.Println("test1")
}

// test1,b,a

func test2() {
	defer func() {
		defer a()
		defer b()
		fmt.Println("test2.defer")
	}()

	fmt.Println("test2")
}

// test2,test2.defer,b,a

func test3() {
	i := 0
	defer func() {
		fmt.Println(i)
	}()
	i++
	fmt.Println("test3")
}

// test3,1

func test4() {
	i := 0
	defer func(i int) {
		fmt.Println(i)
	}(i)
	i++
	fmt.Println("test4")
}

// test4,0

func test5() int {
	i := 0
	defer func() {
		i++
	}()
	return i
}

// 1 -> 0

func test6() int {
	i := 0
	defer func(i int) {
		i++
	}(i)
	return i
}

// 0

func test7() (i int) {
	defer func() {
		i++
	}()
	return
}

// 1

func test8() (i int) {
	defer func(i int) {
		i++
	}(i)
	return
}

// 0

func a() {
	fmt.Println("a")
}

func b() {
	fmt.Println("b")
}
