package main

import "fmt"

/*
*
Assignment
Create a new variable called msg on line 11 and use the appropriate formatting function to return a string that contains the following:

# Hi NAME, your open rate is OPENRATE percentNEWLINE

Replace NAME with the variable name,
Replace OPENRATE with the variable openRate rounded to the nearest "tenths" place, e.g 10.523 should be rounded down to 10.5
The word percent should appear as part of the string following the open rate value
Replace NEWLINE with the newline \n escape sequence.
For example, with the inputs "Jimmy McGill" and 2.5, the expected output would be:

Hi Jimmy McGill, your open rate is 2.5 percent
*/
func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	// don't edit above this line

	msg := fmt.Sprintf("Hi %s, your open rate is %v percent\n", name, openRate)

	// ?

	// don't edit below this line

	fmt.Print(msg)
}
