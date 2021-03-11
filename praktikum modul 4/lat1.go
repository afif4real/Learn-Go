package main

import (
	"fmt"
)

func main() {
	var x, i int
	var input, hasil, temp int

	fmt.Scan(&x)

	temp = 1

	i = 0
	for i < x {
		fmt.Scan(&input)
		if input > -1 && input < 10 {
			hasil = hasil + (input * temp)
			temp = temp * 10
			i = i + 1
		}
	}
	fmt.Println(hasil)
}
