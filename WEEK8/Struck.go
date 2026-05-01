package main

import "fmt"

func main() {

	type BukuFavorit struct {
		Judul   string
		ISBN    int
		Author  string
		Halaman int
	}

	var Buku1 = BukuFavorit{
		Judul:   "One Piece",
		ISBN:    9780763276535,
		Author:  "Mardi",
		Halaman: 15,
	}

	var Buku2 = BukuFavorit{
		Judul:   "Naruto",
		ISBN:    9720463776538,
		Author:  "Acim",
		Halaman: 20,
	}

	var Buku3 = BukuFavorit{
		Judul:   "Boruto",
		ISBN:    9120463786535,
		Author:  "Rizky",
		Halaman: 25,
	}

	fmt.Println("Judul Buku 1:", Buku1.Judul)
	fmt.Println("ISBN Buku 1:", Buku1.ISBN)
	fmt.Println("Author Buku 1:", Buku1.Author)
	fmt.Println("Halaman Buku 1:", Buku1.Halaman)

	fmt.Println("Judul Buku 2:", Buku2.Judul)
	fmt.Println("ISBN Buku 2:", Buku2.ISBN)
	fmt.Println("Author Buku 2:", Buku2.Author)
	fmt.Println("Halaman Buku 2:", Buku2.Halaman)

	fmt.Println("Judul Buku 3:", Buku3.Judul)
	fmt.Println("ISBN Buku 3:", Buku3.ISBN)
	fmt.Println("Author Buku 3:", Buku3.Author)
	fmt.Println("Halaman Buku 3:", Buku3.Halaman)
}
