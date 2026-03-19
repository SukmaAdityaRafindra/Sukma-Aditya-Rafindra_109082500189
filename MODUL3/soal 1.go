package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := n; i >= 1; i-- {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n int, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n int, r int) int {
	return faktorial(n) / faktorial(r) * faktorial(n-r)
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	p1 := permutasi(a, c)
	k1 := kombinasi(a, c)

	p2 := permutasi(b, d)
	k2 := kombinasi(b, d)

	fmt.Println(p1, k1)
	fmt.Println(p2, k2)
}
