package main

import "fmt"

// 函数的定义，其中函数的参数需要指定类型；函数的返回值也需要指定类型
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(18, 36))
}