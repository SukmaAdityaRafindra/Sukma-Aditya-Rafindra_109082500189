package main

import "fmt"

func main() {
	arr := [6]int{6, 7, 3, 5, 8, 2}
	fmt.Println("Array")
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
