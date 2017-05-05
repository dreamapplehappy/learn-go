package main

import "fmt"

func main() {
	s := make([]int, 5)
	fmt.Println(s, len(s), cap(s))
	s = make([]int, 0, 5)
	fmt.Println(s, len(s), cap(s))
	s = s[:cap(s)]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
}
