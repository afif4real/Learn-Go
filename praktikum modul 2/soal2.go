package main

import "fmt"

func main() {
	var luas float64
	var jari int
	fmt.Scanln(&jari)
	luas = (22.0 / 7.0 * float64(jari))
	fmt.Println("Luas lingkaran dengan jari-jari = ", jari, "adalah", luas)
}
