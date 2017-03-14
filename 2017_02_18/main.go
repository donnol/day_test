package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []Person{
		Person{
			Name: "jd",
			Age:  18,
		},
		Person{
			Name: "fish",
			Age:  14,
		},
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].Name > data[j].Name
	})
	fmt.Println(data)

	sort.Slice(data, func(i, j int) bool {
		return data[i].Age > data[j].Age
	})
}

type Person struct {
	Name string
	Age  int
}
