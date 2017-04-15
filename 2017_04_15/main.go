package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 值类型变量初始化为零值
	var i int
	fmt.Printf("%#v\n", i) // 0
	var dz Dog
	fmt.Printf("%#v\n", dz) // {}

	// 引用类型变量初始化为nil
	var d *Dog
	fmt.Printf("%#v\n", d) // nil

	fmt.Println("=====")
	print0(dz)
	fmt.Println("=====")
	print1(d)
	fmt.Println("=====")
	print2(d)
}

type Stringer interface {
	String() string
}

type Dog struct {
	Name string
}

func (d *Dog) String() string {
	return d.Name
}

func print0(s Dog) {
	fmt.Printf("%#v\n", s) // (*main.Dog)(nil)
	sd := &s
	fmt.Println(reflect.TypeOf(sd))
	if sd == nil {
		fmt.Println("1 nil") // 1 nil
	} else {
		fmt.Printf("2 not nil: %+v\n", sd.String())
	}
}

func print1(s *Dog) {
	fmt.Printf("%#v\n", s) // (*main.Dog)(nil)
	fmt.Println(reflect.TypeOf(s))
	if s == nil {
		fmt.Println("1 nil") // 1 nil
	} else {
		fmt.Printf("2 not nil: %+v\n", s.String())
	}
}

// 接口参数的nil必须是(nil, nil)才与nil相等
// 当参数是结构体时，这个参数就不再是(nil,nil)，而是(*struct{}, nil)了，所以与nil不相等
func print2(s Stringer) {
	fmt.Printf("%#v\n", s) // (*main.Dog)(nil)
	fmt.Println(reflect.TypeOf(s))
	if s == nil {
		fmt.Println("1 nil") //
	} else {
		fmt.Printf("2 not nil: %+v\n", s.String()) // panic: runtime error: invalid memory address or nil pointer dereference
	}
}
