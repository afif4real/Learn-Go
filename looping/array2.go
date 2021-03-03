package main

import "fmt"

func main() {
	// Iterate numeric array
	numbers := [3]int{6, 9, 3}
	for i := range numbers {
		fmt.Printf("%d ", numbers[i])
	}
}
