package main

import (
	"fmt"
)

type calculator func(a, b float64) float64

func add(a, b float64) float64 {
	return a + b
}

func substract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("除数不可为零")
		return 0.0
	} else {
		return a / b
	}
}

func calculate(a, b float64, operator calculator) float64 {
	return operator(a, b)
}

func main() {
	var (
		fn, sn, result     float64
		sign_of_operation  string
		selectedcalculator calculator
	)
	fmt.Print("请输入第一个数：")
	fmt.Scanln(&fn)
	fmt.Print("请输入运算符号：")
	fmt.Scanln(&sign_of_operation)
	fmt.Print("请输入第二个数：")
	fmt.Scanln(&sn)

	switch sign_of_operation {
	case "+":
		selectedcalculator = add
	case "-":
		selectedcalculator = substract
	case "/":
		selectedcalculator = divide
	case "*":
		selectedcalculator = multiply
	default:
		fmt.Print("无效输入，请重来")
		return
	}

	result = calculate(fn, sn, selectedcalculator)
	fmt.Printf("%8v答案是：%v", "", result)
}
