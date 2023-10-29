package main

import (
	"fmt"
	"math"
)

func SC(r float64) float64 {
	a := math.Pi * r * r
	return a
}

func main() {
	var a float64
	fmt.Print("请输入圆的半径:")
	fmt.Scan(&a)
	fmt.Printf("圆的面积为：%v", SC(a))
}
