package main

import (
	"math/cmplx"
	"fmt"
)
// 数据类型 int8, uint8, float32, complex64等
var (
	ToBe bool = true
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
