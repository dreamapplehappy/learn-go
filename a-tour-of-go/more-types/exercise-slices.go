package main

import "fmt"

func main() {
	s := Pic(6, 6)
	fmt.Print(s)
}

func Pic(dx, dy int) [][]uint8 {
	// 这里应该使用二维数组
	s := make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
	}
	return s
}