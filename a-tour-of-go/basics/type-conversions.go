package main

import (
	"math"
	"fmt"
)

func main() {
	// Go在不同类型的项之间赋值时需要显式转换
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
