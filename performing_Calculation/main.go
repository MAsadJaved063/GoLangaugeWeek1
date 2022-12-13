package main

import "fmt"

//main function is where the it all begin
func main() {
	fmt.Print("My weight on the surface of mars is ")
	fmt.Print(149 * 0.3783)
	fmt.Print(" lbs and i would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.\n")

	//pass list of arguments in print() function
	fmt.Print("My weight on the surface of mars is ", 149*0.3783, " Ibs and i would be ", 41*365.3456/687, " years old.")

	num1 := 2
	num2 := 3
	sum := num1 + num2
	subtraction := num1 - num2
	multiplication := num1 * num2
	division := num2 / num1
	println("Sum of the numbers is ", sum)
	println("Subtarction of the numbers is ", subtraction)
	println("Multiplication of the numbers is ", multiplication)
	println("Division of the numbers is ", division)

}
