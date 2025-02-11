package main

type (
	card struct {
		idx        int    // 0..            // Deuces 1st & Aces 13th per suit - Rank in Deck (Suits: 1-13C, 14-26D, 27-39H, 40-52S)
		suitedName string // "Ace of Spades"
		noSuitName string // "Ace" 			// Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace
		suitPlural string // "Spades"  		// Clubs, Diamonds, Hearts, Spades, Jokers
		suitSingle string // "Spade"		// Club,  Diamond,  Heart,  Spade,  Joker
		rankPip    string // "A" 			// 2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A
		suitPip    string // "♠" 			// ♣, ♦, ♥, ♠
		suitPipUC  string // "U+2660" 		// U+2663, U+2666, U+2665, U+2660 - &clubs; &diams; &hearts; &spades;
		suitRank   int    // 0-3 			// Suit authority among suits
		rankInSuit int    // 13 			// Rank within Suit 1-13 [2 -> Ace]
		// opRankInSuit int    // 1 		// Rank within Suit when Ace < Two]
		// cardGlyph string // "🂡" 		 // Unicode glyph for card of rank & suit //= https://en.wikipedia.org/wiki/Playing_cards_in_Unicode
		// cardUC string // "U+1F0A1"		// Unicode glyph encoding for card of rank & suit
		//  slang 	   linked-list  // 2, Deuce // Key-Value Pairs "GameName" = "SlangTerm"  (dupe keys?)
	}

	player struct {
		idx            int
		name           string
		hand           []card
		kickers        []card
		tieBreakPipStr string
		handScore      int
		accountBalance int
		currentWager   int
		currentInPot   int
		debtToPot      int
		strategy       int
		wins           int
		losses         int
	}

	util struct{}
)

const (
	numberOfPlayers int = 5
	cardsPerPlayer  int = 5
	startingBalance int = 1000
	betMinimum          = 10
	rankPips            = "W23456789TJQKA" // W=Wild(Joker)
	suitPips            = "♣♦♥♠"
)

func cardStrToCardPip(cardStr string) string { // TODO: Use rankPips more elegantly to resolve? [conv to int, ref pos in rankPips?]
	var newCardStr string

	if cardStr == "1" {
		newCardStr = "2"
	}

	if cardStr == "2" {
		newCardStr = "3"
	}

	if cardStr == "3" {
		newCardStr = "4"
	}

	if cardStr == "4" {
		newCardStr = "5"
	}

	if cardStr == "5" {
		newCardStr = "6"
	}

	if cardStr == "6" {
		newCardStr = "7"
	}

	if cardStr == "7" {
		newCardStr = "8"
	}

	if cardStr == "8" {
		newCardStr = "9"
	}

	if cardStr == "9" {
		newCardStr = "T"
	}

	if cardStr == "10" {
		newCardStr = "J"
	}

	if cardStr == "11" {
		newCardStr = "Q"
	}

	if cardStr == "12" {
		newCardStr = "K"
	}

	if cardStr == "13" {
		newCardStr = "A"
	}

	return newCardStr
}

func getScoreDesc(scoreInt int) string {
	var scoreDesc string

	if scoreInt == 0 {
		scoreDesc = "High-Card"
	}

	if scoreInt == 4 {
		scoreDesc = "Straight"
	}

	if scoreInt == 5 {
		scoreDesc = "Flush"
	}

	if scoreInt == 8 {
		scoreDesc = "Straight-Flush"
	}

	if scoreInt == 9 {
		scoreDesc = "Royal-Flush"
	}

	if scoreInt == 1 {
		scoreDesc = "One-Pair"
	}

	if scoreInt == 2 {
		scoreDesc = "Two-Pair"
	}

	if scoreInt == 3 {
		scoreDesc = "Trips"
	}

	if scoreInt == 7 {
		scoreDesc = "Quads"
	}

	if scoreInt == 6 {
		scoreDesc = "Full-House"
	}

	return scoreDesc
}
