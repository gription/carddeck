package main

import (
	"math/rand/v2"
	"strings"
)

func createDeck() []card {
	accumDeck := []card{}
	for i := 1; i <= 52; i++ {
		c := card{idx: i}
		c.suitedName = suitedNames[c.idx-1]
		c.noSuitName = strings.Fields(c.suitedName)[0]
		c.suitPlural = strings.Fields(c.suitedName)[2]
		c.suitSingle = strings.TrimSuffix(c.suitPlural, "s")
		c.rankInSuit = c.idx % 13
		c.rankPip = string(rankPips[c.rankInSuit])
		c.suitPip = glyphRef[c.suitSingle]
		c.suitPipUC = glyphRef[c.suitPip]
		c.suitRank = rankBySuit[c.suitSingle]
		accumDeck = append(accumDeck, c)
	}

	return accumDeck[:]
}

func deal(theDeck []card, numberOfCards int) ([]card, []card) {
	var cardsDealt []card
	var currentCard card
	for i := 1; i <= numberOfCards; i++ {
		theDeck, currentCard = pickACard(theDeck)
		logCard(currentCard, true)
		cardsDealt = append(cardsDealt, currentCard)
	}
	return theDeck, cardsDealt[:]
}

func pickACard(theDeck []card) ([]card, card) {
	cardDrawn := theDeck[rand.IntN(len(theDeck))]
	theDeck = removeCard(theDeck, cardDrawn)
	return theDeck, cardDrawn
}

func removeCard(theDeck []card, cardToRemove card) []card {
	var idx int = findCardIndex(theDeck, cardToRemove)
	theDeck[idx] = theDeck[len(theDeck)-1] // move to back line
	theDeck = theDeck[:len(theDeck)-1]     // then lop it off - fast
	return theDeck
}

func findCardIndex(theDeck []card, cardInQuestion card) int {
	for idx, thatCard := range theDeck {
		if thatCard == cardInQuestion {
			return idx
		}
	}
	return -1 // no card found
}
