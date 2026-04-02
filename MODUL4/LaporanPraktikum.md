# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>

<p align="center">[Sukma Aditya Rafindra] - [109082500189]</p>

## Unguided 

### 1. [Soal 1]
#### soal 1.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL4/Output/soal1_modul4.png)
Program ini membaca beberapa input angka lalu menghitung nilai faktorialnya menggunakan fungsi rekursif dan menampilkan hasilnya di terminal.


### 2. [Soal 2]
#### soal 2.go

```go
package main

import "fmt"

func hitungSkor(soal, waktu_m *int) {
	*waktu_m = 0
	*soal = 0
	var waktu int
	for i := 1; i <= 8; i++ {
		fmt.Scan(&waktu)

		if waktu < 301 {
			*soal++
			*waktu_m += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, waktu int
	cek := true
	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		var banyaksoal, waktu_terdikit int
		hitungSkor(&banyaksoal, &waktu_terdikit)

		if cek {
			soal = banyaksoal
			waktu = waktu_terdikit
			pemenang = nama
			cek = false
		} else if banyaksoal > soal || (soal == banyaksoal && waktu_terdikit < waktu) {
			soal = banyaksoal
			waktu = waktu_terdikit
			pemenang = nama
		}
	}

	fmt.Println(pemenang, soal, waktu)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL4/Output/soal2_modul4.png)
Program ini membaca data waktu pengerjaan soal sebanyak 8 kali, lalu menghitung berapa soal yang selesai dalam waktu kurang dari 301 detik sekaligus menjumlahkan total waktunya, kemudian menampilkan hasil akhirnya.



### 3. [Soal 3]
#### soal 3.go

```go
package main

import "fmt"

func deret(n int) {
	fmt.Printf("%d ", n)
	for n != 1 && n > 0 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Printf("%d ", n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	deret(n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL4/Output/soal3_modul4.png)
Program ini menampilkan urutan bilangan dari suatu angka dengan aturan: jika genap dibagi 2, jika ganjil dikali 3 lalu ditambah 1, hingga mencapai 1.


