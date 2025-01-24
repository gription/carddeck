/*
Package: carddeck - Provides a virtual card deck, and the tools most frequently required to manipulate it in a tabletop card game (without jokers?)
*/
package main

import (
	"fmt"
	"time"
)

// * It seems next would be either present cards for discard selection,
// * create the discard mechanism, or start the wager system, or 'ai' discard/wager process?

func main() {
	fmt.Printf("\n --==< Begin: carddeck v0.2.. >==--")

	var u util

	defer timeTrack(time.Now(), "carddeck Complete", u)

	theDeck := createDeck()

	theDeck, players := establishPlayers(theDeck, numberOfPlayers, cardsPerPlayer)

	logDeck(theDeck, "full", true)

	// . Sort, Score & Reveal the players' cards..
	for i := range players {
		fmt.Print(players[i].name, ": ", descHand(players[i].hand), "\n")
		players[i].hand, players[i].handScore, players[i].kickers, players[i].tieBreakPipStr = sortAndScore(players[i].hand)

		fmt.Println("----------------------------------------------------")
	}

	theButton := getRandomPlayer(numberOfPlayers)
	fmt.Println("Player", theButton, "is the first dealer. (Has 'the button', or bets/discards last.)")

	dealFromTheBottom(players, 0)

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	defer timeTrack(time.Now(), "carddeck Complete", u)

	u.logMe("--==< Log Terminated >==--")
	u.logMe("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
}

func timeTrack(start time.Time, name string, u util) {
	u.logMe("--==< Log from func >==--")
	// Sandwich code of concern with the following:  elapsed = interim between
	// defer timeTrack(time.Now(), "comment_for_log_indicating_what_was_tracked")
	elapsed := time.Since(start)
	// u.logMe("%s took %s", name, elapsed)
	// u.logMe("%s took %s", name, elapsed)
	fmt.Printf("%s took %s", name, elapsed)
	u.logMe("nookty mcclockity here")
}
