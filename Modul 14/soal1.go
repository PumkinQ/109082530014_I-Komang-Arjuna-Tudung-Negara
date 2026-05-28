package main

import "fmt"

func rumahkerabat(n int) {
	var m, idx_min int
	for i := 0; i < n; i++ {
		var arr []int
		fmt.Scan(&m)
		arr = append(arr, m)
		for j := 0; j < m; j++ {
			fmt.Scan(arr[&j])
			for x := 0; x < m-1; x++ {
				idx_min := x
				for y := x + 1; y < m; y++ {
					if arr[y] < arr[idx_min] {
						idx_min = y
					}
				}
				// Swap (Tukar posisi)
				arr[x], arr[idx_min] = arr[idx_min], arr[x]
			}
			for k := 0; k < m; k++ {
				fmt.Print(arr[k], " ")
			}
			fmt.Println() // Pindah baris untuk daerah berikutnya
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	rumahkerabat(n)
}
