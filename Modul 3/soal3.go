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
