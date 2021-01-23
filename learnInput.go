package main

import "fmt"

//meminta inputan user di golang menggunakan Scan atau Scanln
//fmt.Scan = utk inputan angka atau 1 kata
//fmt.Scan = utk inputan > 1 kata
//disini memakai %d karena lebih suitable dgn int

func main() {
	// Add your code below:
	var angka int
	fmt.Println("masukan angka")
	fmt.Scan(&angka)

	fmt.Printf("angka: %d", angka)
}
