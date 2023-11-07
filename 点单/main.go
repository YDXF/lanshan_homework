package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		input  string
		price  int
		a      = 5
		b      = 8
		Coffee = "Coffee"
		Coke   = "Coke"
	)
	const (
		c = 10
	)

	fmt.Printf("\n%-15v $ %2v\n", "Milk", a)
	fmt.Printf("%-15v $ %2v\n", "Coffee", b)
	fmt.Printf("%-15v $ %2v\n\n", "Coke", c)

	fmt.Print("What do you want?\n")
	fmt.Scanln(&input)

	if strings.Contains(input, Coffee) {
		price = b
	} else if strings.Contains(input, Coke) {
		price = c
	} else {
		price = a
	}
	fmt.Printf("You have to pay $%v.\n\n", price)
	fmt.Println("程序已执行完毕。按 Enter 键退出...")
	fmt.Scanln()
}
