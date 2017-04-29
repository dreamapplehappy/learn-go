package main

import (
	"math"
	"fmt"
)

func pow(a, b, lim float64) float64 {
	// if 语句也可以在条件表达式执行之前执行一个简单的语句
	if v := math.Pow(a, b); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))
}
