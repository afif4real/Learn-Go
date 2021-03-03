package main

import "fmt"

func main() {
	var strings [3]string
	var numbers [3]int

	strings[0] = "hello world"
	strings[1] = "hello world"
	strings[2] = "hello world"

	numbers[0] = 21
	numbers[1] = 2
	numbers[2] = 3

	fmt.Println(strings[2])
	fmt.Println(numbers[2])
}
