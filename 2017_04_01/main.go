package main

import "fmt"
import "reflect"

const (
	a     = 10
	d int = 30
)

func main() {
	var b uint8
	b = 20
    m := a + b
	fmt.Println(m)
	fmt.Println(reflect.TypeOf(a))

	c := 0x7f
	fmt.Println(reflect.TypeOf(c))
	n := b + c
	fmt.Println(n)

    l := b + d
	fmt.Println(l)
}
