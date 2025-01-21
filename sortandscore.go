package main

import (
	"fmt"
	"sort"
	"strconv"
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
//? maybe add kickers to player struct .. maybe even 2nd pair?

func sortAndScore(aHand []card) ([]card, int) {
	score := 1

	// sort cards descending by rankInSuit {slice of structs}
	sort.Slice(aHand, func(i, j int) bool {
		return aHand[i].rankInSuit > aHand[j].rankInSuit
	})

	seqScore, highCard := checkForSeqs(aHand)
	seqScoreDesc := getSeqScoreDesc(seqScore)
	fmt.Println("SeqScore: [", seqScore, "]", seqScoreDesc, ": High:", cardStrToCardPip(strconv.Itoa(highCard)))

	if seqScore == 0 {
		setScore, kicker, setPips := checkForSets(aHand)
		setScoreDesc := getSetScoreDesc(setScore)
		fmt.Println("SetScore:", setScoreDesc, "Of:", setPips, "- Kicker: [", len(kicker), "]", descHand(kicker))
	}

	// & log hand
	for i := range aHand {
		fmt.Println(i, aHand[i].suitedName, "- Rank:", aHand[i].rankInSuit)
	}
	//* PENDING: Refactor to compile/refine scores - seqScore + setScore logic for = score
	return aHand, score
}

func checkForSeqs(aHand []card) (seqScore int, highCard int) {
	isStraight := true
	isFlush := true
	seqScore = 2

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

	if isStraight {
		seqScore = 1

		if isFlush {
			seqScore = 3

			if highCard == 13 {
				seqScore = 4
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
			return 4, kickers, setPips // quads
		}

		if qty == 3 {
			for _, qty := range fullSetMap {
				if qty == 2 {
					return 5, kickers, setPips // full house
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
