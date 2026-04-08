package main

import "fmt"

func main() {
	var a, b, c, d int
	var hasilc1, hasilc2, hasilp1, hasilp2 int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {

		permutation(a, c, &hasilp1)
		combination(a, c, &hasilc1)
		permutation(b, d, &hasilp2)
		combination(b, d, &hasilc2)
		fmt.Println(hasilp1, hasilc1)
		fmt.Print(hasilp2, hasilc2)
	}

}

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}

func permutation(n, r int, hasil *int) {
	var hasilf, hasilnr int
	factorial(n, &hasilf)
	factorial(n-r, &hasilnr)

	*hasil = hasilf / hasilnr

}

func combination(n int, r int, hasil *int) {
	var hasilf, hasilnr, inputr int
	factorial(n, &hasilf)
	factorial(n-r, &hasilnr)
	factorial(r, &inputr)
	*hasil = hasilf / (inputr * hasilnr)

}
