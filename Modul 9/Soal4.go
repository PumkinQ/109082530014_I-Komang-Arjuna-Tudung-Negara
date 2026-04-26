package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0
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
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var tAsli tabel = t
	balikanArray(&t, n)
	for i := 0; i < n; i++ {
		if tAsli[i] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks: ")
	isiArray(&tab, &m)

	isPalindrom := palindrom(tab, m)

	fmt.Print("Reverse teks: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	fmt.Printf("Palindrom? %t\n", isPalindrom)
}
