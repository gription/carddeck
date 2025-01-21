package main

import (
	"fmt"
	"sort"
	"strconv"
)

//%                             compiledScore
//, seqScore [Sequence Score - Where 5 cards are 'active' & 'kicker' is the highest of those five]
//| 0 = Nothing         0   K       x
//| 1 = Straight        1   K       4
//| 2 = Flush           2   K       5
//| 3 = Straight Flush  3   K       8
//| 4 = Royal Flush     4   K       9
//. V = Hand Indicator
//. K = rankInSuit of highest card in sequence/hand
//, setScore [Set Score - If no Sequences apply tally the sets, and track what's left as the 'kicker' - 2-Pair = Anomaly]
//| 0 = High Card   0   K           0
//| 1 = Pair        1   K           1
//| 2 = Two Pair    2   K           2
//| 3 = Trips       3   K           3
//| 4 = Quads       4   K           7
//| 5 = Full House  5   V           6
//. V =  int - Hand Indicator
//. K =  slice of 'Kickers' (inactive but tiebreakers)
//? maybe add kickers to player struct .. maybe even 2nd pair?

func sortAndScore(aHand []card) ([]card, int, []card, string) {
	sort.Slice(aHand, func(i, j int) bool { // sort cards descending by rankInSuit {slice of structs}
		return aHand[i].rankInSuit > aHand[j].rankInSuit
	})

	seqScore, highCard := checkForSeqs(aHand)
	// seqScoreDesc := getScoreDesc(seqScore)

	highSeqPips := cardStrToCardPip(strconv.Itoa(highCard))

	// fmt.Println("seqScore:", seqScore, "------------ SeqScore: [", seqScore, "]", seqScoreDesc, ": High:", highSeqPips)

	setScore, kickers, setPips := checkForSets(aHand)
	// setScoreDesc := getScoreDesc(setScore)
	// fmt.Println("setScore:", setScore, "------------------ SetScore:", setScoreDesc, "Of:", setPips, "- kickers: [", len(kickers), "]", descHand(kickers))

	handScore := seqScore
	tieBreakPipStr := highSeqPips

	if handScore == 0 { // if not a sequence, reset score & relevant pips (kickers unchanged)
		handScore = setScore
		tieBreakPipStr = setPips
	}

	fmt.Println("--==< handScore:", handScore, getScoreDesc(handScore), "( Ext:", tieBreakPipStr, ") - Kickers: [", len(kickers), "]", descHand(kickers))

	// & log hand
	// for i := range aHand {
	// 	fmt.Println(i, aHand[i].suitedName, "- Rank:", aHand[i].rankInSuit)
	// }
	return aHand, handScore, kickers, tieBreakPipStr
}

func checkForSeqs(aHand []card) (seqScore int, highCard int) {
	isStraight := true
	isFlush := true
	seqScore = 5 // default = flush

	rankInSuitMap := make(map[int]int)

	var priorCard int

	for cardNum, card := range aHand { // check for not straight
		rankInSuitMap[cardNum] = card.rankInSuit // map[0:13 1:11 2:6 3:5 4:1] no, map[0:13 1:12 2:11 3:10 4:9] yes

		if cardNum > 0 {
			if card.rankInSuit != priorCard-1 {
				isStraight = false
			}
		}

		priorCard = card.rankInSuit
		highCard = rankInSuitMap[0]
	}

	suitRankMap := make(map[int]int)
	for cardNum, card := range aHand {
		suitRankMap[cardNum] = card.suitRank
	}

	for cardNum := 1; cardNum < len(aHand); cardNum++ { // check for not flush - map[0:1 1:1 2:1 3:1 4:1]
		if suitRankMap[cardNum] != suitRankMap[0] {
			isFlush = false
			seqScore = 0

			break
		}
	}

	if isStraight { // straight
		seqScore = 4

		if isFlush { // straight-flush
			seqScore = 8

			if highCard == 13 { // royal-flush
				seqScore = 9
			}
		}
	}

	return seqScore, highCard
}

func checkForSets(aHand []card) (int, []card, string) {
	setScore := 1 // defaults to just the pair..
	pairCount := 0

	origSetMap := make(map[int]int)
	for _, card := range aHand { // compile k/v pairs reflecting rankInSuit/qty
		origSetMap[card.rankInSuit]++ // ie: map[7:2 11:1 13:2] = 2 Pairs (Aces & Eights + Queen Kicker)
	}

	fullSetMap := copyHashMapII(origSetMap) // copy reqd as splitKickers dropping non-kickers from origSetMap {scope?}

	setPips := extractSets(fullSetMap)

	kickers := splitKickers(aHand, origSetMap)

	if (len(fullSetMap)) == (len(aHand)) {
		return 0, aHand, setPips // not even a pair
	}

	for rankInSuit := range fullSetMap {
		if fullSetMap[rankInSuit] > 1 {
			pairCount++ // tally pairs
		}
	}

	for _, qty := range fullSetMap {
		if qty == 4 {
			return 7, kickers, setPips // quads
		}

		if qty == 3 {
			for _, qty := range fullSetMap {
				if qty == 2 {
					return 6, kickers, setPips // full house
				}

				setScore = 3 // trips
			}
		}
	}

	if pairCount == 2 {
		setScore = 2 // two pair
	}

	return setScore, kickers, setPips // defaults to just the pair..
}

func splitKickers(aHand []card, setMap map[int]int) []card {
	kickers := []card{}

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

func extractSets(setMap map[int]int) string {
	var setPips string

	for cardVal, cardQty := range setMap {
		if cardQty > 1 {
			stringVer := strconv.Itoa(cardVal)
			stringVer = cardStrToCardPip(stringVer) // TODO: Sort to assure highest is always first in string
			setPips = setPips + stringVer
		}
	}

	if setPips == "" {
		setPips = "N/A"
	}

	return setPips
}
