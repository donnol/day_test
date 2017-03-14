package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now, now.Add(-time.Hour))
}
