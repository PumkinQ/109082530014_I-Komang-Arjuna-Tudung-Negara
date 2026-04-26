package main

import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat titik
	r     int
}

func cekjarak(p1, p2 titik) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func cekdalam(c lingkaran, p titik) bool {
	return cekjarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var arr [2]lingkaran
	var p titik

	for i := 0; i < 2; i++ {
		fmt.Scan(&arr[i].pusat.x, &arr[i].pusat.y, &arr[i].r)
	}

	fmt.Scan(&p.x, &p.y)

	cek1 := cekdalam(arr[0], p)
	cek2 := cekdalam(arr[1], p)

	if cek1 && cek2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if cek1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if cek2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
