package main

import (
	"fmt"
)

func main() {

	fmt.Println("starting Project card")
	cards := newDeck()
	cards.print()

	/*
		//splitting deck
		fmt.Println("test deal function")
		hand, remainingCards := deal(cards, 5)
		hand.print()
		fmt.Println("")
		remainingCards.print()

		//deck to byte

		//cards.decttobyteslice()
		//lll := cards.decttobyteslice()

		// savingto file

		//error := ioutil.WriteFile("card.txt", cards.decttobyteslice(), 0644)
		//if error != nil {
		//	fmt.Println("error occured")
		//}

		fmt.Println("returning to deck from file ")
		dd := filetodeck("card.txt")
		dd.print()
	*/
	fmt.Println("After Shuffle")
	cards.shuffle()
	cards.print()

}
