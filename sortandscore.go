package main

import (
	"fmt"
	"sort"
)

//, seqScore [Sequence Score - Where 5 cards are 'active' & 'kicker' is the highest of those five]
//| 0 = Nothing         V   K
//| 1 = Straight        V   K
//| 2 = Flush           V   K
//| 3 = Straight Flush  V   K
//| 4 = Royal Flush     V   K
//. V = Hand Indicator
//. K = rankInSuit of highest card in sequence/hand
//, setScore [Set Score - If no Sequences apply tally the sets, and track what's left as the 'kicker' - 2-Pair = Anomaly]
//| 0 = High Card   V   K
//| 1 = Pair        V   K
//| 2 = Two Pair    V   K   ? Integrate 2nd_Pair_Val into Kicker string ?
//| 3 = Trips       V   K
//| 4 = Quads       V   K
//| 5 = Full House  V   V
//. V =  int - Hand Indicator
//. K =  slice of 'Kickers' (inactive but tiebreakers)
//* TBD = rankInSuit of top set/pair (+2ND PAIR) {Probably make V slice or struct w/ 2+ values (+1 for 2nd pair}
//? maybe add kickers to player struct .. maybe even 2nd pair?

func sortAndScore(aHand []card) ([]card, int) {
	// var sortedHand []card
	score := 1

	// sort cards {slice of structs} by rankInSuit
	sort.Slice(aHand, func(i, j int) bool {
		return aHand[i].rankInSuit > aHand[j].rankInSuit
	})

	// seqScore: 0=Nothing, 1=Straight, 2=Flush, 3=Straight-Flush, 4=Royal-Flush
	seqScore, highCard := checkForSeqs(aHand)
	fmt.Println("Final seqScore:", seqScore, ": highCard:", highCard)
	fmt.Println("------------------------------")

	if seqScore == 0 {
		// setScore: 0=High Card, 1=Pair, 2=Two Pair, 3=Trips, 4=Quads, 5=Full House [kicker = slice of 'extra' []cards]
		setScore, kicker := checkForSets(aHand)
		fmt.Println("Final setScore:", setScore, ": kicker:", kicker)
		fmt.Println("====================================================")
	}
	// & log hand
	// for i := range aHand {
	// 	fmt.Println(i, aHand[i].suitedName, "- Rank:", aHand[i].rankInSuit)
	// }

	//* NEXT: Refactor to refine scores - seqScore + setScore logic for = score
	return aHand, score
}

func checkForSeqs(aHand []card) (seqScore int, highCard int) {
	// seqScore := 0 // defaults to no sequences
	isStraight := true
	isFlush := false
	// highCard := 0
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

func checkForSets(aHand []card) (int, []card) {
	setScore := 1 // defaults to just the pair..
	pairCount := 0

	setMap := make(map[int]int)
	for _, card := range aHand { // compile k/v pairs reflecting rankInSuit/qty
		setMap[card.rankInSuit]++ // ie: map[7:2 11:1 13:2] = 2 Pairs (Aces & Eights + Queen Kicker)
	}

	handCfg := copyHashMapII(setMap)       // copy reqd as splitKickers
	kickers := splitKickers(aHand, setMap) // was dropping non-kickers from setMap {scope?}

	if (len(handCfg)) == (len(aHand)) {
		return 0, aHand // not even a pair
	}

	for rankInSuit := range handCfg {
		if handCfg[rankInSuit] > 1 {
			pairCount++ // tally pairs
		}
	}

	for _, qty := range handCfg {
		if qty == 4 {
			return 4, kickers // quads
		}

		if qty == 3 {
			for _, qty := range handCfg {
				if qty == 2 {
					return 5, kickers // full house
				}

				setScore = 3 // trips
			}
		}
	}

	if pairCount == 2 {
		setScore = 2 // two pair
	}

	return setScore, kickers // defaults to just the pair..
}

func splitKickers(aHand []card, setMap map[int]int) []card {
	kickers := []card{} // {} = pre-allocate?

	for key, val := range setMap {
		if val != 1 { // delete all but singular cards
			delete(setMap, key)
		}
	}

	for _, card := range aHand { // convert setMap back to cards
		for rank := range setMap {
			if rank == card.rankInSuit {
				kickers = append(kickers, card)
			}
		}
	}

	return kickers
}

func copyHashMapII(originalMap map[int]int) map[int]int {
	newMap := make(map[int]int)

	for key, value := range originalMap {
		newMap[key] = value
	}

	return newMap
}
