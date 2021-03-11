package main //blm bener

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		seed          int64
		result, guess int
		resultDadang  int
	)
	fmt.Println("bilangan rahasia:")
	fmt.Scan(&seed)
	fmt.Println("tebakan anda:")
	fmt.Scan(&guess)
	rand.Seed(seed)
	resultDadang = rand.Intn(6) + 1
	fmt.Printf("Dadang: %d", resultDadang)
	result = rand.Intn(6) + 1

	if guess == result {
		fmt.Printf("Nilai dadu &d pemenang adalah anda", result)
	} else if resultDadang == result {
		fmt.Printf("Nilai dadu &d pemenang adalah Dadang", resultDadang)
	} else if guess == resultDadang && guess == result {
		fmt.Println("Nilai dadu", result, "seri")
	} else {
		fmt.Println("Nilai dadu", result, "Tidak ada pemenang")
	}
}
