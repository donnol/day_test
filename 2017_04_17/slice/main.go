package main

import (
	"fmt"
)

type t_struct struct {
	player string
	id     int
}

func main() {
	dataA := make([]t_struct, 1)
	dataB := make([]*t_struct, 1)

	var playerA t_struct
	playerA.player = "tom"
	playerA.id = 1
	dataA[0] = playerA
	dataA[0].id = 2 // ERROR, why?
	fmt.Println(dataA)

	playerB := new(t_struct)
	dataB[0] = playerB
	dataB[0].player = "rick"
	dataB[0].id = 3
	fmt.Println(dataB)
}
