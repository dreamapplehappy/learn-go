package main

import "fmt"

// *T表示指向T类型的指针，其零值是nil
// &操作符会生成一个指向其操作数的指针
// *操作符表示指针指向的底层值

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
