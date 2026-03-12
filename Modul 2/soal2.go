package main

import "fmt"

//Program Mengitung hasil dan peserta donasi
func main() {
	var w1, w2, w3, w4, acuan1, acuan2 string
	var hasil bool
	i := 1
	hasil = true
	fmt.Printf("Percobaan %d:", i)
	fmt.Scan(&w1, &w2, &w3, &w4)
	acuan1 = w1 + w2 + w3 + w4
	for i = 2; i <= 5; i++ {
		fmt.Printf("Percobaan %d:", i)
		fmt.Scan(&w1, &w2, &w3, &w4)

		acuan2 = w1 + w2 + w3 + w4
		if acuan2 != acuan1 {
			hasil = false
		}
	}
	fmt.Print("Hasil :")
	fmt.Print(hasil)
}
