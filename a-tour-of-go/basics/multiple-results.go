package main

import "fmt"

// 函数可以返回多个值
func swap(a string, b string) (string, string) {
	return b, a
}

func main() {
	// 字符串使用双引号
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
