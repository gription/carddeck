package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n -=& Begin: carddeck v0.2.. *=-\n")
	theDeck := createDeck()

	//, Create x players, each holding x cards.  highRollers[p#][hand][cards]{card-details}
	highRollers := []player{} //?  {} at end?  struct?
	for i := 0; i < numberOfPlayers; i++ {
		p := player{idx: i}
		highRollers = append(highRollers, p)
		theDeck, highRollers[i].hand = deal(theDeck, cardsPerPlayer)
		fmt.Println()

		// fmt.Println("highRollers #: ", highRollers[i].idx, "hand: ", highRollers[i].hand)
	}

	// fmt.Print("\n\n")
	// fmt.Println("highRollers #: ", highRollers[1].idx, "- hand: ", highRollers[1].hand)
	// fmt.Println("highRollers #: ", highRollers[2].idx, "- hand: ", highRollers[2].hand)
	// fmt.Println("highRollers #: ", highRollers[3].idx, "- hand: ", highRollers[3].hand)
	// fmt.Println("highRollers #: ", highRollers[0].idx, "- hand: ", highRollers[0].hand)

	fmt.Print("\nOne lucky card:")
	logCard(highRollers[0].hand[1], true)

	logDeck(theDeck, "full", true)
}

// ? are all linters engaging?
