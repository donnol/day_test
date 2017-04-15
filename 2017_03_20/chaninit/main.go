// 多个goroutine初始化顺序
package main

import "fmt"

var c = make(chan bool)

func main() {
	go printName("jd", c)
	go printName("fish", c)

	for {
		select {
		case single := <-c:
			fmt.Println(single)
		}
	}
}

func printName(name string, c chan bool) {
	fmt.Println(name)
	c <- true
}
