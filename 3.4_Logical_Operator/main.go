package main

import "fmt"

func main() {
	var haveTorch = true
	var litTorch = false
	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	}
	var year = 2000
	var IsLeapOrNot = (year%4 == 0 && year%100 != 0) || year%400 == 0
	if IsLeapOrNot {
		fmt.Println("Year is leap year bro")
	} else {
		fmt.Println("Not a leap year")
	}
}
