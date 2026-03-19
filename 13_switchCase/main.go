package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case - Ludo Game")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 & u can open in snake || move 1 spot")
	case 2:
		fmt.Println("move 2 spot")
	case 3:
		fmt.Println("move 3 spot")
	case 4:
		fmt.Println("move 4 spot")
	case 5:
		fmt.Println("move 5 spot")
	case 6:
		fmt.Println("ypu can start in Ludo || move 6 spot || roll dice again")
	default:
		fmt.Println("what was that!!")
	}
}
