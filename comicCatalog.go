package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pageNum uint
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hachet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."

	year = 1997
	pageNum = 14
	grade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "is published by", publisher, "in the year", year, "It has a total of:", pageNum, "pages and this copy has a grade of:", grade)

	fmt.Println("---------------------------------------------------------------------------------")

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNum = 160
	grade = 9.0
	fmt.Println(title, "written by", writer, "drawn by", artist, "is published by", publisher, "in the year", year, "It has a total of:", pageNum, "pages and this copy has a grade of:", grade)

}
