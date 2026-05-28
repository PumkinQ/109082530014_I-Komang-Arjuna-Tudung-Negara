package main

import "fmt"

func main() {

	var n, m, idx_min, temp int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&m)
		var arr []int
		for j := 0; j < m; j++ {

			fmt.Scan(&temp)
			arr = append(arr, temp)

		}
		for x := 0; x < m-1; x++ {
			idx_min = x
			for y := x + 1; y < m; y++ {
				if arr[y] < arr[idx_min] {
					idx_min = y
				}
			}
			arr[x], arr[idx_min] = arr[idx_min], arr[x]
		}
		fmt.Println("HASIL Loop ke-", i+1)
		for k := 0; k < m; k++ {
			fmt.Print(arr[k], " ")
		}
		fmt.Println()
		fmt.Println()
	}

}
