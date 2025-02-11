/*
Package: carddeck - Provides a virtual card deck, and the tools most frequently required to manipulate it in a tabletop card game (without jokers?)
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()

	fmt.Printf("\n --==< Begin: carddeck v0.2.. >==--")

	var u util

	//& Prep deck & players
	theDeck := createDeck()
	theDeck, players := establishPlayers(theDeck, numberOfPlayers, cardsPerPlayer)
	logDeck(theDeck, "full", true)

	//& Sort, Score & Reveal the players' cards..
	for i := range players {
		fmt.Print(players[i].name, ": ", descHand(players[i].hand), "\n")
		players[i].hand, players[i].handScore, players[i].kickers, players[i].tieBreakPipStr = sortAndScore(players[i].hand)

		fmt.Println("----------------------------------------------------")
	}

	dealFromTheBottom(players, 0) //% cheat to p#

	//& Determine thebutton & init wager
	theButton := getRandomPlayer(numberOfPlayers)
	currentPlayer := theButton
	whoBetting := getNextPlayer(currentPlayer, players)

	fmt.Println("Player", theButton, "is on 'the button' and the bet's to player", whoBetting)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	var pot int

	players, pot = getWagers(players, whoBetting, pot)

	fmt.Println("Pot Total: $", pot)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	logPlayers(players)

	//& Close out
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	timeTrack(startTime, "[EOF] App init to exit", u)
}

func getWagers(players []player, whoBetting int, pot int) ([]player, int) {
	currentPlayers := players
	playerBetting := whoBetting
	potAccum := pot
	proxToButton := 0

	if len(players) == 1 {
		fmt.Println(proxToButton)
	}

	goRoundTheTable(players, playerBetting, pot)

	return currentPlayers, potAccum
}

func goRoundTheTable(players []player, whoBetting int, pot int) ([]player, int) {
	currentPlayers := players
	playerBetting := whoBetting
	potAccum := pot

	proxToButton := 1

	var wager, betToYou int

	for range players {
		var foldStatus, checkStatus bool

		if playerBetting == 0 {
			currentPlayers = interactiveBet(players)
			wager = currentPlayers[0].currentWager
			// TODO: Check foldStatus for 'nextHand'

			if wager == 0 {
				checkStatus = true
			}
		} else {
			currentPlayers[playerBetting], wager, foldStatus, checkStatus = playerTurn(currentPlayers[playerBetting], betToYou, proxToButton, len(currentPlayers), potAccum)
		}

		if foldStatus {
			// TODO: Remove player from currentPlayers (fold/lose)
			proxToButton--
		}

		betToYou = currentPlayers[playerBetting].currentWager
		potAccum = potAccum + wager

		fmt.Print("goingRoundTable: >>>>> ",
			players[playerBetting].name,
			" Bal: $", players[playerBetting].accountBalance,
			" Accumulated POT: $", potAccum,
			" (of which ", players[playerBetting].currentInPot, " is theirs) ",
			"(Folding? ", foldStatus,
			") Checking? ", checkStatus,
			"\n",
			"*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*\n")

		proxToButton++

		playerBetting = getNextPlayer(playerBetting, players)
	}

	return currentPlayers, potAccum
}

func interactiveBet(players []player) []player {
	// var wager int
	//+ PLAYER INPUT: Commented out for dev speed
	wager := 246
	// if _, err := fmt.Scanln(&wager); err != nil {
	// 	fmt.Println("Invalid input. Please enter an integer:", err)
	// }

	wager, err := strconv.Atoi(strconv.Itoa(wager)) // Convert input to integer
	// wager, err := strconv.Atoi(strings.TrimSpace(fmt.Sprint(wager)))
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer.")

		return players
	}

	if wager < 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")

		return players
	}

	players[0].currentWager = wager
	players[0].accountBalance = players[0].accountBalance - wager
	players[0].currentInPot = players[0].currentInPot + wager

	fmt.Print(":::::::::::::::::::::: PLAYER BETS: You currently have: $ ", players[0].accountBalance, " Whatcha wanna bet?: $ ", wager, "\n")

	return players
}

func playerTurn(currentPlayer player, betToYou int, proxToButton int, numPlayersLeft int, pot int) (modifiedPlayer player, wager int, foldStatus bool, checkStatus bool) {
	// TODO: for each player determine WAGER & FOLDSTATUS based on hand, betToYou, proxToButton and numPlayersLeft
	//? herein lies the players logic
	thisPlayer := currentPlayer
	numPlayersAhead := numPlayersLeft - proxToButton
	numPlayersInPot := numPlayersLeft - numPlayersAhead - 1
	confidence := 1
	maxBet := thisPlayer.accountBalance
	pName := currentPlayer.name

	var avgAmountIn, minimumBet int

	if betToYou > betMinimum {
		minimumBet = betToYou
	} else {
		minimumBet = betMinimum
	}

	if pot != 0 {
		avgAmountIn = pot / numPlayersInPot // accuracy sufficient for calc
	}

	if betToYou > maxBet { // REVIEW: IF betToYou exceeds balance and want to stay (~borrow from pot?) - leads to target shooting minigame? ;o)
		thisPlayer.debtToPot = thisPlayer.debtToPot + (betToYou - maxBet)
	}

	// REVIEW: strategy needs to be float? range 1.0 to 2.0
	// TODO: determine foldStatus [if handScore < XXX AND numPlayersLeft > YYY AND wager > % of accountBalance AND currentWager < % accountBalance AND riskAssessment < 5 ???] // THEN WAGER = 0

	wager = confidence * minimumBet

	if wager == minimumBet {
		checkStatus = true
	} else {
		if wager == 0 {
			foldStatus = true
		}
	}

	thisPlayer.currentWager = wager
	thisPlayer.accountBalance = thisPlayer.accountBalance - wager
	thisPlayer.currentInPot = thisPlayer.currentInPot + wager

	fmt.Print("playerTurn +++++++++++ ",
		pName, " (@btn+", proxToButton,
		") [POT STATUS: $", pot,
		" w/ ", numPlayersInPot, " invested",
		" {Avg~ $: ", avgAmountIn, "", " ea}] ",
		" with: ", numPlayersAhead, " Ahead ",
		"- Bets(Min =", betToYou, "): ", wager,
		"\n")

	return thisPlayer, wager, foldStatus, checkStatus
}

// player
// 		hand           []card
// 		kickers        []card
// 		tieBreakPipStr string
// 		handScore      int
// 		accountBalance int
// 		currentWager   int
//      currentInPot   int
// 		strategy       int
