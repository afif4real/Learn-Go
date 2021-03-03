package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 1; i <= 6; i++ {
		fmt.Printf("%6s\n", strings.Repeat("*", i))
	}
}
