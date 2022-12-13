package main

import "fmt"

func main() {
	//Declare the variable in single line

	var num1, num2 int = 3, 9
	fmt.Printf("Number one is %v and number two is %v \n", num1, num2)
	//Assignment Operator
	var weight = 149.0
	weight = weight * 0.378
	fmt.Println("weight is ", weight)
	weight *= 0.2
	fmt.Printf("Now weight is %0.1f", weight)
	var age = 41
	age++
	fmt.Println("\nIncremented age is ", age)

}
