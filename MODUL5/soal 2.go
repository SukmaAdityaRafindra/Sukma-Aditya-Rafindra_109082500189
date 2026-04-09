package main

import "fmt"

func pola(n, i int) {
	if i > n {
		return
	}
	fmt.Println(string(make([]byte, i, i)))
	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()
	pola(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	pola(n, 1)
}
