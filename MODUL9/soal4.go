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
