package main

type player struct {
	idx       int
	num       int // Friendly reference
	name      string
	wins      int
	losses    int
	handScore int
	handDesc  string
}

func makePlayer(playerNum int) player {
	p := player{
		idx:       playerNum,
		num:       playerNum + 1,
		name:      "Anonymous",
		wins:      0,
		losses:    0,
		handScore: 0,
		handDesc:  "Hand score & description pending evaluation?",
	}
	names := []string{"Arthur", "Betty", "Charlie", "Denise", "Eddie", "Fran", "George", "Henrietta"}
	p.name = names[playerNum]
	return p
}

func establishPlayers(playerCount int) []player {
	allPlayers := []player{}
	for i := 0; i < playerCount; i++ { // for each player i
		allPlayers = append(allPlayers, makePlayer(i))
	}
	return allPlayers[:]
}
