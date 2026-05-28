# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>

<p align="center">[I Komang Arjuna Tudung Negara] - [109082530014]</p>

## Unguided

### 1. [Soal]

Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya
Hercules sangat suka mengunjungi semua kerabatnya itu.
Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program
rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut
membesar menggunakan algoritma selection sort.

#### soal1.go

```go
package main

import "fmt"

func main() {

	var n, m, idx_min, temp int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Scan(&m)
		var arr []int
		for j := 0; j < m; j++ {

			fmt.Scan(&temp)
			arr = append(arr, temp)

		}
		for x := 0; x < m-1; x++ {
			idx_min = x
			for y := x + 1; y < m; y++ {
				if arr[y] < arr[idx_min] {
					idx_min = y
				}
			}
			arr[x], arr[idx_min] = arr[idx_min], arr[x]
		}
		fmt.Println("HASIL Loop ke-", i+1)
		for k := 0; k < m; k++ {
			fmt.Print(arr[k], " ")
		}
		fmt.Println()
		fmt.Println()
	}

}


```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal1.png)
[penjelasan]
Penulis membuat inputan n yaitu jumlah dearah dan di dalam setiap daerah melakukan perulangan untuk mengambil m yaitu banyaknya rumah, lalu mengisi setiap angka ke dalam slice. Setelah data terkumpul, penulis menerapkan algoritma Selection Sort untuk mengurutkan nomor rumah secara membesar. Selanjutnya, program mengeprint isi slice tersebut secara keseluruhan setelah terurut.

## Unguided

### 2. [Soal]

Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu
diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena
nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah
program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil
lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor
genap terurut mengecil.

#### soal2.go

```go
package main

import "fmt"

func main() {
	var n, m, temp int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			fmt.Scan(&temp)
			if temp%2 != 0 {
				ganjil = append(ganjil, temp)
			} else {
				genap = append(genap, temp)
			}
		}

		for x := 0; x < len(ganjil)-1; x++ {
			idx := x
			for y := x + 1; y < len(ganjil); y++ {
				if ganjil[y] < ganjil[idx] {
					idx = y
				}
			}
			ganjil[x], ganjil[idx] = ganjil[idx], ganjil[x]
		}

		for x := 0; x < len(genap)-1; x++ {
			idx := x
			for y := x + 1; y < len(genap); y++ {
				if genap[y] > genap[idx] {
					idx = y
				}
			}
			genap[x], genap[idx] = genap[idx], genap[x]
		}

		for _, val := range ganjil {
			fmt.Print(val, " ")
		}
		for _, val := range genap {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}


```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal2.png)
[penjelasan]
Penulis membuat input n untuk jumlah daerah yang harus diproses. lalu setiap perulangan daerah program membaca m dan langsung memilah setiap nomor rumah yang diinput ke dalam dua slice, selanjutnaya dibuat loop, satu untuk bilangan ganjil dan satu untuk bilangan genap, dengan logika temp % 2 != 0.

## Unguided

### 3. [Soal]

Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan
tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan
sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu
problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba
untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?

#### soal3.go

```go
package main

import "fmt"

func main() {
	var data []int
	var input int

	for {
		fmt.Scan(&input)
		if input == -5313 {
			break
		}

		if input > 0 {

			data = append(data, input)

			n := len(data)
			for i := n - 1; i > 0 && data[i] < data[i-1]; i-- {
				data[i], data[i-1] = data[i-1], data[i]
			}
		} else if input == 0 {

			n := len(data)
			if n%2 != 0 {

				fmt.Println(data[n/2])
			} else {
				median := (data[n/2-1] + data[n/2]) / 2
				fmt.Println(median)
			}
		}
	}
}


```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal.3.png)
[penjelasan]
Penulis membuat slice kosong dan membuat perulangan infinit untuk membaca input bilangan secara terus-menerus. Setiap kali bilangan positif dimasukkan maka program akan menambahkan bilangan tersebut ke dalam slice dan langsung melakukan Insertion Sort dengan cara menyisipkan bilangan baru ke posisi yang tepat agar slice selalu dalam kondisi terurut.Kalau program menemukan input bernilai 0, maka akan menghitung nilai median dari data yang sudah terkumpul
