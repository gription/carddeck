package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

// TODO: Maintain central deck by ref rather than constantly passing value?
// REVIEW: Any need to shuffledDeck if all 'deals/draws' use RNG?
// TODO: Break out universal deck management into separate files (isolate from game logic: deckManagement.go, fiveCardDraw.go, sevenCardStud.go ?)
// TODO: Intro queries for game type, player count, player name, etc..
// TODO: Store player prefs & history to file

const (
	playerCount    int = 4
	cardsPerPlayer int = 5
)

type player struct {
	idx       int
	num       int // Friendly reference
	name      string
	wins      int
	losses    int
	handScore int
	handDesc  string
}

type card struct {
	idx        int    // 0..            // for indexed manipulation
	deckSeq    int    // 40				// Deuces 1st & Aces 13th per suit - Rank in Deck (Suits: 1-13C, 14-26D, 27-39H, 40-52S)  // TODO: DEPRECATE [idx instead]
	suitedName string // "Ace of Spades"
	noSuitName string // "Ace" 			// Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace
	suitPlural string // "Spades"  		// Clubs, Diamonds, Hearts, Spades, Jokers
	suitSingle string // "Spade"		// Club,  Diamond,  Heart,  Spade,  Joker
	rankPip    string // "A" 			// 2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A
	suitPip    string // "♠" 			// ♣, ♦, ♥, ♠
	suitPipUC  string // "U+2660" 		// U+2663, U+2666, U+2665, U+2660
	suitRank   int    // 0-3 			// Suit authority among suits
	rankInSuit int    // 13 			// Rank within Suit 1-13 [2 -> Ace]
	// opRankInSuit int    // 1 		// Rank within Suit when Ace < Two]
	// cardGlyph string // "🂡" 		 // Unicode glyph for card of rank & suit
	// cardUC string // "U+1F0A1"		// Unicode glyph encoding for card of rank & suit
	//  slang 	   linked-list  // 2, Deuce // Key-Value Pairs "GameName" = "SlangTerm"  (dupe keys?)
}

// REVIEW: Add individual cards, uc's, and friendly names to defineGlyphs {IF unicode display of cards provides ANY(?) value to end product.}
// REVIEW: Potentially incorporate image files. ie: ..\assets\images\club.png - need comparable size rank pip images, or large Font {embed font w/suit glyphs & sufficient points [1/72nd"]}
func defineGlyphs() map[string]string { //. Retrieve glyphs by name, and unicode by glyph
	glyphRef := map[string]string{ //TODO: Need to return looked-up value, not table
		"Club":        "♣",
		"Diamond":     "♦",
		"Heart":       "♥",
		"Spade":       "♠️",
		"♣":           "U+2663",
		"♦":           "U+2666",
		"♥":           "U+2665", // Coerce Heart Red  &#x2665;&#xFE0F; - Coerce Black &#x2665;&#xFE0E;
		"♠":           "U+2660",
		"BLACK JOKER": "U+1F0CF",
		"🃏":           "U+1F0CF",
	}
	return glyphRef
}

func getSuitRank() map[string]int { // TODO: Collapse into makeCard
	rankBySuit := map[string]int{
		"Club":    1,
		"Diamond": 2,
		"Heart":   3,
		"Spade":   4,
	}
	return rankBySuit
}

func logDeck(deck []card, format string, yorn bool) int {
	if yorn {
		fmt.Printf("\n -=* Printing the current deck *=-")
		// fmt.Println(deck[i])
		fmt.Println("Cards in deck:", len(deck))
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

func logCard(currentCard card, yorn bool) {
	if yorn {
		fmt.Print("logCard: ", currentCard.suitedName, " - ", currentCard.rankPip, currentCard.suitPip, " [Suit ", currentCard.suitRank, "/4] - Ranked: ",
			currentCard.rankInSuit, " amongst ", currentCard.suitPlural, ", and Card# ", currentCard.deckSeq, " in (Aces high) Ordered-Deck sequence.\n")
	}
}

func makePlayer(playerNum int) player {
	p := player{
		idx:       playerNum,
		num:       playerNum + 1,
		name:      "Anonymous",
		wins:      0,
		losses:    0,
		handScore: 0,
		handDesc:  "Hand score & description pending evaluation?",
	}
	names := []string{"Arthur", "Betty", "Charlie", "Denise", "Eddie", "Fran", "George", "Henrietta"}
	p.name = names[playerNum]
	return p
}

func establishPlayers(playerCount int) []player {
	allPlayers := []player{}
	for i := 0; i < playerCount; i++ { // for each player i
		allPlayers = append(allPlayers, makePlayer(i))
	}
	return allPlayers[:]
}

func makeCard(newCardNum int) card { //? *card
	c := card{deckSeq: newCardNum}
	c.idx = newCardNum - 1
	suitedNames := [53]string{"fillerSuitedName_0",
		"Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs", "Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs", "Ace of Clubs",
		"Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts", "Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts", "Ace of Hearts",
		"Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds", "Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds", "Ace of Diamonds",
		"Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades", "Six of Spades", "Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades", "Ace of Spades",
	}
	c.suitedName = suitedNames[newCardNum]
	c.noSuitName = strings.Fields(c.suitedName)[0]
	c.suitPlural = strings.Fields(c.suitedName)[2]
	c.suitSingle = strings.TrimSuffix(c.suitPlural, "s")
	rankPips := [53]string{"fillerRankPip_0", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	c.rankPip = rankPips[newCardNum]
	c.suitPip = defineGlyphs()[c.suitSingle]
	c.suitPipUC = defineGlyphs()[c.suitPip]
	c.suitRank = getSuitRank()[c.suitSingle]
	rankInSuit := newCardNum
	for rankInSuit > 13 {
		rankInSuit = rankInSuit - 13
	}
	c.rankInSuit = rankInSuit
	// c.opRankInSuit = 1
	// c.cardGlyph = "🂡"
	// c.cardUC = "U+1F0A1"
	return c //?  &c
}

func createDeck() []card {
	accumDeck := []card{}
	for i := 1; i <= 52; i++ {
		accumDeck = append(accumDeck, makeCard(i))
	}
	return accumDeck[:] // returns slice
}

func findCardIndex(theDeck []card, cardInQuestion card) int {
	for idx, thatCard := range theDeck {
		if thatCard == cardInQuestion {
			return idx
		}
	}
	return -1 // no card found
}

func removeCard(theDeck []card, cardToRemove card) []card {
	var idx int = findCardIndex(theDeck, cardToRemove)
	theDeck[idx] = theDeck[len(theDeck)-1] // move to back line
	theDeck = theDeck[:len(theDeck)-1]     // then lop it off - fast
	return theDeck

}

func pickACard(theDeck []card) ([]card, card) {
	cardDrawn := theDeck[rand.IntN(len(theDeck))]
	theDeck = removeCard(theDeck, cardDrawn)
	return theDeck, cardDrawn
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

func main() {
	fmt.Printf("\n -=& Let the games begin.. *=-\n")
	theDeck := createDeck()

	// Define playerHands[0-->playerCount-1] &  Deal #cardsPerPlayer to #playerCount players
	var aHand []card
	var cardShark []player = establishPlayers(playerCount)
	playerHands := make([][]card, playerCount)
	for i := 1; i <= playerCount; i++ { // for each player i
		theDeck, aHand = deal(theDeck, cardsPerPlayer)  // deal * cards into aHand
		playerHands[i-1] = make([]card, cardsPerPlayer) // init player# for cards * x
		for j := 0; j < cardsPerPlayer; j++ {           // for each card_j
			playerHands[i-1][j] = aHand[j] // put card {aHand[#j]} into playerHand#(i-1) slot j
		}
	}

	// Print out dealt hands for verification
	for i := 0; i < playerCount; i++ {
		fmt.Printf("Player %d's Hand (%d Cards): ", i+1, len(aHand))
		fmt.Println("Player Details: ", cardShark[i])
		for j := 0; j < len(playerHands[i]); j++ {
			fmt.Print(playerHands[i][j].suitedName, ", ")
		}
		fmt.Println()
	}

	// Refactor & consolidate funcs that won't be called again.  Also - cap names for pub/share?
	// Cleanup display of state - start considering presentation?
	// Research stat tables to teach Denise & Eddie how to play...
	//* send player# and correlated hand# off to be sorted, scored, and strategy determined[fold,draw#,bet,bluff,raise,etc] @ sortAndScoreHand(playerHands[i])

	logDeck(theDeck, "full", true)
}

// func sortAndScoreHand(hand []card) (sortedHand [][]card, score int, desc string) {

// 	return sortedHand, score, desc
// }
