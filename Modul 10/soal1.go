package main

import "fmt"

func main() {
	var n int
	var arr [1000]float64

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	if n > 0 {
		min := arr[0]
		max := arr[0]

		for i := 1; i < n; i++ {
			if arr[i] < min {
				min = arr[i]
			}
			if arr[i] > max {
				max = arr[i]
			}
		}
		fmt.Printf("%g %g\n", min, max)
	}
}
