package main

import (
	"fmt"
	"math"
)

func SqrtHelper(x float64, z float64) float64 {
	var next = z - ((z * z) -x) / (2 * z)
	fmt.Println(next)
	return next
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z = SqrtHelper(x, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(100), math.Sqrt(100))
}
