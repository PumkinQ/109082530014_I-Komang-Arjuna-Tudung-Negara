# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>

<p align="center">[I Komang Arjuna Tudung Negara] - [109082530014]</p>

## Unguided

### 1. [Soal]

Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai
suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat
diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku
ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci
tersebut.

#### soal1.go

```go
package main

import "fmt"

func fibo(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(fibo(n))
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal1.png)
[penjelasan]
Penulis mendeklarasikan variable n lalu mebuat inputan dan print yang dimana print tsb langsung menginputan nilai n ke func fibo, pada func fibo ketika n adalah 0 maka akan mengembalikan nilai 0, ketika n == 1 maka mengembalikan 1 dan ketika tidak keduanya maka akan mengembalikan nilai n - 1 ditambah n-2

## Unguided

### 2. [Soal]

Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan
menggunakan fungsi rekursif. N adalah masukan dari user.
Contoh masukan dan keluaran:

#### soal2.go

```go
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



```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal2.png)
[penjelasan]
Penulis mendeklarasikan variable n dan inputan n, lalu dibuat inputan ke func bintang nilai n dan 0, pada func bintang nilai tsb diubah jadi n dan f tipe data int, di dalam func tsb ketika f lebih besar sama dengan n maka tidak mengembalikan nilai apapun atau bakal stop, tapi ketika tidak maka akan mengirim nilai f ke func cetakbintang selanjutnya bakal ngeprint "_", lalu memasukan nilai n dan f+1 ke func bintang yang menyebabkan rekursif terjadi, di dalam func cetak bintang terdapat percabangan ketika f ==0 maka bakal return, dan selanjutnya print _

## Unguided

### 3. [Soal]

Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari
suatu N, atau bilangan yang apa saja yang habis membagi N.
Masukan terdiri dari sebuah bilangan bulat positif N.
Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
Contoh masukan dan keluaran:

#### soal3.go

```go
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

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%203/output/Soal.3.png)
[penjelasan]
Pada program diatas penulis mendeklarasi n bertipe data int, lalu inputan n, lalu di input ke func mod dengan nilai n dan n, lalu di mod terdapat percabangan ketika b lebih besar dari 0 maka akan menjalankan inputan ke func mod kembali dengan nilai n dan b-1 yang menyebabkan rekusif lalu dibuat percabangan lagi ketika habis dibagi b maka bakal mengeprint b yaitu nilai faktor tsb
