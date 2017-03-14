package main

// 必须是可导出方法
// 否则会报错：panic: plugin: symbol add not found in plugin jdscript.com/day_test/2017_03_03/plugin
func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}
