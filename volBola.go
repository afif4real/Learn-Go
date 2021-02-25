package main

import (
	"fmt"
)

func main() {
	fmt.Println("jari-jari:")
	var jari float64 = 5
	fmt.Scanln(&jari)
	var vol float64 = ((4 / 3) * 3.14 * (jari * jari * jari))
	var luas float64 = (4.0 * 3.14 * (jari * jari))
	fmt.Println("bola dengan jejari", jari, "memiliki volume", vol, "dan luas kulit", luas)
}
