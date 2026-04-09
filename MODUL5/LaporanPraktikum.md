# <h1 align="center">Laporan Praktikum Modul 5 - ... </h1>

<p align="center">[Sukma Aditya Rafindra] - [109082500189]</p>

## Unguided 

### 1. [Soal 1]
#### soal 1.go

```go
package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fib(i), " ")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL5/Output/SOAL1_MODUL5.png)
Program ini menghitung deret Fibonacci, di mana setiap angka didapat dari penjumlahan dua angka sebelumnya.



### 2. [Soal 2]
#### soal 2.go

```go
package main

import "fmt"

func pola(n, i int) {
	if i > n {
		return
	}
	fmt.Println(string(make([]byte, i, i)))
	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()
	pola(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	pola(n, 1)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL5/Output/SOAL2_MODUL5.png)
Program ini menampilkan bintang dari 1 sampai N (bentuk segitiga).




### 3. [Soal 3]
#### soal 3.go

```go
package main

import "fmt"

func faktor(n, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	faktor(n, 1)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL5/Output/SOAL3_MODUL5.png)
Program ini tujuannya untuk menampilkan faktor dari suatu bilangan N (angka yang bisa membagi N habis).


### 4. [Soal 4]
#### soal 4.go

```go
ppackage main

import "fmt"

func turunNaik(n int) {
	if n == 0 {
		return
	}

	fmt.Print(n, " ")
	turunNaik(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	turunNaik(n)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL5/Output/SOAL4_MODUL5.png)
Program ini menampilkan angka dari N ke 1, lalu balik lagi ke N.


### 5. [Soal 5]
#### soal 5.go

```go
package main

import "fmt"

func ganjil(i, n int) {
	if i > n {
		return
	}

	if i%2 != 0 {
		fmt.Print(i, " ")
	}

	ganjil(i+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	ganjil(1, n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL5/Output/SOAL5_MODUL5png)
Program ini menampilkan angka ganjil aja dari 1 sampai N.



### 6. [Soal 6]
#### soal 6.go

```go
package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	hasil := pangkat(x, y)
	fmt.Println("Hasil:", hasil)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/SukmaAdityaRafindra/Sukma-Aditya-Rafindra_109082500189/blob/main/MODUL5/Output/SOAL6_MODUL5.png)
Program ini menghitung pangkat dari dua bilangan (x^y) pakai rekursif.

