package main

import "fmt"

func main() {
	fruits := [4]string {
		"apple",
		"orange",
		"banana",
		"pear",
	}

	a := fruits[0:2]
	b := fruits[1:3]

	fmt.Println(fruits)
	fmt.Println(a)
	fmt.Println(b)

	a[0] = "apple1"
	fmt.Println(fruits)
	fmt.Println(a)

	a[1] = "orange1"
	fmt.Println(fruits)
	fmt.Println(a)
	fmt.Println(b)

}
