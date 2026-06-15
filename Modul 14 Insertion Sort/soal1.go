package main

import "fmt"

func main() {
	var input, temp, i, j int
	var angka []int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		angka = append(angka, input)
	}

	n := len(angka)

	for i = 1; i <= n-1; i++ {
		j = i
		temp = angka[j]

		for j > 0 && temp < angka[j-1] {
			angka[j] = angka[j-1]
			j = j - 1
		}

		angka[j] = temp
	}

	for i = 0; i < n; i++ {
		fmt.Print(angka[i], " ")
	}
	fmt.Println()

	if n < 2 {
		fmt.Println("Data berjarak 0")
		return
	}

	jarak := angka[1] - angka[0]
	tetap := true

	for i = 1; i < n-1; i++ {
		if angka[i+1]-angka[i] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
