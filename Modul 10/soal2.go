package main

import "fmt"

func main() {
	var x, y int
	var arr [1000]float64

	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&arr[i])
	}

	var totalBeratSemua float64
	var jumlahWadah int

	for i := 0; i < x; i += y {
		var beratWadah float64
		for j := 0; j < y && i+j < x; j++ {
			beratWadah += arr[i+j]
		}
		fmt.Printf("%g ", beratWadah)
		totalBeratSemua += beratWadah
		jumlahWadah++
	}
	fmt.Println()

	if jumlahWadah > 0 {
		rataRata := totalBeratSemua / float64(jumlahWadah)
		fmt.Printf("%g\n", rataRata)
	}
}
