package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) < r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	L1 := didalam(cx1, cy1, r1, x, y)
	L2 := didalam(cx2, cy2, r2, x, y)

	if L1 && L2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if L1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if L2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
