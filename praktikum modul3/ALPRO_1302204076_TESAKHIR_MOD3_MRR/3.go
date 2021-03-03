package main //Afif Fajar Rayhan

import "fmt"

func main() {
	var x1, x2, x3, x4, x5 int
	var y1, y2, y3, y4, y5 string

	fmt.Scanln(&x1, &y1, &x2, &y2, &x3, &y3, &x4, &y4, &x5, &y5)

	quads := (x1 == x2 && x1 == x3 && x1 == x4 && x1 != x5) || (x1 == x2 && x1 == x3 && x1 != x4 && x1 == x5) || (x1 == x2 && x1 != x3 && x1 == x4 && x1 == x5) || (x1 != x2 && x1 == x3 && x1 == x4 && x1 == x5)

	fmt.Println(quads)
}
