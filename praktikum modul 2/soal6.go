package main

import "fmt"

func main() {
	var nilai, n float64
	var jumlah, rata2 float64
	fmt.Scanln(&nilai)
	n = 0
	for nilai != -1 {
		n++
		jumlah += nilai
		fmt.Scanln(&nilai)
	}

	if n == 0 {
		rata2 = 0.0
	} else {
		rata2 = jumlah / n
	}
	fmt.Println(rata2)
}
