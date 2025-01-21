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

	//. lucky card (how to address a player card)
	// fmt.Print("One lucky card for the dealer:")
	// logCard(players[0].hand[1], true)

	// . Sort, Score & Reveal the players' cards..
	for i := range players {
		fmt.Print(players[i].name, ": ", descHand(players[i].hand), "\n")
		players[i].hand, players[i].handScore, players[i].kickers, players[0].tieBreakPipStr = sortAndScore(players[i].hand)

		fmt.Println("----------------------------------------------------")
	}

	//* DEAL FROM THE BOTTOM... (insensitive to deck contents)
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeFlush()
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeStraight()
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeStraightFlush()
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeRoyalFlush()
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeHighCard()
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeOnePair()
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeTwoPair()
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeTrips()
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeQuads()
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeFullHouse()
	// players[0].hand, players[0].handScore, players[0].kickers, players[0].tieBreakPipStr = gimmeTrash()
	// logPlayers(players[0])
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
}
