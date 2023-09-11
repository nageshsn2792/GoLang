package main

import "fmt"

func main() {
	// Assigning value to variables
	hello := "Hello!!!"
	fmt.Println(hello)
	// Integer type variable
	programno := 1
	fmt.Println(programno)
	// Printing the String value in fmt.Println statement
	fmt.Println("Welcome to cards game!")

	//Creating the array of string values
    cards:= deck{"1 of hearts", "2 of sphades", "7 of diamonds", "8 of Khilwar"}

	// Looping the arrayand printing the values
	for _, value := range cards{
		fmt.Println(value);
	}

	// Calling a function written in another go file and retrieve the deck of cards
	cardsDeck:= NewDeck()

	// Calling a function to print the deck of cards
	cardsDeck.print()
}