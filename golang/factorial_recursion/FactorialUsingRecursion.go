package main

import "fmt"

func factorial(n int) int { // using recursion finding factorial
	if n == 1 || n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	var n int
	fmt.Print("Enter n :")
	fmt.Scan(&n)

	fmt.Println(factorial(n)) // i/p = 5 , o/p = 120
}
