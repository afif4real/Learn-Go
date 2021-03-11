package main

import "fmt"

func main() {
	var (
		input, i, n int
	)
	fmt.Scanln(&input)
	n = 0
	i = 1

	for i < input+1 {
		if input%i == 0 { // utk mencari faktor
			fmt.Print(i, " ")
			i++
			n++
		} else if input%i != 0 { //utk mencari faktor
			i++
		}
	}
	fmt.Println()
	if n != 2 {
		fmt.Println("bukan ole2")
	} else if n == 2 {
		fmt.Println("ole2")
	}
}
