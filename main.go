/*
Package: carddeck - Provides a virtual card deck, and the tools most frequently required to manipulate it in a tabletop card game (without jokers?)
*/
package main

import (
	"fmt"
	"strings"
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
}

func descHand(playerHand []card) string {
	var handDesc string

	for i := range playerHand {
		thumbNext := playerHand[i].suitedName
		handDesc = handDesc + thumbNext + ", "
	}

	handDesc = strings.TrimSuffix(handDesc, ", ")

	return handDesc
}

// func sortAndScore((playerHand []card) ([]card string) {
//? Given a hand, how do we decide how to evaluate it?  What to discard?  What to bet?  Do I bluff?  Do I fold?  Do I raise?  Has the card we need been played?
//     return sortedHand, score
// }
//? player numbers from 0 or 1 ?

func establishPlayers(theDeck []card, numberOfPlayers int, cardsPerPlayer int) ([]card, []player) {
	players := []player{} //? players := []player{}  diff?  either works -  BUT?
	//                            //? var players []player      Happier wif pre-alloc from top choice?
	shoe := theDeck

	var aHand []card

	for i := range numberOfPlayers {
		shoe, aHand = deal(shoe, cardsPerPlayer)

		p := player{
			idx:       i,
			name:      playerNames[i],
			hand:      aHand,
			handScore: 0,
			wins:      0,
			losses:    0,
		}
		players = append(players, p)
	}

	return shoe, players
}
