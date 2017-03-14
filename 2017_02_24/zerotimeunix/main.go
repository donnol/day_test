package main

import (
	"fmt"
	"time"
)

func main() {
	zeroTime := time.Time{}

	fmt.Println(zeroTime, zeroTime.Unix())
}
