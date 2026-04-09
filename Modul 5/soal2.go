package main

import "fmt"

func cetakbintang(f int) {
	if f == 0 {
		return
	}
	fmt.Print("*")
	cetakbintang(f - 1)
}

func bintang(n, f int) {
	if f >= n {
		return
	}
	cetakbintang(f)
	fmt.Println("*")
	bintang(n, f+1)

}

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n, 0)
}
