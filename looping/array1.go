package main // array iteration example

import "fmt"

func main() {
	// Iterate numeric array
	numbers := [3]int{6, 9, 3}
	for key, value := range numbers {
		fmt.Printf("%d = %d\n", key, value)
	}
}
