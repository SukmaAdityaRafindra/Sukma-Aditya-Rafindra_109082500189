package main

import "fmt"

func main() {
	arr := [6]int{2, 8, 5, 3, 7, 6}
	fmt.Println("Array")
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
