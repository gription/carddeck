/*
Package: carddeck - Provides a virtual card deck, and the tools most frequently required to manipulate it in a tabletop card game (without jokers?)
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("\n --==< Begin: carddeck v0.2.. >==--")

	theDeck := createDeck()
	theDeck, players := establishPlayers(theDeck, numberOfPlayers, cardsPerPlayer)
	logDeck(theDeck, "full", true)
	// logPlayers(players)
	// fmt.Print("One lucky card for the dealer:")
	// logCard(players[0].hand[1], true)

	for i := range players {
		players[i].hand, players[i].handScore = sortAndScore(players[i].hand)
	}

	fmt.Println("Sorted...")
	logPlayers(players)
}

func sortAndScore(aHand []card) ([]card, int) {
	// var sortedHand []card
	score := 1

	// sort cards {slice of structs} by rankInSuit
	sort.Slice(aHand, func(i, j int) bool {
		return aHand[i].rankInSuit > aHand[j].rankInSuit
	})

	setScore := checkForSets(aHand)

	// & log hand
	for i := range aHand {
		fmt.Println(i, aHand[i].suitedName, "- Rank:", aHand[i].rankInSuit)
	}

	fmt.Print("Final setScore:", setScore, "\n\n")

	seqScore := checkForSeqs(aHand)
	fmt.Println("seqScore: :", seqScore)
	//* NEXT: implement seqScore + setScore logic for = handScore
	//* NEXT: Refactor to refine scores (set/pair PipValues & Kickers)
	//* NEXT: figure out 'kickers' to address ties..
	//? yrst assert certain hands?
	return aHand, score
}

func checkForSeqs(aHand []card) int {
	//, seqScore
	//| 0 = Nothing         V   K
	//| 1 = Straight        V   K
	//| 2 = Flush           V   K
	//| 3 = Straight Flush  V   K
	//| 4 = Royal Flush     V   K
	//. V = rankInSuit of highest card in sequence/hand
	//. K = String of 'Kickers'
	//* Now test for seqScore (straight, flush, straight-flush, royal-flush)
	seqScore := len(aHand)

	return seqScore
}

func checkForSets(aHand []card) int {
	//, setScore
	//| 0 = High Card   V   K
	//| 1 = Pair        V   K
	//| 2 = Two Pair    V   K   Integrate 2nd_Pair_Val into Kicker string
	//| 3 = Trips       V   K
	//| 4 = Quads       V   K
	//| 5 = Full House  V   V
	//. V = int - rankInSuit of top set/pair
	//. K = String of 'Kickers'
	setScore := 1 // has at least a Pair
	pairCount := 0

	setTally := make(map[int]int) // k/v pairs reflecting card/qty
	for _, card := range aHand {
		setTally[card.rankInSuit]++
	}

	if (len(setTally)) == (len(aHand)) {
		return 0 // not even a pair
	}

	// sort setTally map into 'keys' slice by descending order
	keys := make([]int, 0, len(setTally))

	for key := range setTally {
		if setTally[key] > 1 {
			pairCount++ // if two pair
		}

		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return setTally[keys[i]] > setTally[keys[j]]
	})

	for k := range keys {
		if setTally[k] == 4 {
			return 4 // quads
		}

		if setTally[k] == 3 {
			for _, k := range keys {
				if setTally[k] == 2 {
					return 5 // full house
				}

				setScore = 3 // trips
			}
		}
	}

	if pairCount == 2 {
		setScore = 2 // two pair
	}

	return setScore // defaults to just the pair..
}

// TODO: splitKickers
// func splitKickers(aHand []card, num int) ([]card, []card) { // divide 'num' of 'focus' cards from leftover 'kicker' cards
//     return focus, kickers
// }
