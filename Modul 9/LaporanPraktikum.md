# <h1 align="center">Laporan Praktikum Modul 9 - ... </h1>

<p align="center">[I Komang Arjuna Tudung Negara] - [109082530014]</p>

## Unguided

### 1. [Soal]

Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila
diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan
koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan
radiusnya.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat titik
	r     int
}

func cekjarak(p1, p2 titik) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func cekdalam(c lingkaran, p titik) bool {
	return cekjarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var arr [2]lingkaran
	var p titik

	for i := 0; i < 2; i++ {
		fmt.Scan(&arr[i].pusat.x, &arr[i].pusat.y, &arr[i].r)
	}

	fmt.Scan(&p.x, &p.y)

	cek1 := cekdalam(arr[0], p)
	cek2 := cekdalam(arr[1], p)

	if cek1 && cek2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if cek1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if cek2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}


```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%209/output/Soal1.png)
[penjelasan]
Penulis membuat tipe bentukan titik dan lingkaran, lalu membuat func cekjarak untuk hitung jarak koordinat pakai rumus Pythagoras. Selanjutnya, ada func cekdalam buat cek apakah jaraknya lebih kecil dari radius. Di bagian main, dibuat inputan untuk dua lingkaran dan satu titik, lalu ada percabangan if-else untuk menentukan posisi titik tersebut ada di dalam lingkaran 1, lingkaran 2, atau keduanya.

## Unguided

### 2. [Soal]

Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program
yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array
memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat
menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.
b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah
genap).
d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh
dari masukan pengguna.
e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.
Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array
tersebut.

#### soal2.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, idxHapus int
	var arr [100]int
	var total float64

	fmt.Println("input isi Array")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Scan(&idxHapus)
	for i := idxHapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		total += float64(arr[i])
	}
	rataRata := total / float64(n)
	fmt.Printf("%.2f\n", rataRata)

	var jumlahKuadrat float64
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rataRata
		jumlahKuadrat += selisih * selisih
	}
	stdDev := math.Sqrt(jumlahKuadrat / float64(n))
	fmt.Printf("%.2f\n", stdDev)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%209/output/Soal2.png)
[penjelasan]
Penulis membuat inputan n dan isi array, lalu selanjutnya mengeprint isi array secara keseluruhan, indeks ganjil, dan indeks genap. Terakhir, ada inputan x untuk cek indeks kelipatan x dan inputan untuk menghapus elemen array dengan cara menggeser indeksnya. Setelah itu, program menghitung ratarata dan standar deviasi pake fungsi akar dari library math.

## Unguided

### 3. [Soal]

Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang
memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang
digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.
Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian
program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan
dalam array adalah nama-nama klub yang menang saja.
Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir
program, tampilkan daftar klub yang memenangkan pertandingan.

#### soal3.go

```go
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string
	var i int = 1

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		i++
	}

	for j := 0; j < len(pemenang); j++ {
		fmt.Printf("Hasil %d: %s\n", j+1, pemenang[j])
	}
	fmt.Println("Pertandingan selesai")
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%209/output/Soal3.png)
[penjelasan]
Penulis membuat inputan nama dua klub, lalu masuk ke perulangan untuk input skor terusmenerus. Selanjutnya, ada percabangan kalo skornya negatif maka perulangan bakal berhenti. Di dalamnya juga dicek siapa yang menang untuk dimasukkan ke slice pemenang, lalu terakhir program bakal mengeprint semua hasil pertandingan yang sudah tersimpan tadi.

## Unguided

### 4. [Soal]

Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk
membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa
apakah membentuk palindrom.

#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0
	for *n < NMAX {
		fmt.Scanf("%c", &input)
		if input == '.' || input == '\n' {
			break
		}
		if input != ' ' {
			t[*n] = input
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var tAsli tabel = t
	balikanArray(&t, n)
	for i := 0; i < n; i++ {
		if tAsli[i] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks: ")
	isiArray(&tab, &m)

	isPalindrom := palindrom(tab, m)

	fmt.Print("Reverse teks: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	fmt.Printf("Palindrom? %t\n", isPalindrom)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%209/output/Soal4.png)
[penjelasan]
Penulis mendeklarasi tipe bentukan tabel untuk menyimpan array karakter, lalu dibuat func isiArray untuk mengambil inputan sampai tanda titik Selanjutnya, ada func balikanArray yang untuk memutar urutan isi array, dan func palindrom untuk mengecek apakah kata aslinya sama dengan kata yang sudah dibalik. Di bagian main, program menjalankan fungsi tersebut lalu menampilkan hasil kebalikan teksnya dan status apakah kata tersebut termasuk palindrom atau bukan.
