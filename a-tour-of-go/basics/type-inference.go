package main

import "fmt"

func main() {
	// 在声明一个变量而不指定其类型时（即使用不带类型的 := 语法 或 var = 表达式语法），变量的类型由右值推导得出
	var i int
	j := i
	num := 1
	fl := 3.14
	c := 5 - 5i
	const f = "%v is of type %T\n"
	fmt.Printf(f, i, i)
	fmt.Printf(f, j, j)
	fmt.Printf(f, num, num)
	fmt.Printf(f, fl, fl)
	fmt.Printf(f, c, c)
}
