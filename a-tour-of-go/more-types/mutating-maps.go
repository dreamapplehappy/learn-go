package main

import "fmt"

func main() {
	var m map[string]int
	// 插入或者修改元素
	m = make(map[string]int)
	m["answer"] = 12
	fmt.Println(m["answer"])
	// 获取元素
	answer := m["answer"]
	fmt.Println(answer)
	// 删除元素
	delete(m, "answer")
	fmt.Println(m)
	// 双赋值
	var result int = 1
	var ok bool = true
	result, ok = m["answer"]
	fmt.Println(result, ok)
}
