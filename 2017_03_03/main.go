package main

import (
	"fmt"
	"plugin"
)

func main() {
	// 插件目录
	p, err := plugin.Open("./plugin/plugin.so")
	if err != nil {
		panic(err)
	}

	// 加载插件方法
	add, err := p.Lookup("Add") // 方法名必须与插件的方法名一致，区分大小写
	if err != nil {
		panic(err)
	}
	sub, err := p.Lookup("Subtract")
	if err != nil {
		panic(err)
	}

	// 运行插件方法
	sum := add.(func(int, int) int)(11, 2)
	fmt.Println(sum)

	subt := sub.(func(int, int) int)(11, 2)
	fmt.Println(subt)
}
