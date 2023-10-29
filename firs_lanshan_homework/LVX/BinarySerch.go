package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		target       = rand.Intn(100) + 1
		arr    []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
			26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
			51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75,
			76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	)
	fmt.Printf("要找的为整数为%v\n", target)
	a := BinarySerch(target, arr)
	switch BinarySerch(target, arr).(type) {
	case int:
		fmt.Printf("该数在数组中的索引为：%v", a)
	case string:
		fmt.Print(a)
	}
}

func BinarySerch(a int, arr []int) interface{} {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		switch {
		case arr[mid] < a:
			left = mid + 1
		case arr[mid] > a:
			right = mid - 1
		case arr[mid] == a:
			return mid
		}
	}
	return "列表中没有查找到该值"
}
