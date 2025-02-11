package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

//? should player numbers from 0 or 1 ?

func establishPlayers(theDeck []card, numberOfPlayers int, cardsPerPlayer int) ([]card, []player) {
	// var wg sync.WaitGroup
	playerNames := []string{"Arthur", "Betty", "Charlie", "Denise", "Eddie", "Fran", "George", "Henrietta"}

	players := []player{} //? players := []player{}  diff?  either works -  BUT?  *** pre-allocate ??
	//                            //? var players []player      Happier wif pre-alloc from top choice?
	shoe := theDeck

	var aHand []card

	for i := range numberOfPlayers {
		// wg.Add
		shoe, aHand = deal(shoe, cardsPerPlayer)

		p := player{
			idx:            i,
			name:           playerNames[i],
			hand:           aHand,
			kickers:        aHand,
			handScore:      0,
			tieBreakPipStr: "",
			accountBalance: startingBalance,
			currentWager:   0,
			currentInPot:   0,
			debtToPot:      0,
			strategy:       1,
			wins:           0,
			losses:         0,
		}
		players = append(players, p)
	}

	return shoe, players
}

func getRandomPlayer(playerCount int) int {
	// generator := rand.New(rand.NewSource(time.Now().UnixNano())) // Create new random generator
	// theButton := generator.Intn(playerCount)
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(playerCount)))
	if err != nil {
		panic(err)
	}

	theButton := nBig.Int64()

	return int(theButton)
}

func getNextPlayer(currentPlayer int, activePlayers []player) int {
	var nextPlayer int
	if currentPlayer == len(activePlayers)-1 {
		nextPlayer = 0
	} else {
		nextPlayer = currentPlayer + 1
	}

	return nextPlayer
}

func logPlayers(players []player) {
	for i := range players {
		fmt.Println("Player", i, "-", players[i].name, "[ Balance:", players[i].accountBalance, "] Holding:", descHand(players[i].hand))
	}
}
