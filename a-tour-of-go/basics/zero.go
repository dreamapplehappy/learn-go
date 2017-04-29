package main

import "fmt"

func main() {
	// 没有明确初始值的变量声明会被赋予一个值，具体的数值可以看下面的代码
	var (
		i int
		str string
		b bool
		f float32
		p complex128
	)
	fmt.Printf("%v, %q, %v, %v, %v\n", i, str, b, f, p)
}
