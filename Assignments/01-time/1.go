package main

import "fmt"

/*
*
Assignment
Keeping track of time in a message-sending application is critical. Imagine getting an appointment reminder an hour after your doctor's visit.

Complete the code using a computed constant to print the number of seconds in an hour.
*/
func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Print("number of seconds in an hour:", secondsInHour)
}
