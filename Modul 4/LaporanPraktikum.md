# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>

<p align="center">[I Komang Arjuna Tudung Negara] - [109082530014]</p>

## Unguided

### 1. [Soal]

Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan
menggunakan persamaan berikut!
P(n, r) =
n!
(n−r)!
, sedangkan C(n, r) =
n!
r!(n−r)!

#### soal1.go

```go
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
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal1.png)
[penjelasan]
Penulis mendeklarasikan variable a, b, c, d untuk inputan, lalu hasilc1, hasilc2, hasilp1, hasilp2 untuk isian jawaban dari soal permutasi dan combinasi, lalu dibuat if else untuk syarat combinasi dan premutasi, lalu dibuat fmt print untuk hasil hasil dari permutasi dan combinasi, lalu dibuat func factorial dengan variable sbg inputan dan hasil sebagai output, permutasi dengan n dan r sebagai inputan lalu hasil sebagai output dan cobinmsi dengan n dan r sebagai inputan lalu hasil sebagai output.

## Unguided

### 2. [Soal]

Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal
yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan
soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program
harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total
soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal.
Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca
di dalam prosedur.
prosedure hitungSkor(in/out soal, skor : integer)
Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah
8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal.
Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan
dalam waktu 5 jam 1 menit (301 menit).
Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang
diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil
diselesaikan.

#### soal2.go

```gopackage main

import "fmt"

func main() {
	var peserta, pemenang string
	var maxSoal, minWaktu int

	maxSoal = -1
	minWaktu = 999999

	for {
		fmt.Scan(&peserta)

		if peserta == "Selesai" {
			break
		}

		var soalSekarang, waktuSekarang int
		hitungSkor(&soalSekarang, &waktuSekarang)

		if soalSekarang > maxSoal || (soalSekarang == maxSoal && waktuSekarang < minWaktu) {
			maxSoal = soalSekarang
			minWaktu = waktuSekarang
			pemenang = peserta
		}
	}

	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
	}
}

func hitungSkor(skor, soal *int) {

	var waktu int

	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}

}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal2.png)
[penjelasan]
Penulis mendeklarasikan variable peserta dan pemenang sebagai string lalu maxsoal dan minsoal sebagai int, lalu dibuat for yang di dalamnya berisi inputan peserta lalu dibuat if then jika peserta di input selesai maka perulangan akan selesai, lalu dibuat var soalsekarang dan waktusekkarang dengn tipe data int, lalu di bawa ke func hitung skor yang dimana di func tsb dilakukan inputan menjadi skor lalu soal dengan output
