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

//=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

//& Unicode block: Playing Cards (U+1F0A0-1F0FF)
// Also, the knight, a specific fool and twenty-one generic trump cards (based on the Bourgeois Tarot) are available but not shown.

//|  ♤      U+2664         White Spade Suit
//|  ♡      U+2661         White Heart Suit
//|  ♢      U+2662         White Diamond Suit
//|  ♧      U+2667         White Club Suit

//   ♠      U+2660         Black Spade Suit
//   ♥      U+2665         Black Heart Suit
//   ♦      U+2666         Black Diamond Suit
//   ♣      U+2663         Black Club Suit

//, Back and Three Jokers (red, black and white)
//.  🂠       U+1F0A0       Playing Card Back
//+  🂿       U+1F0BF       Red Joker
//   🃏︎       U+1F0CF       Black Joker
//|  🃟       U+1F0DF       White Joker

//. U+1F0A1	        U+1F0B1	        U+1F0C1	            U+1F0D1
//|   🂡                🂱               🃁                   🃑
//~ Ace of Spades	Ace of Hearts	Ace of Diamonds	    Ace of Clubs
//. U+1F0A2	        U+1F0B2	        U+1F0C2	            U+1F0D2
//|   🂢                🂲               🃂                   🃒
//~ Two of Spades	Two of Hearts	Two of Diamonds	    Two of Clubs
//. U+1F0A3	        U+1F0B3	        U+1F0C3	            U+1F0D3
//|   🂣                🂳               🃃                   🃓
//~ Three of Spades	Three of Hearts	Three of Diamonds	Three of Clubs
//. U+1F0A4	        U+1F0B4	        U+1F0C4	            U+1F0D4
//|   🂤                🂴               🃄                   🃔
//~ Four of Spades	Four of Hearts	Four of Diamonds	Four of Clubs
//. U+1F0A5	        U+1F0B5	        U+1F0C5	            U+1F0D5
//|   🂥                🂵               🃅                   🃕
//~ Five of Spades	Five of Hearts	Five of Diamonds	Five of Clubs
//. U+1F0A6	        U+1F0B6	        U+1F0C6	            U+1F0D6
//|   🂦                🂶               🃆                   🃖
//~ Six of Spades	Six of Hearts	Six of Diamonds	Six of Clubs
//. U+1F0A7	        U+1F0B7	        U+1F0C7	            U+1F0D7
//|   🂧                🂷               🃇                   🃗
//~ Seven of Spades	Seven of Hearts	Seven of Diamonds	Seven of Clubs
//. U+1F0A8	        U+1F0B8	        U+1F0C8	            U+1F0D8
//|   🂨                🂸               🃈                   🃘
//~ Eight of Spades	Eight of Hearts	Eight of Diamonds	Eight of Clubs
//. U+1F0A9	        U+1F0B9	        U+1F0C9	            U+1F0D9
//|   🂩                🂹               🃉                   🃙
//~ Nine of Spades	Nine of Hearts	Nine of Diamonds	Nine of Clubs
//. U+1F0AA	        U+1F0BA	        U+1F0CA	            U+1F0DA
//|   🂪                🂺               🃊                   🃚
//~ Ten of Spades	Ten of Hearts	Ten of Diamonds	Ten of Clubs
//. U+1F0AB	        U+1F0BB	        U+1F0CB	            U+1F0DB
//|   🂫                🂻               🃋                   🃛
//~ Jack of Spades	Jack of Hearts	Jack of Diamonds	Jack of Clubs
//. U+1F0AD	        U+1F0BD	        U+1F0CD	            U+1F0DD
//|   🂭                🂽               🃍                   🃝
//~ Queen of Spades	Queen of Hearts	Queen of Diamonds	Queen of Clubs
//. U+1F0AE	        U+1F0BE	        U+1F0CE	            U+1F0DE
//|   🂮                🂾               🃎                   🃞
//~ King of Spades	King of Hearts	King of Diamonds	King of Clubs
