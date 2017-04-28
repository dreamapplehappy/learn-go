package main

import "fmt"

// 如果返回语句没有返回值，它会默认返回函数命名的变量
func split(num int) (x, y int) {
	x = num * 4 / 9
	y = num - x
	return
}

func main() {
	fmt.Println(split(17))
}
