package main

import (
	"fmt"
	"math/rand"
)

var (
	play_again   = "y"
	max_attempts = 7
)

func main() {
	x := rand.Intn(100) + 1
	result := play_session(x)
	fmt.Println(result)
	fmt.Print("Want to play again? (y/n): ")
	fmt.Scan(&play_again)
	for play_again == "y" {
		x = rand.Intn(100) + 1
		result = play_session(x)
		fmt.Println(result)
		fmt.Print("Want to play again? (y/n): ")
		fmt.Scan(&play_again)
	}
}

func play_session(x int) string {
	result := "You lost"
	counter := 1
	guess := 0
	fmt.Print("Enter a number from 1 to 100: ")
	fmt.Scan(&guess)
	for guess != x && counter <= max_attempts {
		counter++
		if guess < x {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
		fmt.Print("Enter a number: ")
		fmt.Scan(&guess)
	}
	if x == guess {
		result = "You win!"
	}
	return result
}
