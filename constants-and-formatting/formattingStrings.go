package main

import "fmt"

func main() {
	//	%v variant prints any value in a default format
	s := fmt.Sprintf("I am %v years old", 10)
	// I am 10 years old
	s2 := fmt.Sprintf("I am %v years old", "way too many")
	// I am way too many years old

	//	Strings
	stri := fmt.Sprintf("I am %s years old", "way too many")
	// I am way too many years old

	//	Integer
	intgr := fmt.Sprintf("I am %d years old", 10)
	// I am 10 years old

	//	Float
	floa1 := fmt.Sprintf("I am %f years old", 10.523)
	// I am 10.523000 years old

	// The ".2" rounds the number to 2 decimal places
	floa2 := fmt.Sprintf("I am %.2f years old", 10.523)
	// I am 10.52 years old

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(stri)
	fmt.Println(intgr)
	fmt.Println(floa1)
	fmt.Println(floa2)
}
