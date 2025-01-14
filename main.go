package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n -=& Begin: carddeck v0.2.. *=-\n")
	theDeck := createDeck()
	players := []player{} //                    //?  {} at end?  struct?
	for i := 0; i < numberOfPlayers; i++ {
		p := player{idx: i}
		players = append(players, p)
		theDeck, players[i].hand = deal(theDeck, cardsPerPlayer)
		fmt.Println()
		// fmt.Println("players #: ", players[i].idx, "hand: ", players[i].hand)
	}

	// fmt.Print("\n\n")
	// fmt.Println("players #: ", players[1].idx, "- hand: ", players[1].hand)
	// fmt.Println("players #: ", players[2].idx, "- hand: ", players[2].hand)
	// fmt.Println("players #: ", players[3].idx, "- hand: ", players[3].hand)
	// fmt.Println("players #: ", players[0].idx, "- hand: ", players[0].hand)

	fmt.Print("\nOne lucky card:")
	logCard(players[0].hand[1], true)

	logDeck(theDeck, "full", true)
}
