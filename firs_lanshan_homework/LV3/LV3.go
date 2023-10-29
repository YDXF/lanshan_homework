package main

import (
	"fmt"
)

func prime_number(a int) {
	var (
		b int
		c string
	)
	for i := 2; i <= a; i++ {
		if a%i == 0 {
			b++
		}
	}
	switch {
	case b == 1 && a != 1:
		c = "是素数"
	default:
		c = "不是素数"
	}
	fmt.Printf("%v%v", a, c)
}

func main() {
	var a int
	for {
		fmt.Print("请输入你想判断的数（输入0退出）：")
		fmt.Scan(&a)
		if a == 0 {
			return
		}
		prime_number(a)
	}
}
