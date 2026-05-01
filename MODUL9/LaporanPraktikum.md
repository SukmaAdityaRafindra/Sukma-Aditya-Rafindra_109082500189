# <h1 align="center">Laporan Praktikum Modul 9 - ... </h1>

<p align="center">[Sukma Aditya Rafindra] - [109082500189]</p>

## Unguided 

### 1. [Soal 1]
#### soal 1.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL9/Output/image.png)
Program ini adalah alat pengecek posisi yang menentukan apakah sebuah titik berada di dalam satu, kedua, atau di luar dua lingkaran yang ditentukan. Dengan menggunakan rumus Pythagoras untuk menghitung jarak antara titik dan pusat lingkaran, program membandingkan hasil tersebut dengan panjang jari-jari untuk memberikan jawaban instan mengenai status keberadaan titik tersebut terhadap kedua area lingkaran.
### 2. [Soal 2]
#### soal 2.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Semua:", arr)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Scan(&x)
	fmt.Print("Kelipatan indeks ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
	fmt.Scan(&idx)
	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println("Setelah hapus:", arr)

	sum := 0
	for _, v := range arr {
		sum += v
	}
	avg := float64(sum) / float64(len(arr))
	fmt.Println("Rata-rata:", avg)

	var total float64
	for _, v := range arr {
		total += math.Pow(float64(v)-avg, 2)
	}
	std := math.Sqrt(total / float64(len(arr)))
	fmt.Println("Std dev:", std)

	var cari, count int
	fmt.Scan(&cari)
	for _, v := range arr {
		if v == cari {
			count++
		}
	}
	fmt.Println("Frekuensi:", count)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL9/Output/image copy.png)
Program ini adalah pengolah kumpulan angka yang mampu memfilter data berdasarkan indeks (ganjil, genap, atau kelipatan tertentu), menghapus elemen di posisi tertentu, serta melakukan analisis statistik seperti menghitung rata-rata dan standar deviasi. Selain itu, program ini juga dapat mencari frekuensi kemunculan suatu angka spesifik di dalam daftar tersebut.


### 3. [Soal 3]
#### soal 3.go

```go
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
	}

	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL9/Output/image copy 2.png)
Program ini mencatat pemenang setiap pertandingan antara dua klub secara berulang hingga pengguna memasukkan skor negatif, lalu menampilkan ringkasan seluruh hasil akhirnya.


### 4. [Soal 4]
#### soal 4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	fmt.Print("Teks : ")
	for *n < NMAX {
		fmt.Scanf("%c", &input)
		if input == '.' || input == '\n' {
			break
		}
		if input != ' ' {
			t[*n] = input
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	asli := t
	balikanArray(&t, n)
	for i := 0; i < n; i++ {
		if asli[i] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	isPal := palindrom(tab, m)

	fmt.Print("Reverse teks : ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	fmt.Printf("Palindrom : %t\n", isPal)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL9/Output/image copy 3.png)
Program ini memproses input teks untuk membalikkan urutan karakternya dan mengecek apakah teks tersebut adalah palindrom (kata yang terbaca sama dari depan maupun belakang) setelah mengabaikan spasi.

