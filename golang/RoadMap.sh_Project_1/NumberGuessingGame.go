package main

import (
	"fmt"

)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Print("\nPlease select the difficulty level:")
	fmt.Println("\n1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	var level int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&level)

	switch level {
	case 1:
		fmt.Println("Great! You have selected the Easy difficulty level.")
	case 2:
		fmt.Println("Great! You have selected the Medium difficulty level.")
	case 3:
		fmt.Println("Great! You have selected the Hard difficulty level.")
	default:
		fmt.Println("Invalid Selection , Select only 1, 2 or 3")
		return
	}
	fmt.Println("Let's start the game!")
}
