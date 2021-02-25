package main

import "fmt"

func main() {
	var x int
	fmt.Println("masukan angka")
	fmt.Scan(&x)
	if x%3 == 0 {
		fmt.Println("Fizz")
	} else if x%5 == 0 {
		fmt.Println("Bazz")
	}
}
