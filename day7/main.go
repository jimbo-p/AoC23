package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards    string
	Bid      int
	CardMap  map[int]int
	CardVals int
	HandType int
	HandRank int
}

func main() {
	Question1()
	Question2()
}

// Create a function to read the data
func ReadData() []Hand {
	file, _ := os.Open("input.txt")
	defer file.Close()
	s := bufio.NewScanner(file)

	hands := []Hand{}
	for s.Scan() {
		line := s.Text()
		hand := strings.Split(line, " ")
		bid, _ := strconv.Atoi(hand[1])
		hands = append(hands, Hand{Cards: hand[0], Bid: bid})
	}

	return hands
}

func Question1() {
	hands := ReadData()
	hands = EnhanceHand(hands)
	hands = HandType(hands)
	hands = SortHands(hands)
	hands = RankHands(hands)

	totalWinnings := 0
	for _, hand := range hands {
		totalWinnings += hand.HandRank
	}
	fmt.Println("Q1: ", totalWinnings)
}

func EnhanceHand(hands []Hand) []Hand {
	for i, hand := range hands {
		cardMap := map[int]int{}
		for j, card := range hand.Cards {
			cardVal := 0
			switch card {
			case 'T':
				cardVal = 10
			case 'J':
				cardVal = 11
			case 'Q':
				cardVal = 12
			case 'K':
				cardVal = 13
			case 'A':
				cardVal = 14
			default:
				cardVal, _ = strconv.Atoi(string(card))
			}
			cardMap[cardVal]++
			hands[i].CardVals += int(math.Pow(float64(10), float64(10-j*2)) * float64(cardVal))
		}
		hands[i].CardMap = cardMap
	}

	return hands
}

func HandType(hands []Hand) []Hand {
	for i, hand := range hands {
		// Five of kind
		if len(hand.CardMap) == 1 {
			hands[i].HandType = 6
			continue
		}
		// Four of kind or Full House
		if len(hand.CardMap) == 2 {
			for _, v := range hand.CardMap {
				if v == 4 {
					hands[i].HandType = 5
					break
				} else {
					hands[i].HandType = 4
				}
			}
			continue
		}
		// Three of kind or Two Pair
		if len(hand.CardMap) == 3 {
			for _, v := range hand.CardMap {
				if v == 3 {
					hands[i].HandType = 3
					break
				} else {
					hands[i].HandType = 2
				}
			}
			continue
		}
		// One Pair
		if len(hand.CardMap) == 4 {
			hands[i].HandType = 1
			continue
		}
		// High Card
		hands[i].HandType = 0
	}
	return hands
}

func SortHands(hands []Hand) []Hand {
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].HandType > hands[j].HandType {
			return true
		} else if hands[i].HandType < hands[j].HandType {
			return false
		} else {
			return hands[i].CardVals > hands[j].CardVals
		}
	})

	return hands
}

func RankHands(hands []Hand) []Hand {
	for i, hand := range hands {
		hands[i].HandRank = (len(hands) - i) * hand.Bid
	}

	return hands
}

func Question2() {
	hands := ReadData()
	hands = EnhanceHand2(hands)
	hands = HandType2(hands)
	hands = SortHands2(hands)
	hands = RankHands2(hands)

	totalWinnings := 0
	for _, hand := range hands {
		totalWinnings += hand.HandRank
	}
	fmt.Println("Q1: ", totalWinnings)
}

func EnhanceHand2(hands []Hand) []Hand {
	for i, hand := range hands {
		cardMap := map[int]int{}
		for j, card := range hand.Cards {
			cardVal := 0
			switch card {
			case 'T':
				cardVal = 10
			case 'J':
				cardVal = 1
			case 'Q':
				cardVal = 12
			case 'K':
				cardVal = 13
			case 'A':
				cardVal = 14
			default:
				cardVal, _ = strconv.Atoi(string(card))
			}
			cardMap[cardVal]++
			hands[i].CardVals += int(math.Pow(float64(10), float64(10-j*2)) * float64(cardVal))
		}
		hands[i].CardMap = cardMap
	}

	return hands
}

func HandType2(hands []Hand) []Hand {
	for i, hand := range hands {
		jokers := hand.CardMap[1]

		// Five of kind
		if len(hand.CardMap) == 1 {
			hands[i].HandType = 6
			continue
		}
		// Four of kind or Full House
		if len(hand.CardMap) == 2 {
			if jokers > 0 {
				hands[i].HandType = 6
				continue
			}
			for _, v := range hand.CardMap {
				if v == 4 {
					hands[i].HandType = 5
					break
				} else {
					hands[i].HandType = 4
				}
			}
			continue
		}
		// Three of kind or Two Pair
		if len(hand.CardMap) == 3 {
			for _, v := range hand.CardMap {
				if v == 3 {
					if jokers > 0 {
						hands[i].HandType = 5
						break
					} else {
						hands[i].HandType = 3
						break
					}
				} else {
					if jokers == 1 {
						hands[i].HandType = 4
					} else if jokers == 2 {
						hands[i].HandType = 5
					} else {
						hands[i].HandType = 2
					}
				}
			}
			continue
		}
		// One Pair
		if len(hand.CardMap) == 4 {
			if jokers > 0 {
				hands[i].HandType = 3
			} else {
				hands[i].HandType = 1
			}
			continue
		}
		// High Card
		if jokers > 0 {
			hands[i].HandType = 1
			continue
		}
		hands[i].HandType = 0
	}
	return hands
}

func SortHands2(hands []Hand) []Hand {
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].HandType > hands[j].HandType {
			return true
		} else if hands[i].HandType < hands[j].HandType {
			return false
		} else {
			return hands[i].CardVals > hands[j].CardVals
		}
	})

	return hands
}

func RankHands2(hands []Hand) []Hand {
	for i, hand := range hands {
		hands[i].HandRank = (len(hands) - i) * hand.Bid
	}

	return hands
}
