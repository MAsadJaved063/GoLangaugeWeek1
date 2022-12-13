package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch case in go language")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice value is ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 you can open")
	case 2:
		fmt.Println("Dice value is 2 you can move 2 spot")
	case 3:
		fmt.Println("Dice value is 3 you can move 3 spot")
	case 4:
		fmt.Println("Dice value is 4 you can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 you can move 5 spot")
		fallthrough //Fallthrough means that its next case will be executed
	case 6:
		fmt.Println("Dice value is 6 you can move 6 spot")
	default:
		fmt.Println("What was that")
	}
}
