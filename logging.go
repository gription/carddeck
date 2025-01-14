package main

//* Package - Development purposes: Not intended for production release.

import "fmt"

func logCard(c card, yorn bool) {
	if yorn {
		fmt.Print(
			"logCard: #",
			c.idx, " - ",
			c.suitedName, " - ",
			c.rankPip, c.suitPip, " [Suit ",
			c.suitRank, "/4] - Ranked: ",
			c.rankInSuit, " amongst ",
			c.suitPlural, ", and Card# ",
			c.idx, " in (Aces high) Ordered-Deck sequence.\n",
		)
	}
}

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
					fmt.Print(c.idx, "\t", c.suitedName, "\t")
				} else {
					fmt.Print(c.idx, "\t", c.suitedName, "\t\t")
				}
			}
		}
	}
	fmt.Print("\n")
	return 0
}
