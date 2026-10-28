package main

import (
	"fmt"
	"math/rand"
)

var (
	attempts = 7
)

func main() {
	x := rand.Intn(100) + 1
	counter := 1
	guess := 0
	fmt.Print("Enter a number from 1 to 100: ")
	fmt.Scan(&guess)
	for guess != x && counter <= attempts {
		counter++
		if guess < x {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
		fmt.Print("Enter a number: ")
		fmt.Scan(&guess)
	}
	if attempts <= counter {
		fmt.Println("You lost! The number was", x)
		return
	}
	fmt.Println("You guessed it!")
}
