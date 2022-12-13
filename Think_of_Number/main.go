package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)

	var distance = rand.Intn(345000001) + 56000000
	fmt.Println(distance)

	//How fast a ship would be Program
	const hoursPerDay = 24
	var Distance = 56000000     //km
	var time = 28 * hoursPerDay //days
	var velocity = Distance / time

	fmt.Printf("Ship would need to travel %v (in km/h)  in order to reach Malacandra in 28 days", velocity)

}
