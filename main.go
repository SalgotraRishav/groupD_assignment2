package groupD_assignment2

import (
	"fmt"
)

func main() {
	// Calling the function
	greatestOfThree()
}

func greatestOfThree() {
	//new commit
	//new commit
	// Prompt user for input
	fmt.Println("Enter three numbers:")

	// Taking three variables of integer type
	var num1, num2, num3 int

	// Read inputs from user and save in num1, num2 and num3 variables
	fmt.Print("Number 1: ")
	fmt.Scanln(&num1)
	fmt.Print("Number 2: ")
	fmt.Scanln(&num2)
	fmt.Print("Number 3: ")
	fmt.Scanln(&num3)

	// Calling the 'greatest' function and storing the return value in num variable
	num := greatest(num1, num2, num3)

	// Displaying the output on the screen
	fmt.Println(num, "is the greatest")
}

// function to calculate the greatest out of three
func greatest(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}
