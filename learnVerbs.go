package main

import "fmt"

//disini kita belajar .format atau f{} punyanya golang yaitu verbs
//buat int pake %d
//buat string mah pake %v aja
func main() {
	floatExample := 1.75
	// Edit the following Printf for the FIRST step
	fmt.Printf("Working with a %T.", floatExample)

	fmt.Println("\n***") // Added for spacing

	yearsOfExp := 3
	reqYearsExp := 15
	// Edit the following Printf for the SECOND step
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years.", yearsOfExp, reqYearsExp)

	fmt.Println("\n***") // Added for spacing

	stockPrice := 3.50
	// Edit the following Printf for the THIRD step
	fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice)
}
