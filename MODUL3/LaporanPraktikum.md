# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>

<p align="center">[Sukma Aditya Rafindra] - [109082500189]</p>

## Unguided 

### 1. [Soal 1]
#### soal 1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL3/Output/SOAL1_MODUL3.png)
Program tersebut digunakan untuk membaca empat bilangan (a, b, c, d), lalu menghitung dan menampilkan hasil permutasi serta kombinasi dari a terhadap c dan b terhadap d dengan memanfaatkan fungsi faktorial.

### 2. [Soal 2]
#### soal 2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	hasil1 := f(g(h(a)))

	hasil2 := g(h(f(b)))

	hasil3 := h(f(g(c)))

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL3/Output/SOAL2_MODUL3.png)
Program tersebut membaca tiga bilangan lalu menghitung dan menampilkan hasil komposisi fungsi f(g(h(a))), g(h(f(b))), dan h(f(g(c))) menggunakan fungsi matematika yang telah didefinisikan.


### 3. [Soal 3]
#### soal 3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow((a-c), 2) + math.Pow((b-d), 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var (
		cx1, cy1, cx2, cy2, r1, r2 float64
		x, y                       float64
	)
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	d1 := didalam(cx1, cy1, r1, x, y)
	d2 := didalam(cx2, cy2, r2, x, y)

	if d1 && d2 {
		fmt.Print("titik didalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Print("titik didalam lingkaran 1")
	} else if d2 {
		fmt.Print("titik didalam lingkaran 2")
	} else {
		fmt.Print("titik diluar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL3/Output/SOAL2_MODUL3.png)
Program ini digunakan untuk menentukan posisi sebuah titik terhadap dua lingkaran dengan menghitung jaraknya dari pusat lingkaran dan membandingkannya dengan radius.

