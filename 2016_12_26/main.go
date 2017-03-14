package main

import (
	"fmt"
	"sync"
)

// 变量
var a = 1

// 会话一
func sessionAdd() {
	for i := 0; i < 100; i++ {
		a += i
	}
}

// 会话二
func sessionSub() {
	for i := 0; i < 100; i++ {
		a -= i
	}
}

// 会话三
func sessionPrint() {
	fmt.Println(a)
}

func main() {
	var wg = new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		defer wg.Done()
		sessionAdd()
	}()
	go func() {
		defer wg.Done()
		sessionSub()
	}()
	go func() {
		defer wg.Done()
		sessionPrint()
	}()
	wg.Wait()

	sessionPrint()
}
