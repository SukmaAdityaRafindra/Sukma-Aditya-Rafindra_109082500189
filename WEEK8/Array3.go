package main

import "fmt"

func main() {
	var Kiri = [6]int{6, 7, 3, 5, 8, 2}
	var Kanan = [6]int{2, 8, 5, 3, 7, 6}

	for i := 0; i < 6; i++ {
		if Kiri[i] < Kanan[i] {
			fmt.Print(Kiri)
		} else {
			fmt.Print(Kanan)
		}
	}
}
