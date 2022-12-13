package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	println("You leave the cave: ", exit)

	exit = strings.Contains(command, "inside")

	println("You leave the cave: ", exit)

	var wearShades = true
	println("Quick check 3.1", wearShades)

	command = "There is a sign near the entrance that reads 'No Minors'"
	exit = strings.Contains(command, "reading")

	println("At age 41, am I a minor?", exit)
}
