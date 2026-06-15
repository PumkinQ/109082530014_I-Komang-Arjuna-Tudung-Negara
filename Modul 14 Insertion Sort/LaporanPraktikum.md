# <h1 align="center">Laporan Praktikum Modul 14 - ... </h1>

<p align="center">[I Komang Arjuna Tudung Negara] - [109082530014]</p>

## Unguided

### 1. [Soal]

Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang
diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan
memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.
Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya
bilangan non negatif saja yang disimpan ke dalam array.
Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan
pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam
array. "Data berjarak x" atau "data berjarak tidak tetap".

#### soal1.go

```go
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
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%2014%20Insertion%20Sort/output/Soal1.png)
[penjelasan]
Penulis membuat inputan n yaitu jumlah dearah dan di dalam setiap daerah melakukan perulangan untuk mengambil m yaitu banyaknya rumah, lalu mengisi setiap angka ke dalam slice. Setelah data terkumpul, penulis menerapkan algoritma Selection Sort untuk mengurutkan nomor rumah secara membesar. Selanjutnya, program mengeprint isi slice tersebut secara keseluruhan setelah terurut.

## Unguided

### 2. [Soal]

Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu
perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
const nMax : integer = 7919
type Buku = <
id, judul, penulis, penerbit : string
eksemplar, tahun, rating : integer >
type DaftarBuku = array [ 1..nMax] of Buku
Pustaka : DaftarBuku
nPustaka: integer

Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang
menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya,
masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris
terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.

Halaman 103 | Modul P raktikum Algoritma dan Pemrograman 2
Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua
adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku
yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

#### soal2.go

```go
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
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%2014/output/Soal2.png)
[penjelasan]
Penulis membuat input n untuk jumlah daerah yang harus diproses. lalu setiap perulangan daerah program membaca m dan langsung memilah setiap nomor rumah yang diinput ke dalam dua slice, selanjutnaya dibuat loop, satu untuk bilangan ganjil dan satu untuk bilangan genap, dengan logika temp % 2 != 0.

## Unguided
