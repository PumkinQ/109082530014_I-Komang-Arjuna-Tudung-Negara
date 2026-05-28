package main

import "fmt"

func main() {
	var n, m, temp int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			fmt.Scan(&temp)
			if temp%2 != 0 {
				ganjil = append(ganjil, temp)
			} else {
				genap = append(genap, temp)
			}
		}

		for x := 0; x < len(ganjil)-1; x++ {
			idx := x
			for y := x + 1; y < len(ganjil); y++ {
				if ganjil[y] < ganjil[idx] {
					idx = y
				}
			}
			ganjil[x], ganjil[idx] = ganjil[idx], ganjil[x]
		}

		for x := 0; x < len(genap)-1; x++ {
			idx := x
			for y := x + 1; y < len(genap); y++ {
				if genap[y] > genap[idx] {
					idx = y
				}
			}
			genap[x], genap[idx] = genap[idx], genap[x]
		}

		for _, val := range ganjil {
			fmt.Print(val, " ")
		}
		for _, val := range genap {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
