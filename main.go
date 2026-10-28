package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(100) + 1
	guess := 0
	fmt.Print("Enter a number from 1 to 100: ")
	fmt.Scan(&guess)
	for guess != x {
		if guess < x {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
		fmt.Print("Enter a number: ")
		fmt.Scan(&guess)
	}
	fmt.Println("You guessed it!")
}
