package main

import "fmt"

func main() {
	var x, a, b byte
	fmt.Scanln(&a, &b)
	x = a
	for x <= b {
		fmt.Printf("Simbol ASCII dari %d adalah %c\n", x, x)
		x = x + 1
	}
}
