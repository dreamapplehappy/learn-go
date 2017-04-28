package main

import "fmt"

// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型外，其它的都可以省略
func add1(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add1(7, 9))
}
