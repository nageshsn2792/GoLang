package main

import "fmt"

type deck []string

func NewDeck() deck{

	card := []string{}
	cardType:= [] string{"club","heart","diamond","spades"}
	cardValue:= [] string{"one","two","three","four"}

	for _, cardtype := range cardType{
		for _, cardValues := range cardValue{
			card = append(card, cardValues+ " of "+ cardtype)
		}
	}
	return card
}

func (d deck) print() {
	// Looping the deck of cards and printing it
	fmt.Print("Inside deck class print fucntion: ")
	for _, value := range d{
		fmt.Println(value);
	}
}