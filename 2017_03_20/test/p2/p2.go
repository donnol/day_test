package p2

import (
	"fmt"
)

var a int

func Print() {
	fmt.Printf("%p\n", &a)
}
