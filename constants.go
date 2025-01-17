package main

type (
	card struct {
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

	player struct {
		idx       int
		name      string
		hand      []card
		handScore int
		wins      int
		losses    int
	}
)

const (
	numberOfPlayers int = 4
	cardsPerPlayer  int = 5
	rankPips            = "W23456789TJQKA" // W=Wild(Joker)
	suitPips            = "♣♦♥♠"
)
