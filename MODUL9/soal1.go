package main

import (
	"fmt"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func diDalam(c Lingkaran, t Titik) bool {
	dx := t.x - c.pusat.x
	dy := t.y - c.pusat.y
	return dx*dx+dy*dy <= c.r*c.r
}

func main() {
	var c1, c2 Lingkaran
	var t Titik

	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.r)
	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.r)
	fmt.Scan(&t.x, &t.y)

	in1 := diDalam(c1, t)
	in2 := diDalam(c2, t)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
