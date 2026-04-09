package main

import "fmt"

func mod(n, b int) {

	if b > 0 {
		mod(n, b-1)
		if n%b == 0 {
			fmt.Print(b, " ")
		}

	}
}

func main() {
	var n int
	fmt.Scan(&n)
	mod(n, n)
}
