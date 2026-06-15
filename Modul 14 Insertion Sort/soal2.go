package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}
	idxMax := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}
	fmt.Printf("%s %s %s %d\n", pustaka[idxMax].judul, pustaka[idxMax].penulis, pustaka[idxMax].penerbit, pustaka[idxMax].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var j int
	for i := 1; i < n; i++ {
		j = i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}
		pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low := 0
	high := n - 1
	found := -1

	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			found = mid
			break
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if found != -1 {
		b := pustaka[found]
		fmt.Printf("%s %s %s %d %d %d\n", b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var Pustaka DaftarBuku
	var n, ratingCari int

	fmt.Scan(&n)
	DaftarkanBuku(&Pustaka, n)
	fmt.Scan(&ratingCari)

	CetakTerfavorit(Pustaka, n)
	UrutBuku(&Pustaka, n)
	Cetak5Terbaru(Pustaka, n)
	CariBuku(Pustaka, n, ratingCari)
}
