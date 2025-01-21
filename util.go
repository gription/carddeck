package main

func copyHashMapII(originalMap map[int]int) map[int]int {
	newMap := make(map[int]int)

	for key, value := range originalMap {
		newMap[key] = value
	}

	return newMap
}

func cardStrToCardPip(cardStr string) string { // TODO: Use rankPips more elegantly to resolve? [conv to int, ref pos in rankPips?]
	var newCardStr string

	if cardStr == "1" {
		newCardStr = "2"
	}

	if cardStr == "2" {
		newCardStr = "3"
	}

	if cardStr == "3" {
		newCardStr = "4"
	}

	if cardStr == "4" {
		newCardStr = "5"
	}

	if cardStr == "5" {
		newCardStr = "6"
	}

	if cardStr == "6" {
		newCardStr = "7"
	}

	if cardStr == "7" {
		newCardStr = "8"
	}

	if cardStr == "8" {
		newCardStr = "9"
	}

	if cardStr == "9" {
		newCardStr = "T"
	}

	if cardStr == "10" {
		newCardStr = "J"
	}

	if cardStr == "11" {
		newCardStr = "Q"
	}

	if cardStr == "12" {
		newCardStr = "K"
	}

	if cardStr == "13" {
		newCardStr = "A"
	}

	return newCardStr
}

func getScoreDesc(scoreInt int) string {
	var scoreDesc string

	if scoreInt == 0 {
		scoreDesc = "High-Card"
	}

	if scoreInt == 4 {
		scoreDesc = "Straight"
	}

	if scoreInt == 5 {
		scoreDesc = "Flush"
	}

	if scoreInt == 8 {
		scoreDesc = "Straight-Flush"
	}

	if scoreInt == 9 {
		scoreDesc = "Royal-Flush"
	}

	if scoreInt == 1 {
		scoreDesc = "One-Pair"
	}

	if scoreInt == 2 {
		scoreDesc = "Two-Pair"
	}

	if scoreInt == 3 {
		scoreDesc = "Trips"
	}

	if scoreInt == 7 {
		scoreDesc = "Quads"
	}

	if scoreInt == 6 {
		scoreDesc = "Full-House"
	}

	return scoreDesc
}
