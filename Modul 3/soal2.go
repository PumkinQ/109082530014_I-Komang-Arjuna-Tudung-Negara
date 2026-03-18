package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Printf("(fogog)(%d) = %d ", a, fx(a))
	fmt.Printf("(gohof)(%d) = %d ", b, gx(b))
	fmt.Printf("(hofog)(%d) = %d ", c, hx(c))
}

func fx(n int) int {
	hasil := ((n + 1) - 2) * ((n + 1) - 2)
	return hasil
}

func gx(n int) int {
	hasil := (((n * n) + 1) - 2)
	return hasil
}

func hx(n int) int {
	hasil := ((n - 2) * (n - 2)) + 1
	return hasil
}
