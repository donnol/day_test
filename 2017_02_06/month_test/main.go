package main

import (
	"fmt"
	"time"
)

func getNum(min, max int) {
	if min > max {
		return
	}
	for i := min; i < max; i++ {
		// fmt.Println(i)
		isNum := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isNum = false
				break
			}
		}
		if isNum {
			fmt.Println(i)
		}
	}
}

func main() {
	// 取素数
	getNum(100, 200)

	// 判断某个日期在当年内的第几天
	loc := time.Now().Location()
	fmt.Println(getDay(time.Date(2017, 02, 06, 0, 0, 0, 0, loc)))

	//
	// var a = 3
	// b :=  a+=a-=a*a
	// fmt.Println(b)
}

func getDay(day time.Time) int {
	return day.YearDay()
}
