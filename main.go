/*
Package: carddeck - Provides a virtual card deck, and the tools most frequently required to manipulate it in a tabletop card game (without jokers?)
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n --==< Begin: carddeck v0.2.. >==--")

	theDeck := createDeck()
	theDeck, players := establishPlayers(theDeck, numberOfPlayers, cardsPerPlayer)
	logDeck(theDeck, "full", true)

	//. lucky card
	// fmt.Print("One lucky card for the dealer:")
	// logCard(players[0].hand[1], true)

	//. Deal the cards..
	// for i := range players {
	// 	players[i].hand, players[i].handScore = sortAndScore(players[i].hand)
	// }

	//* DEAL FROM THE BOTTOM...
	// players[0].hand, players[0].handScore = gimmeFlush()
	// players[0].hand, players[0].handScore = gimmeStraight()
	// players[0].hand, players[0].handScore = gimmeStraightFlush()
	// players[0].hand, players[0].handScore = gimmeRoyalFlush()

	// players[0].hand, players[0].handScore = gimmeHighCard()
	players[0].hand, players[0].handScore = gimmeOnePair()
	// players[0].hand, players[0].handScore = gimmeTwoPair()
	// players[0].hand, players[0].handScore = gimmeTrips()
	// players[0].hand, players[0].handScore = gimmeQuads()
	// players[0].hand, players[0].handScore = gimmeFullHouse()
	// players[0].hand, players[0].handScore = gimmeTrash()
	logPlayers(players)
}
