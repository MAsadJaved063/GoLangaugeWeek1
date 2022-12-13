package main

import "fmt"

func main() {
	const lightSpeed = 299792 //km/s
	var distance = 56000000   //km
	fmt.Println(distance/lightSpeed, " Seconds")

	distance = 401000000

	fmt.Println(distance/lightSpeed, " Seconds")

	//Quick 2.3
	const hoursPerDay = 24
	speed := 100800     //km/h
	distance = 96300000 //km

	result := (distance / speed) / hoursPerDay
	fmt.Printf("It will take %v Days", result)

}
