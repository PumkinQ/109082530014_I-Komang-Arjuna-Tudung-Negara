package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, idxHapus int
	var arr [100]int
	var total float64

	fmt.Println("input isi Array")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Scan(&idxHapus)
	for i := idxHapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		total += float64(arr[i])
	}
	rataRata := total / float64(n)
	fmt.Printf("%.2f\n", rataRata)

	var jumlahKuadrat float64
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rataRata
		jumlahKuadrat += selisih * selisih
	}
	stdDev := math.Sqrt(jumlahKuadrat / float64(n))
	fmt.Printf("%.2f\n", stdDev)
}
