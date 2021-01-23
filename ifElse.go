package main

import "fmt"

//comparison disini sama kek python
//pake == dan != juga
//&& = and
//|| = or
//! = not
func main() {
	lockCombo := "2-35-19"
	robAttempt := "1-1-1"

	// Add your code below:
	if lockCombo == robAttempt {
		fmt.Println("The vault is now opened")
	} else {
		fmt.Println("yauda apaan kek")
	}
	rightTime := true
	rightPlace := true

	// Edit this condition for the FIRST checkpoint
	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}
}
