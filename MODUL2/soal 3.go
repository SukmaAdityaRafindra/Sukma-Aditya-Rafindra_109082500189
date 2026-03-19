package main

import "fmt"

func main() {
	var gram int
	var Kg, Sisa int
	var biayaKg, biayaSisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)

	Kg = gram / 1000
	Sisa = gram % 1000
	biayaKg = Kg * 10000

	if Kg > 10 {
		biayaSisa = 0
	} else {
		if Sisa >= 500 {
			biayaSisa = Sisa * 5
		} else {
			biayaSisa = Sisa * 15
		}
	}
	total = biayaKg + biayaSisa

	fmt.Println("Detail berat:", Kg, "Kg+", Sisa, "gr")
	fmt.Println("Detail biaya: Rp. ", biayaKg, "+ Rp. ", biayaSisa)
	fmt.Println("Total biaya Rp.", total)
}
