package main

//& Package - Development purposes: Not intended for production release.

import (
	"fmt"
	"log"
	"os"
	"time"
)

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
		fmt.Println("====================================================")
	}

	return 0
}

func dealFromTheBottom(players []player, pNum int) { //, DEAL FROM THE BOTTOM... (insensitive to deck contents)
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeFlush()
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeStraight()
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeStraightFlush()
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeRoyalFlush()
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeHighCard()
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeOnePair()
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeTwoPair()
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeTrips()
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeQuads()
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeFullHouse()
	// players[pNum].hand, players[pNum].handScore, players[pNum].kickers, players[pNum].tieBreakPipStr = gimmeTrash()
	//. lucky card (how to address a single player card)
	fmt.Print("--==< And one lucky card for the dealer:")
	logCard(players[pNum].hand[getRandCardNum(len(players[pNum].hand))], true)
}

func (u *util) logMe(msg string) {
	logFile, err := os.OpenFile("log/carddeck.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		log.Println("Error opening application log: ", err)
		fmt.Println("Error opening application log: ", err)
	}

	logit := log.New(logFile, "carddeck - ", log.LstdFlags|log.Lmicroseconds)
	logit.Println(msg)

	if err := logFile.Close(); err != nil {
		logit.Printf("Error closing log file: %s", err)
		fmt.Printf("Error closing log file: %s %T\n", err, u) // REVIEW: ref'd 'u' only to alleviate error
		logit.Fatal(err)
	}
}

func timeTrack(startTime time.Time, msgToLog string, u util) {
	// time.Sleep(2 * time.Second) //, delay
	combinedMsg := "Timed Event: " + msgToLog + " - Elapsed Time: " + (time.Since(startTime)).String() + " :._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._."
	fmt.Println(combinedMsg)
	u.logMe(combinedMsg)
}
