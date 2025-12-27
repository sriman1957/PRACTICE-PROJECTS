package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func getRandomNumber() int {
	return rand.Intn(100) + 1
}

func main() {
	fmt.Println("Number Guessing Game.....")

	time.Sleep(2 * time.Second)

	generatedNumber := getRandomNumber()

	for {
		fmt.Print("Enter a NUMBER between (1-100) : ")
		var guessedNumber int
		_, err := fmt.Scan(&guessedNumber)

		if err != nil {
			fmt.Println("Invalid input! Please enter a whole number.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if guessedNumber > 100 {
			fmt.Println("Error: You've provided wrong input, try to run the code again.")
			os.Exit(1)
		}

		if guessedNumber > generatedNumber {
			fmt.Println("Too big...,try again")
		} else if guessedNumber < generatedNumber {
			fmt.Println("Too small...,try again")
		} else {
			fmt.Println("Exactly same number,you are a genius....")
			break
		}
	}

}
