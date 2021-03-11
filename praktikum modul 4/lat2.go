package main

import (
	"fmt"
)

func main() {
	var n, max, orang, i, nilai int

	fmt.Scanln(&n)
	i = 0
	orang = 0
	max = 0
	for i < n {
		fmt.Scan(&nilai)
		if nilai > max {
			orang = 1
			max = nilai
		} else if nilai == max {
			orang++
		}
		i++
	}
	fmt.Printf("nilai terbesar adalah %d yang diperoleh %d orang mahasiswa", max, orang)
}
