package main

import (
	"math"
	"fmt"
)

// 注意这个函数的返回值是string 因为可能有负数出现
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// 使用了 fmt.Sprint
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
