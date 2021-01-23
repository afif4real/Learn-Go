package main

import "fmt"

//disini kita belajar ttg Sprint basically sprint itu sama dengan return di python
//Sprintf = kyk return terus pake .format gt
func main() {
	step1 := "Breathe in..."
	step2 := "Breathe out..."

	// Add your code below:
	meditation := fmt.Sprintln(step1, step2)
	fmt.Print(meditation)

	fmt.Print("-------------------------------------------------------\n")

	template := "I wish I had a %v."
	pet := "puppy"

	var wish string

	// Add your code below:
	wish = fmt.Sprintf(template, pet)

	fmt.Println(wish)
}
