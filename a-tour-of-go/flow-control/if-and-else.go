package main

import (
	"math"
	"fmt"
)

func pow1(a, b, lim float64) float64 {
	// 在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用
	if v := math.Pow(a, b); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 注意不要忽略掉返回值，这里就不可以使用v变量了
	return lim
}

func main() {
	fmt.Println(pow1(3, 2, 10))
	fmt.Println(pow1(3, 3, 10))
}
