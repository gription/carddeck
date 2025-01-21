package main

//? should player numbers from 0 or 1 ?

func establishPlayers(theDeck []card, numberOfPlayers int, cardsPerPlayer int) ([]card, []player) {
	playerNames := []string{"Arthur", "Betty", "Charlie", "Denise", "Eddie", "Fran", "George", "Henrietta"}

	players := []player{} //? players := []player{}  diff?  either works -  BUT?  *** pre-allocate ??
	//                            //? var players []player      Happier wif pre-alloc from top choice?
	shoe := theDeck

	var aHand []card

	for i := range numberOfPlayers {
		shoe, aHand = deal(shoe, cardsPerPlayer)

		p := player{
			idx:       i,
			name:      playerNames[i],
			hand:      aHand,
			handScore: 0,
			wins:      0,
			losses:    0,
		}
		players = append(players, p)
	}

	return shoe, players
}
