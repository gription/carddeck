package main

import (
	"fmt"
)

const (
	playerCount    int = 4
	cardsPerPlayer int = 5
)

func main() {
	fmt.Printf("\n -=& Let the games begin.. *=-\n")
	theDeck := createDeck()

	// Define playerHands[0-->playerCount-1] &  Deal #cardsPerPlayer to #playerCount players
	var aHand []card
	var cardShark []player = establishPlayers(playerCount)
	playerHands := make([][]card, playerCount)
	for i := 1; i <= playerCount; i++ { //, for each player i
		theDeck, aHand = deal(theDeck, cardsPerPlayer)  // deal * cards into aHand
		playerHands[i-1] = make([]card, cardsPerPlayer) // init player# for cards * x
		for j := 0; j < cardsPerPlayer; j++ {           // for each card_j
			playerHands[i-1][j] = aHand[j] // put card {aHand[#j]} into playerHand#(i-1) slot j
		}

		playerState := []any{cardShark[i-1], playerHands[i-1]} //? any was interface{} (interface{} - allows us to store elements of different types in slice)
		fmt.Println("Player - playerState[0]:", playerState[0])
		fmt.Println("Hand - playerState[1]:", playerState[1])
		fmt.Print("--------\n")

		// fmt.Println("Player:", thisPlayerState[0].([]player))
		// fmt.Println("Hand:", thisPlayerState[1].([][]card))

	}

	// Print out dealt hands for verification
	for i := 0; i < playerCount; i++ {

		fmt.Printf("Player %d's Hand (%d Cards): ", i+1, len(aHand))
		fmt.Println("Player Details: ", cardShark[i])
		for j := 0; j < len(playerHands[i]); j++ {
			fmt.Print(playerHands[i][j].suitedName, ", ")
		}
		fmt.Println()
		// fmt.Println("pplayerState: playerState[1])", playerState[1])
	}

	//* send player# and correlated hand# off to be sorted, scored, and strategy determined[fold,draw#,bet,bluff,raise,etc] @ sortAndScoreHand(playerHands[i])
	// thisPlayerState := []interface{}{cardShark, playerHands}
	//` fmt.Println("\n\n-------- thisPlayerState: ", thisPlayerState)
	//` panic: interface conversion: interface {} is []main.player, not main.player
	// fmt.Println("Player:", thisPlayerState[0].([]player))
	// fmt.Println("Hand:", thisPlayerState[1].([][]card))

	// fmt.Println("Hand:", thisPlayerState[0])

	//&
	// var playerState state
	// for i :=0; i < playerCount; i++ {
	// 	playerState[i] = establishPlayerState(cardShark[i], playerHands[i])

	// }

	logDeck(theDeck, "full", true)
}

//    thisPlayerState := []interface{}{cardShark, playerHands}
//     fmt.Println("Player:", thisPlayerState[0].(Player).Name)
//     fmt.Println("Hand:", thisPlayerState[1].([]Card))

// func sortAndScoreHand(hand []card) (sortedHand [][]card, score int, desc string) {

// 	return sortedHand, score, desc
// }
