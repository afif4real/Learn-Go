package main //Afif Fajar Rayhan

import "fmt"

func main() {
	var (
		s1, s2, s3, s4, s5 rune
		n                  rune
	)
	fmt.Scanf("%c%c%c%c%c", &s1, &s2, &s3, &s4, &s5)
	fmt.Scan(&n)
	fmt.Printf("%c%c%c%c%c", s1+n, s2+n, s3+n, s4+n, s5+n)

}
