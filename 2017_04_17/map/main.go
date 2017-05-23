package main

import (
	"fmt"
)

type t_struct struct {
	player string
	id     int
}

func main() {
	dataA := make(map[string]t_struct)
	dataB := make(map[string]*t_struct)

	var playerA t_struct
	playerA.player = "tom"
	playerA.id = 1
	dataA["classA"] = playerA
	// dataA["classA"].id = 2 // ERROR, why?
	fmt.Println(dataA)

	playerB := new(t_struct)
	dataB["classB"] = playerB
	dataB["classB"].player = "rick"
	dataB["classB"].id = 3
	fmt.Println(dataB)
}
