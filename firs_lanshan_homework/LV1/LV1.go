package main

import (
	"fmt"
)

func add(a int, b int) int {
	c := a + b
	return c
}

func main() {
	var a, b int
	fmt.Print("请输入相加的两个整数：")
	fmt.Scan(&a, &b)
	fmt.Printf("它们的和为：%v", add(a, b))
}
