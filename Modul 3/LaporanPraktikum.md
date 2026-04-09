# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>

<p align="center">[I Komang Arjuna Tudung Negara] - [109082530014]</p>

## Unguided

### 1. [Soal]

Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)

#### soal1.go

```go
package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println("permutation: ", permutation(a, c))
		fmt.Println("combination: ", combination(a, c))
		fmt.Println("permutation: ", permutation(b, d))
		fmt.Println("combination: ", combination(b, d))
	}
}

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutation(n, r int) int {
	hasil := factorial(n) / factorial((n - r))
	return hasil
}

func combination(n, r int) int {
	hasil := factorial(n) / (factorial(r) * factorial((n - r)))
	return hasil
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal1.png)
[penjelasan]
Penulis mendeklarasikan variable a, b, c, d. lalu membuat inputan , lalu memanggil hasil dari permutasi dan kombinasi, dimana hanya nilai n dan r yang di bawa, untuk di proses permutasi, dan kombinasi, dari dua func itu juga memanggil hasil dari factorial pada func actorial untuk di eksekusi sendiri.

## Unguided

### 2. [Soal]

Diberikan tiga buah fungsi matematika yaitu f (x) = x
2
, g (x) = x − 2 dan h (x) = x +

1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x)
   dalam bentuk function.
   Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.
   Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b),
   dan baris ketiga adalah (hofog)(c)!

#### soal2.go

```go
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


```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal2.png)
[penjelasan]
Penulis mendeklarasikan variable 1,b,c dilakukan inputan dan output di atas, lalu dibuat 3 func untuk fogog, gohof, dan hofog, lalu dari outputan di panggil langsung func tsb

## Unguided

### 3. [Soal]

[Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius
r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) < r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	L1 := didalam(cx1, cy1, r1, x, y)
	L2 := didalam(cx2, cy2, r2, x, y)

	if L1 && L2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if L1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if L2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal.3.png)
[penjelasan]
Pada program diatas penulis mendeklarasi beberapa variable lalu digunakan inputan scan, lalu dibuat func dalam untuk membandingakan dengan radius, lalu dimasukan ke variabel L1 dan L 2, lalu doberikan struktur ko ntrol if else, yang jika l1 daan l2 mempunyai nilai sama maka titik 1 dan 2, kalau tidak, maka titik dalam lingkaran 1, dan kalau tidak maka titik dalam lingkaran 2, dan kalau tidak maka titik di luar lingkaran 1 dan 2
