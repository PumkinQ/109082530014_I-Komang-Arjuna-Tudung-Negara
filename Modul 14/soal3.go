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
