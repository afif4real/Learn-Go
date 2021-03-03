package main

import "fmt"

func main() {
	var (
		a, b, c, d, e int
		k1, k2, k3    byte
	)
	fmt.Scanln(&a, &b, &c, &d, &e)
	fmt.Scanf("%c%c%c", &k1, &k2, &k3) //%c = karakter, %d = int, %t = bool, %f = float, %s = string
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)
	fmt.Printf("%c%c%c", k1+1, k2+1, k3+1)
}
