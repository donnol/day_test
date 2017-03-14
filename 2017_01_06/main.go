// 现有一个类A，除了使用它的实例调用它的方法Print--a.Print()外，还能怎么调用呢？
// 写出一种即可
package main

import "fmt"

// 1
// type A struct {
// }

// func (a *A) Print() {
// 	fmt.Println("this is A")
// }

// func main() {
// 	a := A{}
// 	a.Print()
// }

// 2
type A struct {
}

func (a *A) Print() {
	fmt.Println("this is A")
}

func (a A) Hello() {
	fmt.Println("hello, this is A")
}

func (a *A) PrintName(name string) {
	fmt.Println("name: ", name)
}

func main() {
	a := A{}

	//     ./main.go:32: invalid method expression A.Print (needs pointer receiver: (*A).Print)
	// ./main.go:32: A.Print undefined (type A has no method Print)
	// A.Print(&a)
	// A.Print(a)

	(*A).Print(&a)

	(*A).PrintName(&a, "a")

	//./main.go:33: cannot use a (type A) as type *A in argument to (*A).Hello
	// (*A).Hello(a)

	A.Hello(a)

	(*A).Hello(&a)
}

// 3
// type I interface {
// 	Print()
// }
// type A struct {
// }

// func (a *A) Print() {
// 	fmt.Println("this is A")
// }

// type B struct {
// }

// func (b *B) Print() {
// 	fmt.Println("this is B")
// }

// func main() {
// 	a := A{}
// 	b := B{}
// 	I.Print(&a)
// 	I.Print(&b)
// }
