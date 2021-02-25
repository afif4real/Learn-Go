package main

import "fmt"

func main() {
	var n, i, jumlah, t1, t2, t3 int
	fmt.Scanln(&n)
	jumlah = 0
	i = 0
	for i < n {
		fmt.Scanln(&t1, &t2, &t3)
		if t1+t2+t3 >= 2 {
			jumlah++
		}
		i++
	}
	fmt.Println(jumlah)
}
