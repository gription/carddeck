/*
Package: carddeck - Provides a virtual card deck, and the tools most frequently required to manipulate it in a tabletop card game (without jokers?)
*/
package main

import (
	"fmt"
)

// shoe dealt from deck

func main() {
	fmt.Printf("\n -=& Begin: carddeck v0.2.. *=-\n")

	theDeck := createDeck()
	theDeck, players := establishPlayers(theDeck, numberOfPlayers, cardsPerPlayer)
	logDeck(theDeck, "full", true)
	logPlayers(players)
	fmt.Print("One lucky card for the dealer:")
	logCard(players[0].hand[1], true)

	num := len(theDeck)
	// fmt.Println()
	randCardNum := getRandCardNum(num)
	fmt.Println("Random num [0-ShoeSize]: ", randCardNum)
}

// func sortAndScore((playerHand []card) ([]card string) {
//? Given a hand, how do we decide how to evaluate it?  What to discard?  What to bet?  Do I bluff?  Do I fold?  Do I raise?  Has the card we need been played?
//     return sortedHand, score
// }
