package main


import (
	"fmt"
	"math/rand" // 按照约定，包名与导入路径的最后一个元素一致
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(100)) // 我们使用rand表示导入的"math/rand"
}
