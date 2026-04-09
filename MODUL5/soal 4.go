package main

import "fmt"

func turunNaik(n int) {
	if n == 0 {
		return
	}

	fmt.Print(n, " ")
	turunNaik(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	turunNaik(n)
}
