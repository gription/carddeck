package main

//* Package - Development purposes: Not intended for production release.

import "fmt"

func logCard(c card, yorn bool) {
	if yorn {
		fmt.Print(
			" 🃏", //                                                         //.  Glyph
			// " logCard: #", c.idx, //                                             //.  logCard: #24
			" - ", c.rankPip, c.suitPip, //                                         //. - Q♥
			" - ", c.suitedName, //                                                 //. - Queen of Hearts
			// " - [Suit: ", c.suitRank, "/4]", //                                  //. - [Suit: 3/4]
			// " - Ranked: ", c.rankInSuit, " amongst ", c.suitPlural, //           //. - Ranked: 11 amongst Hearts
			// " - and Card#: ", c.idx, " in (Aces high) Ordered-Deck sequence.\n", //. - and Card#: 24 in (Aces high) Ordered-Deck sequence.
			"\n",
		)
	}
}

func logDeck(deck []card, format string, yorn bool) int {
	if yorn {
		for i := range deck {
			if format == "full" {
				c := deck[i]

				if i%6 == 0 {
					fmt.Print("\n")
				}

				if len(c.suitedName) > 15 {
					fmt.Print(c.idx, "\t", c.suitedName, "\t")
				} else {
					fmt.Print(c.idx, "\t", c.suitedName, "\t\t")
				}
			}
		}

		fmt.Println("\n --==<", len(deck), "Cards remain in deck. >==--")
	}

	return 0
}

func logPlayers(players []player) {
	for i := range numberOfPlayers {
		fmt.Print(players[i].name, ": ", descHand(players[i].hand), "\n")
	}
}
