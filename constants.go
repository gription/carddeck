package main

const (
	numberOfPlayers int = 4
	cardsPerPlayer  int = 5
	rankPips            = "W23456789TJQKA" // W=Wild(Joker)
	suitPips            = "♣♦♥♠"
)

var (
	names       = []string{"Arthur", "Betty", "Charlie", "Denise", "Eddie", "Fran", "George", "Henrietta"}
	rankBySuit  = map[string]int{"Club": 1, "Diamond": 2, "Heart": 3, "Spade": 4, "Joker": 5}
	suitedNames = [52]string{
		"Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs", "Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs", "Ace of Clubs",
		"Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts", "Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts", "Ace of Hearts",
		"Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds", "Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds", "Ace of Diamonds",
		"Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades", "Six of Spades", "Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades", "Ace of Spades",
	}
)

var glyphRef = map[string]string{ // Maps are reference types, so they are always passed by reference. You don't need a pointer.
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

type card struct {
	idx        int    // 0..            // Deuces 1st & Aces 13th per suit - Rank in Deck (Suits: 1-13C, 14-26D, 27-39H, 40-52S)  // TODO: DEPRECATE [idx instead]
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

type player struct {
	idx       int
	name      string
	hand      []card
	handScore int
	handDesc  string
	wins      int
	losses    int
}
