package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := a[:]
	fmt.Println(s)
	s = a[0:]
	fmt.Println(s)
	s = a[:10]
	fmt.Println(s)
	s = a[0:10]
	fmt.Println(s)
}
