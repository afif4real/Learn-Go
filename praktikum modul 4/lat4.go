package main //blm selesai

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		seed           int64
		n              int
		tebak1, tebak2 string
	)
	//kemungkinan : ganjil besar (5), ganjil kecil(1 3),
	//genap besar (4 6), genap kecil (2)


	fmt.Print("Angka rahasia: ")
	fmt.Scanln(&seed)
	rand.Seed(seed)
	fmt.Print("Anda: ")
	fmt.Scanln(&tebak1, &tebak2)
	n = rand.Intn(6) + 1
	if n%2 == 1 { //GANJIL
		if n >= 4 {
			if tebak1 == "ganjil" {
				if tebak2 == "kecil" {
					fmt.Printf("Nilai dadu %d, Anda ka
if tebak2 == "kecil" {
					fmt.Printf("Nilai dadu %d, Anda kalah\n", n)
				} else if tebak2 == "besar" {
					fmt.Printf("Nilai dadu %d, Anda menang\n", n)
				}
			} else if tebak1 == "genap" {
				fmt.Printf("Nilai dadu %d, Anda kalah\n", n)
			}
} else {
			if tebak1 == "ganjil" {
				if tebak2 == "kecil" {
					fmt.Printf("Nilai dadu %d, Anda menang\n", n)
				} else if tebak2 == "besar" {
					fmt.Printf("Nilai dadu %d, Anda kalah\n", n)
				}
			} else if tebak1 == "genap" {
				fmt.Printf("Nilai dadu %d, Anda kalah\n", n)
			}
		}
} else if n%2 == 0 {
		if n >= 4 {
			if tebak1 == "genap" {
				if tebak2 == "kecil" {
					fmt.Printf("Nilai dadu %d, Anda kalah\n", n)
				} else if tebak2 == "besar" {
					fmt.Printf("Nilai dadu %d, Anda menang\n", n)
				}
			} else if tebak1 == "ganjil" {
				fmt.Printf("Nilai dadu %d, Anda kalah\n", n)
			}
} else { //if n < 4
			if tebak1 == "genap" {
				if tebak2 == "kecil" {
					fmt.Printf("Nilai dadu %d, Anda menang\n", n)
				} else if tebak2 == "besar" {
					fmt.Printf("Nilai dadu %d, Anda kalah\n", n)
				}
			} else if tebak1 == "ganjil" {
				fmt.Printf("Nilai dadu %d, Anda kalah\n", n)
			}
		}
	}
}