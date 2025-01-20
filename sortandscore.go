package main

import (
	"fmt"
	"sort"
)

func sortAndScore(aHand []card) ([]card, int) {
	// var sortedHand []card
	score := 1

	// sort cards {slice of structs} by rankInSuit
	sort.Slice(aHand, func(i, j int) bool {
		return aHand[i].rankInSuit > aHand[j].rankInSuit
	})

	seqScore, highCard := checkForSeqs(aHand) // seqScore: 0=Nothing, 1=Straight, 2=Flush, 3=Straight-Flush, 4=Royal-Flush
	fmt.Println("Final seqScore:", seqScore, "highCard:", highCard)
	fmt.Println("------------------------------")

	if seqScore == 0 {
		setScore, kicker := checkForSets(aHand) // setScore: 0=High Card, 1=Pair, 2=Two Pair, 3=Trips, 4=Quads, 5=Full House [kicker = slice of 'extra' []card ]
		fmt.Println("Final setScore:", setScore, "kicker:", kicker)
		fmt.Println("====================================================")
	}
	// & log hand
	// for i := range aHand {
	// 	fmt.Println(i, aHand[i].suitedName, "- Rank:", aHand[i].rankInSuit)
	// }

	//* NEXT: implement seqScore + setScore logic for = handScore
	//* NEXT: Refactor to refine scores (set/pair PipValues & Kickers to address ties..)
	return aHand, score
}

func checkForSeqs(aHand []card) (int, int) {
	seqScore := 0 // defaults to no sequences
	isStraight := true
	isFlush := false
	highCard := 0
	straightAssessment := make(map[int]int)

	var priorCard int

	for cardNum, card := range aHand {
		straightAssessment[cardNum] = card.rankInSuit // map[0:13 1:11 2:6 3:5 4:1] no, map[0:13 1:12 2:11 3:10 4:9] yes

		if cardNum > 0 {
			if card.rankInSuit != priorCard-1 {
				isStraight = false
			}
		}

		priorCard = card.rankInSuit
		highCard = straightAssessment[0]
	}

	seqArray := make(map[int]int) // k/v pairs reflecting rankInSuit/suit - map[0:1 1:1 2:1 3:1 4:1] flush
	for cardNum, card := range aHand {
		seqArray[cardNum] = card.suitRank
	}

	suitSum := 0

	for cardNum := range seqArray { // count suits for flushMetric
		if len(seqArray) > 1 {
			suitSum = suitSum + seqArray[cardNum]
		}
	}

	if suitSum/len(aHand) == seqArray[0] { // flushMetric
		isFlush = true
		seqScore = 2
	}

	if isStraight {
		seqScore = 1

		if isFlush {
			seqScore = 3

			if highCard == 13 {
				seqScore = 4
			}
		}
	}

	fmt.Println("isStraight:", isStraight, "isFlush:", isFlush)

	return seqScore, highCard
}

// ** kickers: must take applied cards out of aHand, sort the remains by rankInHand if >1 and return sorted 'minihand'
// * intent: build string of 'kickers' then call splitKickers with aHand and kickers string - extracting cards from aHand per kickers values & returning minihand of kicker cards
func checkForSets(aHand []card) (int, []card) {
	setScore := 1 // defaults to just the pair..
	pairCount := 0

	var kickers string

	setTally := make(map[int]int)
	for _, card := range aHand { // compile k/v pairs reflecting rankInSuit/qty map[7:2 11:1 13:2]
		setTally[card.rankInSuit]++
	}

	if (len(setTally)) == (len(aHand)) {
		kickerCards := splitKickers(aHand, kickers)

		return 0, kickerCards //* not even a pair
	}

	for rankInSuit := range setTally {
		if setTally[rankInSuit] > 1 {
			pairCount++ // tally pairs
		}
	}

	for _, qty := range setTally {
		if qty == 4 {
			kickerCards := splitKickers(aHand, kickers)

			return 4, kickerCards //* quads
		}

		if qty == 3 {
			for _, qty := range setTally {
				if qty == 2 {
					kickerCards := splitKickers(aHand, kickers)

					return 5, kickerCards //* full house
				}

				setScore = 3 // trips
			}
		}
	}

	if pairCount == 2 {
		setScore = 2 // two pair
	}

	kickerCards := splitKickers(aHand, kickers)

	return setScore, kickerCards //* defaults to just the pair..
}

// TODO: splitKickers
func splitKickers(aHand []card, kickers string) []card { // divide 'num' of 'focus' cards from leftover 'kicker' cards
	var kickerCards []card

	for i := range aHand {
		for j := range kickers {
			if j == aHand[i].rankInSuit { //?rankPip
				kickerCards = append(kickerCards, aHand[i])
			}
		}
	}
	// .cards: func removeCard(theDeck []card, cardToRemove card) []card {  ...	return modifiedDeck }

	return kickerCards
}

//, seqScore
//| 0 = Nothing         V   K
//| 1 = Straight        V   K
//| 2 = Flush           V   K
//| 3 = Straight Flush  V   K
//| 4 = Royal Flush     V   K
//. V = rankInSuit of highest card in sequence/hand
//. K = String of 'Kickers'
//, setScore
//| 0 = High Card   V   K
//| 1 = Pair        V   K
//| 2 = Two Pair    V   K   Integrate 2nd_Pair_Val into Kicker string
//| 3 = Trips       V   K
//| 4 = Quads       V   K
//| 5 = Full House  V   V
//. V = int - rankInSuit of top set/pair
//. K = String of 'Kickers'
