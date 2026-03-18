package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println("permutation: ", permutation(a, c))
		fmt.Println("combination: ", combination(a, c))
		fmt.Println("permutation: ", permutation(b, d))
		fmt.Println("combination: ", combination(b, d))
	}
}

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutation(n, r int) int {
	hasil := factorial(n) / factorial((n - r))
	return hasil
}

func combination(n, r int) int {
	hasil := factorial(n) / (factorial(r) * factorial((n - r)))
	return hasil
}
