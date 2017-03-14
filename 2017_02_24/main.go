package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["name"] = "fish"

	modMap(m)

	fmt.Println(m)
}

func modMap(m map[string]string) {
	// 删除
	// delete(m, "name")

	// 修改
	// m["name"] = "jd"

	// 添加
	m["age"] = "18"
}

// 运行结果
// ➜  2017_02_24 git:(master) ✗ go run  main.go
// # command-line-arguments
// ./main.go:8: undefined: fish
// ➜  2017_02_24 git:(master) ✗
// ➜  2017_02_24 git:(master) ✗ go run  main.go
// map[]
// ➜  2017_02_24 git:(master) ✗
// ➜  2017_02_24 git:(master) ✗ go run  main.go
// map[name:jd]
// ➜  2017_02_24 git:(master) ✗
// ➜  2017_02_24 git:(master) ✗
// ➜  2017_02_24 git:(master) ✗ go run  main.go
// map[name:fish age:18]

// 结果表明
//     都会修改map

// 总结
//     如果要使用map作为参数传递时，不要在函数内修改它
