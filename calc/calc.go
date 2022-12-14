package calc

import "fmt"

func init() {
	fmt.Println("导入calc包时自动调用")
}

func Add(a, b int) (res int) {
	res = a + b
	return
}

func Sub(a, b int) int {
	return a - b
}
