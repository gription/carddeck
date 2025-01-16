package main

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func getRandCardNum(shoeSize int) int {
	shoeSize64 := int64(shoeSize) // max 51

	randBigInt, err := rand.Int(rand.Reader, big.NewInt(shoeSize64))
	if err != nil {
		panic(err)
	}

	randInt := int(randBigInt.Int64()) // int <- Int64 <- *big.Int

	return randInt
}

func createDeck() []card {
	accumDeck := []card{}

	for i := 1; i <= 52; i++ {
		c := card{
			idx:        i,
			suitedName: "",
			noSuitName: "",
			suitPlural: "",
			suitSingle: "",
			rankInSuit: 0,
			rankPip:    "",
			suitPip:    "",
			suitPipUC:  "",
			suitRank:   0,
		}
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
	modifiedDeck := theDeck

	var cardsDealt []card

	var currentCard card

	for i := 1; i <= numberOfCards; i++ {
		modifiedDeck, currentCard = pickACard(modifiedDeck)
		logCard(currentCard, false)
		cardsDealt = append(cardsDealt, currentCard)
	}

	return modifiedDeck, cardsDealt[:]
}

func pickACard(theDeck []card) ([]card, card) {
	cardDrawn := theDeck[getRandCardNum(len(theDeck))]
	modifiedDeck := removeCard(theDeck, cardDrawn)

	return modifiedDeck, cardDrawn
}

func removeCard(theDeck []card, cardToRemove card) []card {
	idx := findCardIndex(theDeck, cardToRemove)

	theDeck[idx] = theDeck[len(theDeck)-1]   // move to back line
	modifiedDeck := theDeck[:len(theDeck)-1] // then lop it off - fast

	return modifiedDeck
}

func findCardIndex(theDeck []card, cardInQuestion card) int {
	for idx, thatCard := range theDeck {
		if thatCard == cardInQuestion {
			return idx
		}
	}

	return -1 // no card found
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
