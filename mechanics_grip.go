package main

import (
	"fmt"
)

//* TESTS: This is repo of hands for testing.

var grip = []card{
	{1, "Two of Clubs", "Two", "Clubs", "Club", "2", "♣", "U+2663", 1, 1},
	{2, "Three of Clubs", "Three", "Clubs", "Club", "3", "♣", "U+2663", 1, 2},
	{3, "Four of Clubs", "Four", "Clubs", "Club", "4", "♣", "U+2663", 1, 3},
	{4, "Five of Clubs", "Five", "Clubs", "Club", "5", "♣", "U+2663", 1, 4},
	{5, "Six of Clubs", "Six", "Clubs", "Club", "6", "♣", "U+2663", 1, 5},
	{6, "Seven of Clubs", "Seven", "Clubs", "Club", "7", "♣", "U+2663", 1, 6},
	{7, "Eight of Clubs", "Eight", "Clubs", "Club", "8", "♣", "U+2663", 1, 7},
	{8, "Nine of Clubs", "Nine", "Clubs", "Club", "9", "♣", "U+2663", 1, 8},
	{9, "Ten of Clubs", "Ten", "Clubs", "Club", "T", "♣", "U+2663", 1, 9},
	{10, "Jack of Clubs", "Jack", "Clubs", "Club", "J", "♣", "U+2663", 1, 10},
	{11, "Queen of Clubs", "Queen", "Clubs", "Club", "Q", "♣", "U+2663", 1, 11},
	{12, "King of Clubs", "King", "Clubs", "Club", "K", "♣", "U+2663", 1, 12},
	{13, "Ace of Clubs", "Ace", "Clubs", "Club", "A", "♣", "U+2663", 1, 13},
	{14, "Two of Hearts", "Two", "Hearts", "Heart", "2", "♥", "U+2665", 3, 1},
	{15, "Three of Hearts", "Three", "Hearts", "Heart", "3", "♥", "U+2665", 3, 2},
	{16, "Four of Hearts", "Four", "Hearts", "Heart", "4", "♥", "U+2665", 3, 3},
	{17, "Five of Hearts", "Five", "Hearts", "Heart", "5", "♥", "U+2665", 3, 4},
	{18, "Six of Hearts", "Six", "Hearts", "Heart", "6", "♥", "U+2665", 3, 5},
	{19, "Seven of Hearts", "Seven", "Hearts", "Heart", "7", "♥", "U+2665", 3, 6},
	{20, "Eight of Hearts", "Eight", "Hearts", "Heart", "8", "♥", "U+2665", 3, 7},
	{21, "Nine of Hearts", "Nine", "Hearts", "Heart", "9", "♥", "U+2665", 3, 8},
	{22, "Ten of Hearts", "Ten", "Hearts", "Heart", "T", "♥", "U+2665", 3, 9},
	{23, "Jack of Hearts", "Jack", "Hearts", "Heart", "J", "♥", "U+2665", 3, 10},
	{24, "Queen of Hearts", "Queen", "Hearts", "Heart", "Q", "♥", "U+2665", 3, 11},
	{25, "King of Hearts", "King", "Hearts", "Heart", "K", "♥", "U+2665", 3, 12},
	{26, "Ace of Hearts", "Ace", "Hearts", "Heart", "A", "♥", "U+2665", 3, 13},
	{27, "Two of Diamonds", "Two", "Diamonds", "Diamond", "2", "♦", "U+2666", 2, 1},
	{28, "Three of Diamonds", "Three", "Diamonds", "Diamond", "3", "♦", "U+2666", 2, 2},
	{29, "Four of Diamonds", "Four", "Diamonds", "Diamond", "4", "♦", "U+2666", 2, 3},
	{30, "Five of Diamonds", "Five", "Diamonds", "Diamond", "5", "♦", "U+2666", 2, 4},
	{31, "Six of Diamonds", "Six", "Diamonds", "Diamond", "6", "♦", "U+2666", 2, 5},
	{32, "Seven of Diamonds", "Seven", "Diamonds", "Diamond", "7", "♦", "U+2666", 2, 6},
	{33, "Eight of Diamonds", "Eight", "Diamonds", "Diamond", "8", "♦", "U+2666", 2, 7},
	{34, "Nine of Diamonds", "Nine", "Diamonds", "Diamond", "9", "♦", "U+2666", 2, 8},
	{35, "Ten of Diamonds", "Ten", "Diamonds", "Diamond", "T", "♦", "U+2666", 2, 9},
	{36, "Jack of Diamonds", "Jack", "Diamonds", "Diamond", "J", "♦", "U+2666", 2, 10},
	{37, "Queen of Diamonds", "Queen", "Diamonds", "Diamond", "Q", "♦", "U+2666", 2, 11},
	{38, "King of Diamonds", "King", "Diamonds", "Diamond", "K", "♦", "U+2666", 2, 12},
	{39, "Ace of Diamonds", "Ace", "Diamonds", "Diamond", "A", "♦", "U+2666", 2, 13},
	{40, "Two of Spades", "Two", "Spades", "Spade", "2", "♠", "U+2660️", 4, 1},
	{41, "Three of Spades", "Three", "Spades", "Spade", "3", "♠️", "U+2660", 4, 2},
	{42, "Four of Spades", "Four", "Spades", "Spade", "4", "♠️", "U+2660", 4, 3},
	{43, "Five of Spades", "Five", "Spades", "Spade", "5", "♠️", "U+2660", 4, 4},
	{44, "Six of Spades", "Six", "Spades", "Spade", "6", "♠️", "U+2660", 4, 5},
	{45, "Seven of Spades", "Seven", "Spades", "Spade", "7", "♠️", "U+2660", 4, 6},
	{46, "Eight of Spades", "Eight", "Spades", "Spade", "8", "♠️", "U+2660", 4, 7},
	{47, "Nine of Spades", "Nine", "Spades", "Spade", "9", "♠️", "U+2660", 4, 8},
	{48, "Ten of Spades", "Ten", "Spades", "Spade", "T", "♠️", "U+2660", 4, 9},
	{49, "Jack of Spades", "Jack", "Spades", "Spade", "J", "♠️", "U+2660", 4, 10},
	{50, "Queen of Spades", "Queen", "Spades", "Spade", "Q", "♠️", "U+2660", 4, 11},
	{51, "King of Spades", "King", "Spades", "Spade", "K", "♠️", "U+2660", 4, 12},
	{52, "Ace of Spades", "Ace", "Spades", "Spade", "A", "♠️", "U+2660", 4, 13},
}

func gimmeFlush() ([]card, int) { // AQ762 - all Diamonds
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: Flush >==----------------------------------------------")

	cardsDealt := []card{
		{27, "Two of Diamonds", "Two", "Diamonds", "Diamond", "2", "♦", "U+2666", 2, 1},
		{31, "Six of Diamonds", "Six", "Diamonds", "Diamond", "6", "♦", "U+2666", 2, 5},
		{39, "Ace of Diamonds", "Ace", "Diamonds", "Diamond", "A", "♦", "U+2666", 2, 13},
		{32, "Seven of Diamonds", "Seven", "Diamonds", "Diamond", "7", "♦", "U+2666", 2, 6},
		{37, "Queen of Diamonds", "Queen", "Diamonds", "Diamond", "Q", "♦", "U+2666", 2, 11},
	}

	return sortAndScore(cardsDealt)
}

func gimmeStraight() ([]card, int) { // 65432
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: Straight >==----------------------------------------------")

	cardsDealt := []card{
		{3, "Four of Clubs", "Four", "Clubs", "Club", "4", "♣", "U+2663", 1, 3},
		{44, "Six of Spades", "Six", "Spades", "Spade", "6", "♠️", "U+2660", 4, 5},
		{40, "Two of Spades", "Two", "Spades", "Spade", "2", "♠", "U+2660️", 4, 1},
		{4, "Five of Clubs", "Five", "Clubs", "Club", "5", "♣", "U+2663", 1, 4},
		{28, "Three of Diamonds", "Three", "Diamonds", "Diamond", "3", "♦", "U+2666", 2, 2},
	}

	return sortAndScore(cardsDealt)
}

func gimmeStraightFlush() ([]card, int) { // 98765 - all Hearts
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: Straight Flush >==----------------------------------------------")

	cardsDealt := []card{
		{20, "Eight of Hearts", "Eight", "Hearts", "Heart", "8", "♥", "U+2665", 3, 7},
		{21, "Nine of Hearts", "Nine", "Hearts", "Heart", "9", "♥", "U+2665", 3, 8},
		{19, "Seven of Hearts", "Seven", "Hearts", "Heart", "7", "♥", "U+2665", 3, 6},
		{18, "Six of Hearts", "Six", "Hearts", "Heart", "6", "♥", "U+2665", 3, 5},
		{17, "Five of Hearts", "Five", "Hearts", "Heart", "5", "♥", "U+2665", 3, 4},
	}

	return sortAndScore(cardsDealt)
}

func gimmeRoyalFlush() ([]card, int) { // AKQJT - all Clubs
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: Royal Flush >==----------------------------------------------")

	cardsDealt := []card{
		{9, "Ten of Clubs", "Ten", "Clubs", "Club", "T", "♣", "U+2663", 1, 9},
		{10, "Jack of Clubs", "Jack", "Clubs", "Club", "J", "♣", "U+2663", 1, 10},
		{11, "Queen of Clubs", "Queen", "Clubs", "Club", "Q", "♣", "U+2663", 1, 11},
		{12, "King of Clubs", "King", "Clubs", "Club", "K", "♣", "U+2663", 1, 12},
		{13, "Ace of Clubs", "Ace", "Clubs", "Club", "A", "♣", "U+2663", 1, 13},
	}

	return sortAndScore(cardsDealt)
}

func gimmeHighCard() ([]card, int) { // AKT63
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: High Card (Ace High) >==----------------------------------------------")

	cardsDealt := []card{
		{52, "Ace of Spades", "Ace", "Spades", "Spade", "A", "♠️", "U+2660", 4, 13},
		{38, "King of Diamonds", "King", "Diamonds", "Diamond", "K", "♦", "U+2666", 2, 12},
		{22, "Ten of Hearts", "Ten", "Hearts", "Heart", "T", "♥", "U+2665", 3, 9},
		{5, "Six of Clubs", "Six", "Clubs", "Club", "6", "♣", "U+2663", 1, 5},
		{41, "Three of Spades", "Three", "Spades", "Spade", "3", "♠️", "U+2660", 4, 2},
	}

	return sortAndScore(cardsDealt)
}

func gimmeOnePair() ([]card, int) { // JJT97
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: One Pair >==----------------------------------------------")

	cardsDealt := []card{
		{23, "Jack of Hearts", "Jack", "Hearts", "Heart", "J", "♥", "U+2665", 3, 10},
		{49, "Jack of Spades", "Jack", "Spades", "Spade", "J", "♠️", "U+2660", 4, 10},
		{22, "Ten of Hearts", "Ten", "Hearts", "Heart", "T", "♥", "U+2665", 3, 9},
		{8, "Nine of Clubs", "Nine", "Clubs", "Club", "9", "♣", "U+2663", 1, 8},
		{32, "Seven of Diamonds", "Seven", "Diamonds", "Diamond", "7", "♦", "U+2666", 2, 6},
	}

	return sortAndScore(cardsDealt)
}

func gimmeTwoPair() ([]card, int) { // AA88Q
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: Two Pair >==----------------------------------------------")

	cardsDealt := []card{
		{26, "Ace of Hearts", "Ace", "Hearts", "Heart", "A", "♥", "U+2665", 3, 13},
		{52, "Ace of Spades", "Ace", "Spades", "Spade", "A", "♠️", "U+2660", 4, 13},
		{33, "Eight of Diamonds", "Eight", "Diamonds", "Diamond", "8", "♦", "U+2666", 2, 7},
		{7, "Eight of Clubs", "Eight", "Clubs", "Club", "8", "♣", "U+2663", 1, 7},
		{24, "Queen of Hearts", "Queen", "Hearts", "Heart", "Q", "♥", "U+2665", 3, 11},
	}

	return sortAndScore(cardsDealt)
}

func gimmeTrips() ([]card, int) { // 333K5
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: Trips >==----------------------------------------------")

	cardsDealt := []card{
		{41, "Three of Spades", "Three", "Spades", "Spade", "3", "♠️", "U+2660", 4, 2},
		{15, "Three of Hearts", "Three", "Hearts", "Heart", "3", "♥", "U+2665", 3, 2},
		{2, "Three of Clubs", "Three", "Clubs", "Club", "3", "♣", "U+2663", 1, 2},
		{51, "King of Spades", "King", "Spades", "Spade", "K", "♠️", "U+2660", 4, 12},
		{17, "Five of Hearts", "Five", "Hearts", "Heart", "5", "♥", "U+2665", 3, 4},
	}

	return sortAndScore(cardsDealt)
}

func gimmeQuads() ([]card, int) { // 4444T
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: Quads >==----------------------------------------------")

	cardsDealt := []card{
		{16, "Four of Hearts", "Four", "Hearts", "Heart", "4", "♥", "U+2665", 3, 3},
		{3, "Four of Clubs", "Four", "Clubs", "Club", "4", "♣", "U+2663", 1, 3},
		{29, "Four of Diamonds", "Four", "Diamonds", "Diamond", "4", "♦", "U+2666", 2, 3},
		{42, "Four of Spades", "Four", "Spades", "Spade", "4", "♠️", "U+2660", 4, 3},
		{9, "Ten of Clubs", "Ten", "Clubs", "Club", "T", "♣", "U+2663", 1, 9},
	}

	return sortAndScore(cardsDealt)
}

func gimmeFullHouse() ([]card, int) { // AA888
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: Full House >==----------------------------------------------")

	cardsDealt := []card{
		{13, "Ace of Clubs", "Ace", "Clubs", "Club", "A", "♣", "U+2663", 1, 13},
		{52, "Ace of Spades", "Ace", "Spades", "Spade", "A", "♠️", "U+2660", 4, 13},
		{46, "Eight of Spades", "Eight", "Spades", "Spade", "8", "♠️", "U+2660", 4, 7},
		{7, "Eight of Clubs", "Eight", "Clubs", "Club", "8", "♣", "U+2663", 1, 7},
		{33, "Eight of Diamonds", "Eight", "Diamonds", "Diamond", "8", "♦", "U+2666", 2, 7},
	}

	return sortAndScore(cardsDealt)
}

func gimmeTrash() ([]card, int) { // can't win 75432
	fmt.Println("--==< FROM THE BOTTOM: Dealing Player 0: Trash >==----------------------------------------------")

	cardsDealt := []card{
		{19, "Seven of Hearts", "Seven", "Hearts", "Heart", "7", "♥", "U+2665", 3, 6},
		{43, "Five of Spades", "Five", "Spades", "Spade", "5", "♠️", "U+2660", 4, 4},
		{3, "Four of Clubs", "Four", "Clubs", "Club", "4", "♣", "U+2663", 1, 3},
		{28, "Three of Diamonds", "Three", "Diamonds", "Diamond", "3", "♦", "U+2666", 2, 2},
		{14, "Two of Hearts", "Two", "Hearts", "Heart", "2", "♥", "U+2665", 3, 1},
	}

	return sortAndScore(cardsDealt)
}

// {1 Two of Clubs Two Clubs Club 2 ♣ U+2663 1 1}
// {2 Three of Clubs Three Clubs Club 3 ♣ U+2663 1 2}
// {3 Four of Clubs Four Clubs Club 4 ♣ U+2663 1 3}
// {4 Five of Clubs Five Clubs Club 5 ♣ U+2663 1 4}
// {5 Six of Clubs Six Clubs Club 6 ♣ U+2663 1 5}
// {6 Seven of Clubs Seven Clubs Club 7 ♣ U+2663 1 6}
// {7 Eight of Clubs Eight Clubs Club 8 ♣ U+2663 1 7}
// {8 Nine of Clubs Nine Clubs Club 9 ♣ U+2663 1 8}
// {9 Ten of Clubs Ten Clubs Club T ♣ U+2663 1 9}
// {10 Jack of Clubs Jack Clubs Club J ♣ U+2663 1 10}
// {11 Queen of Clubs Queen Clubs Club Q ♣ U+2663 1 11}
// {12 King of Clubs King Clubs Club K ♣ U+2663 1 12}
// {13 Ace of Clubs Ace Clubs Club A ♣ U+2663 1 13}
// {14 Two of Hearts Two Hearts Heart 2 ♥ U+2665 3 1}
// {15 Three of Hearts Three Hearts Heart 3 ♥ U+2665 3 2}
// {16 Four of Hearts Four Hearts Heart 4 ♥ U+2665 3 3}
// {17 Five of Hearts Five Hearts Heart 5 ♥ U+2665 3 4}
// {18 Six of Hearts Six Hearts Heart 6 ♥ U+2665 3 5}
// {19 Seven of Hearts Seven Hearts Heart 7 ♥ U+2665 3 6}
// {20 Eight of Hearts Eight Hearts Heart 8 ♥ U+2665 3 7}
// {21 Nine of Hearts Nine Hearts Heart 9 ♥ U+2665 3 8}
// {22 Ten of Hearts Ten Hearts Heart T ♥ U+2665 3 9}
// {23 Jack of Hearts Jack Hearts Heart J ♥ U+2665 3 10}
// {24 Queen of Hearts Queen Hearts Heart Q ♥ U+2665 3 11}
// {25 King of Hearts King Hearts Heart K ♥ U+2665 3 12}
// {26 Ace of Hearts Ace Hearts Heart A ♥ U+2665 3 13}
// {27 Two of Diamonds Two Diamonds Diamond 2 ♦ U+2666 2 1}
// {28 Three of Diamonds Three Diamonds Diamond 3 ♦ U+2666 2 2}
// {29 Four of Diamonds Four Diamonds Diamond 4 ♦ U+2666 2 3}
// {30 Five of Diamonds Five Diamonds Diamond 5 ♦ U+2666 2 4}
// {31 Six of Diamonds Six Diamonds Diamond 6 ♦ U+2666 2 5}
// {32 Seven of Diamonds Seven Diamonds Diamond 7 ♦ U+2666 2 6}
// {33 Eight of Diamonds Eight Diamonds Diamond 8 ♦ U+2666 2 7}
// {34 Nine of Diamonds Nine Diamonds Diamond 9 ♦ U+2666 2 8}
// {35 Ten of Diamonds Ten Diamonds Diamond T ♦ U+2666 2 9}
// {36 Jack of Diamonds Jack Diamonds Diamond J ♦ U+2666 2 10}
// {37 Queen of Diamonds Queen Diamonds Diamond Q ♦ U+2666 2 11}
// {38 King of Diamonds King Diamonds Diamond K ♦ U+2666 2 12}
// {39 Ace of Diamonds Ace Diamonds Diamond A ♦ U+2666 2 13}
// {40 Two of Spades Two Spades Spade 2 ♠ U+2660️ 4 1}
// {41 Three of Spades Three Spades Spade 3 ♠️ U+2660 4 2}
// {42 Four of Spades Four Spades Spade 4 ♠️ U+2660 4 3}
// {43 Five of Spades Five Spades Spade 5 ♠️ U+2660 4 4}
// {44 Six of Spades Six Spades Spade 6 ♠️ U+2660 4 5}
// {45 Seven of Spades Seven Spades Spade 7 ♠️ U+2660 4 6}
// {46 Eight of Spades Eight Spades Spade 8 ♠️ U+2660 4 7}
// {47 Nine of Spades Nine Spades Spade 9 ♠️ U+2660 4 8}
// {48 Ten of Spades Ten Spades Spade T ♠️ U+2660 4 9}
// {49 Jack of Spades Jack Spades Spade J ♠️ U+2660 4 10}
// {50 Queen of Spades Queen Spades Spade Q ♠️ U+2660 4 11}
// {51 King of Spades King Spades Spade K ♠️ U+2660 4 12}
// {52 Ace of Spades Ace Spades Spade A ♠️ U+2660 4 13}
