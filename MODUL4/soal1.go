package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	p1 := faktorial(a) / faktorial(a-c)
	c1 := faktorial(a) / (faktorial(c) * faktorial(a-c))

	p2 := faktorial(b) / faktorial(b-d)
	c2 := faktorial(b) / (faktorial(d) * faktorial(b-d))

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
