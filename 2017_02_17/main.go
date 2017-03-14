// 切片长度和容量
package main

import "fmt"

func main() {
	s := make([]byte, 5)
	s = s[2:4]
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
