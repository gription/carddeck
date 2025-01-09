package main

import "fmt"

func logDeck(deck []card, format string, yorn bool) int {
	if yorn {
		fmt.Println("\n -=*", len(deck), "Cards remain in deck:")
		for i := 0; i < len(deck); i++ {
			if format == "full" {
				c := deck[i]
				if i%6 == 0 {
					fmt.Print("\n")
				}
				if len(c.suitedName) > 15 {
					fmt.Print(c.deckSeq, "\t", c.suitedName, "\t")
				} else {
					fmt.Print(c.deckSeq, "\t", c.suitedName, "\t\t")
				}
			}
		}
	}
	fmt.Print("\n")
	return 0
}

func createDeck() []card {
	accumDeck := []card{}
	for i := 1; i <= 52; i++ {
		accumDeck = append(accumDeck, makeCard(i))
	}
	return accumDeck[:] // returns slice
}

func deal(theDeck []card, numberOfCards int) ([]card, []card) {
	var cardsDealt []card
	var currentCard card
	for i := 1; i <= numberOfCards; i++ {
		theDeck, currentCard = pickACard(theDeck)
		logCard(currentCard, false)
		cardsDealt = append(cardsDealt, currentCard)
	}
	return theDeck, cardsDealt[:]
}
