package main

import "fmt"

/**
Build a Player Summary
Create a small Go program that uses variables with the correct types and prints a short player summary.

Task
Inside main.go:

Create a playerName variable for the player's name: Mira
Create a coins variable for the starting coin amount: 42
Create a potionCost variable for the cost of one potion: 5
Create an isVIP variable to show that the player is a VIP: true
Create a variable for the coins left after buying 3 potions
Print these four lines, in this order:
Player: Mira
Coins before: 42
Coins after: 27
VIP: true
Rules
Use variables for the values
Use code to calculate the coins left after buying 3 potions
Use the correct Go types for each variable
Print only those four lines
*/

func main() {
	playerName := "Mira"
	coins := 42
	potionCost := 5
	isVIP := true
	coinsAfter := coins - (potionCost * 3)
	fmt.Println("Player:", playerName)
	fmt.Println("Coins before:", coins)
	fmt.Println("Coins after:", coinsAfter)
	fmt.Println("VIP:", isVIP)

}
