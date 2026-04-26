# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>

<p align="center">[I Komang Arjuna Tudung Negara] - [109082530014]</p>

## Unguided

### 1. [Soal]

Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar.
Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak
kelinci yang akan dijual.
Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan
bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya
N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.
Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan
terbesar.

#### soal1.go

```go
package main

import "fmt"

func main() {
	var n int
	var arr [1000]float64

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	if n > 0 {
		min := arr[0]
		max := arr[0]

		for i := 1; i < n; i++ {
			if arr[i] < min {
				min = arr[i]
			}
			if arr[i] > max {
				max = arr[i]
			}
		}
		fmt.Printf("%g %g\n", min, max)
	}
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal1.png)
[penjelasan]
Penulis mendeklarasikan variabel n dan array dengan kapasitas 1000 untuk menampung berat kelinci. Pertama, dibuat inputan n untuk jumlah anak kelinci, lalu dilanjutkan dengan perulangan untuk mengisi nilai ke dalam array. Selanjutny dibuat variabel min dan max yang diinilisisialisasi dengan nilai elemen pertama array. lalu dilakukan perulangan dan percabangan if untuk membandingkan setiap isi array, jika ada yang lebih kecil maka akan masuk ke variabel min dan jika lebih besar akan masuk ke variabel max, lalu hasilnya langsung diprint.

## Unguided

### 2. [Soal]

Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program
ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan
dijual.
Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan
y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya
ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil
yang menyatakan banyaknya ikan yang akan dijual.
Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan
total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang
dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah
sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

#### soal2.go

```go
package main

import "fmt"

func main() {
	var x, y int
	var arr [1000]float64

	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&arr[i])
	}

	var totalBeratSemua float64
	var jumlahWadah int

	for i := 0; i < x; i += y {
		var beratWadah float64
		for j := 0; j < y && i+j < x; j++ {
			beratWadah += arr[i+j]
		}
		fmt.Printf("%g ", beratWadah)
		totalBeratSemua += beratWadah
		jumlahWadah++
	}
	fmt.Println()

	if jumlahWadah > 0 {
		rataRata := totalBeratSemua / float64(jumlahWadah)
		fmt.Printf("%g\n", rataRata)
	}
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal2.png)
[penjelasan]
Penulis membuat inputan x untuk jumlah ikan dan y untuk kapasitas wadah, serta array untuk menyimpan berat masing masing ikan. Selanjutny digunakan nestedloop untuk menjumlahkan berat ikan ke dalam setiap wadah sesuai kapasitas y yang ditentukan. Di dalam perulangan tersebut, selanjutnya total berat tiap wadah diprint secara berurutan, lalu dihitung juga total berat keseluruhan dan jumlah wadahnya. Terakhir, lalu program menghitung rata-rata berat per wadah dengan membagi total berat semua ikan dengan jumlah wadah yang digunakan.

## Unguided

### 3. [Soal]

Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data
berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data
yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
Buatlah program dengan spesifikasi subprogram sebagai berikut:

#### soal3.go

```go
package main

import "fmt"

const NMAX int = 100

type arrBalita [NMAX]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arr arrBalita
	var bMin, bMax float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	if n > 0 {
		hitungMinMax(arr, n, &bMin, &bMax)
		rata := rerata(arr, n)

		fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
		fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
		fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
	}
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal.3.png)
[penjelasan]
Penulis menggunakan tipe bentukan arrBalita untuk berguna menyimpan data berat dalam array. selanjutnya pada program ini dibuat dua subprogram, yaitu func hitungMinMax yang menggunakan parameter pointer untuk mencari nilai terkecil dan terbesar, lalu func rerata untuk menghitung nilai rata rata berat balita. Di bagian main penulis membuat inputan jumlah data dan mengisi array menggunakan perulangan, lalu selanjutnya memanggil fungsi fungsi tadi untuk menampilkan hasil berat minimum, maksimum, dan rata ratanya dengan format dua angka di belakang koma.
