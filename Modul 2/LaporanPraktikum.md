# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>

<p align="center">[I Komang Arjuna Tudung Negara] - [109082530014]</p>

## Unguided

### 1. [Soal]

Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?

#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%202/output/Soal1.png)
[penjelasan]
Penulis mendeklarasikan variable satu dua tiga dan juga temp sebagai string, lalu dibuat inputan pada variabel satu dua dan tiga, selanjutnya akan dilakukan penambahan variable string satu + dua + tiga, lalu selanjutnya dilakukan pemindahan nilai variable, variable satu dimasukan ke variable temp(tempat meletakkan data sementara) dua diletakan di variable satu, tiga diletakan di variabel dua, dan temp diletakan di variabel tiga, yang dimana jadinya nilai dari variabel temp adalah satu, nilai dari variabel satu adalah dua, nilai dari variabel dua adalah tiga, dan nilai dari variabel tiga adalah isi dari variabel temp yaitu satu, jadi output akhirnya menjadi dua tiga satu

## Unguided

### 2. [Soal]

Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan
praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana
susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta
untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan
warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,
‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.

#### soal2.go

```go
package main

import "fmt"

//Program Mengitung hasil dan peserta donasi
func main() {
	var w1, w2, w3, w4, acuan1, acuan2 string
	var hasil bool
	i := 1
	hasil = true
	fmt.Printf("Percobaan %d:", i)
	fmt.Scan(&w1, &w2, &w3, &w4)
	acuan1 = w1 + w2 + w3 + w4
	for i = 2; i <= 5; i++ {
		fmt.Printf("Percobaan %d:", i)
		fmt.Scan(&w1, &w2, &w3, &w4)

		acuan2 = w1 + w2 + w3 + w4
		if acuan2 != acuan1 {
			hasil = false
		}
	}
	fmt.Print("Hasil :")
	fmt.Print(hasil)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%202/output/Soal2.png)
[penjelasan]
Penulis mendeklarasikan variable w1, w2, w3, w4, acuan1, acuan2 bertipe data string dan
variabel hasil bertipe data boolean, lalu dideklarasikan i bertipe data integer bernilai 1,
dan variabel hasil bernilai true, kemudian di buat output untuk percobaan pertama
sekaligun inputan untuk percobaan pertama, selanjutnya dibuatkan variabel acuan1 yang
mana variabel itu untuk menyimpan nilai dari inputan pertama, kemudian dibuat struktur
kontrol for loop dari i = 2 sampai 5, di dalam for loop terdapat output percobaan dari 2
sampai 5 lalu terdapat inputan w1, w2, w3, w4, lalu di masukan ke variabel acuan2, dan
dibuat lagi struktur kontrol if else yang jika acuan2 tidak sama dengan acuan1 maka
variabel hasil akan false, selanjutnya akan dibuat output hasil.

## Unguided

### 3. [Soal]

PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka,
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan
sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

#### soal3.go

```go
package main

import "fmt"

func main() {
	var parsel int
	var kg int
	var gram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&parsel)
	kg = parsel / 1000
	gram = (parsel % 1000)
	fmt.Println("Detail berat: ", kg, "kg +", gram, "gr")
	detailkg := kg * 10000
	if kg >= 10 {
		gram = 0
	} else if gram < 500 {
		gram = gram * 15
	} else if gram >= 500 {
		gram = gram * 5
	}
	fmt.Println("Detail biaya: Rp.", detailkg, "+ Rp.", gram)
	total := detailkg + gram
	fmt.Print("Total biaya: Rp.", total)
}

```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/PumkinQ/109082530014_I-Komang-Arjuna-Tudung-Negara/blob/main/Modul%202/output/Soal3.png)
[penjelasan]
Pada program diatas penulis mendeklarasikan variable parsel, kg, gram sebagai integer
lalu membuat inputan parsel, kemudian melakukan konversi untuk mendapatkan nilai
gram dan kg dengan cara membagi 1000, gram dengan cara memoodulus parsel dengan
1000, lalu untuk menkonversi nilai kg ke rupiah di kali 10000, kemudian dibuat if else
jika kg >= 10 maka gram akan bernilai 0, jika gram < 500 maka gram akan dikali 15, dan
jika gram >= 500 maka gram dikali 5, juga pada contoh 3 yang diberikan ada kesalahan penulisan
detail biaya yang apabila total kg lebih dari 10kg maka sisa berat akan digeratiskan
