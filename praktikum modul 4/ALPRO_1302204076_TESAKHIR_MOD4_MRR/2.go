package main

import "fmt"

func main() {
	var (
		n, d1, d2 int
	)
	x := 0
	fmt.Scan(&n)
	for x < n {
		fmt.Scan(&d1, &d2)
		if (d1+d2)%2 == 0 { // kalo genap
			fmt.Println(d1)
		} else if (d1+d2)%1 == 0 { // kalo ganjil
			fmt.Println(d2)
		}
		x++ //x + 1 sebagai counter jika x == n maka looping akan berhenti
	}
}
