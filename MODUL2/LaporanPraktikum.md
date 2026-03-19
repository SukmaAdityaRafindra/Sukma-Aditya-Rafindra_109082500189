# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>

<p align="center">[Sukma Aditya Rafindra] - [109082500189]</p>

## Unguided 

### 1. [Soal 1]
#### soal 1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL2/Output/soal1.png)
berfungsi untuk menerima tiga input string dari pengguna, menampilkan urutan awalnya, kemudian menukar posisi ketiga string tersebut menggunakan variabel sementara (temp), lalu menampilkan hasil urutan akhirnya.

### 2. [Soal 2]
#### soal 2.go

```go
package main

import "fmt"

func main() {
	var w1, w2, w3, w4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&w1, &w2, &w3, &w4)

		if !(w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "unggu") {
			berhasil = false
		}
	}
	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL2/Output/soal2.png)
berfungsi untuk melakukan 5 kali percobaan input empat kata, kemudian memeriksa apakah urutan kata yang dimasukkan adalah “merah kuning hijau ungu”, dan jika ada yang tidak sesuai maka nilai berhasil akan menjadi false, lalu hasil akhirnya ditampilkan.


### 3. [Soal 3]
#### soal 3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL2/Output/soal3.png)
berfungsi untuk menghitung total biaya pengiriman berdasarkan berat parsel yang dimasukkan dalam gram, dengan mengubahnya ke kilogram dan sisa gram, kemudian menentukan biaya tambahan dari sisa berat sebelum menampilkan detail berat dan total biaya pengiriman.

