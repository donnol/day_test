// 比较golang struct的receiver
package main

import "fmt"

func main() {
	//
	a1 := A{
		"jd",
	}
	fmt.Println(a1.Name)
	a1.SetName()
	fmt.Println(a1.Name)
	a1.SetNameByPtr()
	fmt.Println(a1.Name)

	fmt.Println("=====")

	//
	a2 := &A{
		"jd",
	}
	fmt.Println(a2.Name)
	a2.SetNameByPtr()
	fmt.Println(a2.Name)
	a2.SetName()
	fmt.Println(a2.Name)

	fmt.Println("=====")

	A{"jd"}.SetName()

	// cannot call pointer method on composite literal
	// cannot take the address of composite literal
	// 		A{"jd"}.SetNameByPtr()

	fmt.Println("=====")

	var a3 A
	a3.SetName()

	a3.SetNameByPtr()
}

type A struct {
	Name string
}

// 当receiver是值的时候，无法改变Name的值
func (a A) SetName() {
	// if a == nil { // cannot convert nil to type A
	// 	fmt.Println("nil a")
	// }
	a.Name = "fish"
	fmt.Println(a.Name)
}

// 当receiver是指针的时候，可以改变Name的值
func (a *A) SetNameByPtr() {
	if a == nil {
		fmt.Println("nil *a")
	}
	a.Name = "fish"
	fmt.Println(a.Name)
}
