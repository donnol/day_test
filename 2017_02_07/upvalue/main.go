// 闭包捕获的外部变量，称为upvalue
package main

import (
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	log.Printf("%p", &wg)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			log.Printf("%p", wg)
			log.Printf("i: %d", i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	log.Println("exit")
}
