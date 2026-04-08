package main

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
