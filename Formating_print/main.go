package main

import "fmt"

func main() {
	//use of printf with %v
	fmt.Printf("My weight on the surface of mars is %0.1f Ibs ", 149*0.378)
	fmt.Printf("and i would be %v years old.\n", 41*365/687)

	//align text
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-5v $%4v\n", "Virgin Galactic", 100)
}
