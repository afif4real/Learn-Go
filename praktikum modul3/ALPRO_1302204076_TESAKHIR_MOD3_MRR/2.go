package main //Afif Fajar Rayhan

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(n / 1000)
	fmt.Println((n / 100) % 10)
	fmt.Println((n / 10) % 10)
	fmt.Println(n % 10)
	fmt.Println(n / 100)
	fmt.Println((n / 10) % 100)
	fmt.Println(n % 100)
	fmt.Println(n / 10)
	fmt.Println(n % 1000)
	fmt.Print(n)
}
