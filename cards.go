package main

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func createDeck() []card {
	var (
		rankBySuit  = map[string]int{"Club": 1, "Diamond": 2, "Heart": 3, "Spade": 4, "Joker": 5}
		suitedNames = [52]string{
			"Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs", "Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs", "Ace of Clubs",
			"Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts", "Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts", "Ace of Hearts",
			"Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds", "Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds", "Ace of Diamonds",
			"Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades", "Six of Spades", "Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades", "Ace of Spades",
		}
		glyphRef = map[string]string{ // Maps are reference types, so they are always passed by reference. You don't need a pointer.
			"Club":        "♣", // Coerce Heart Red  &#x2665;&#xFE0F; - Coerce Black &#x2665;&#xFE0E;
			"Diamond":     "♦",
			"Heart":       "♥",
			"Spade":       "♠️",
			"♣":           "U+2663",
			"♦":           "U+2666",
			"♥":           "U+2665",
			"♠":           "U+2660",
			"BLACK JOKER": "U+1F0CF",
			"🃏":           "U+1F0CF",
		}
	)

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

		if c.rankInSuit == 0 { // handles Aces and zeroes
			c.rankInSuit = 13
		}

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

func getRandCardNum(shoeSize int) int {
	shoeSize64 := int64(shoeSize) // overkill much? - max 51

	randBigInt, err := rand.Int(rand.Reader, big.NewInt(shoeSize64))
	if err != nil {
		panic(err)
	}

	randInt := int(randBigInt.Int64()) // int <- Int64 <- *big.Int

	return randInt
}

func descHand(playerHand []card) string {
	var handDesc string

	for i := range playerHand {
		nextCard := playerHand[i].suitedName
		handDesc = handDesc + nextCard + ", "
	}

	handDesc = strings.TrimSuffix(handDesc, ", ")

	return handDesc
}
