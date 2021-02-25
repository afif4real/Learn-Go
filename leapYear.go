package main

import (
	"fmt"
)

func main() {
	var tahun int
	fmt.Print("masukan tahun: ")
	fmt.Scan(&tahun)
	var leapYear bool = (tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0))
	fmt.Println("Tahun Kabisat:", leapYear)

}
