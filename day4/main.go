package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	CardNumber  int
	WinningCard []int
	YourCard    []int
	Score       int
	Copies      int
}

type CardMap map[int]Card

func main() {
	Question1()
	Question2()
}

// Create a function to read the data
func ReadData() []Card {
	file, _ := os.Open("input.txt")
	defer file.Close()
	s := bufio.NewScanner(file)

	Cards := []Card{}
	for s.Scan() {
		card := Card{}
		split1 := strings.Split(s.Text(), "|")                       // left winners, right yours
		split2 := strings.Split(split1[0], ":")                      // left card number, right winners
		yourCard := strings.Split(strings.TrimSpace(split1[1]), " ") // left card number, right yours
		winners := strings.Split(strings.TrimSpace(split2[1]), " ")  // left card number, right winners

		for _, v := range yourCard {
			v, _ := strconv.Atoi(strings.TrimSpace(v))
			if v == 0 {
				continue
			}
			card.YourCard = append(card.YourCard, v)
		}

		for _, v := range winners {
			v, _ := strconv.Atoi(strings.TrimSpace(v))
			if v == 0 {
				continue
			}
			card.WinningCard = append(card.WinningCard, v)
		}

		// regex to extract the card number
		re := regexp.MustCompile("[0-9]+")
		card.CardNumber, _ = strconv.Atoi(re.FindString(split2[0]))

		Cards = append(Cards, card)
	}

	return Cards
}

// Create a function to read the data
func ReadDataMap() CardMap {
	file, _ := os.Open("input.txt")
	defer file.Close()
	s := bufio.NewScanner(file)

	cardMap := make(CardMap)
	for s.Scan() {
		card := Card{}
		split1 := strings.Split(s.Text(), "|")                       // left winners, right yours
		split2 := strings.Split(split1[0], ":")                      // left card number, right winners
		yourCard := strings.Split(strings.TrimSpace(split1[1]), " ") // left card number, right yours
		winners := strings.Split(strings.TrimSpace(split2[1]), " ")  // left card number, right winners

		for _, v := range yourCard {
			v, _ := strconv.Atoi(strings.TrimSpace(v))
			if v == 0 {
				continue
			}

			card.YourCard = append(card.YourCard, v)
		}

		for _, v := range winners {
			v, _ := strconv.Atoi(strings.TrimSpace(v))
			if v == 0 {
				continue
			}
			card.WinningCard = append(card.WinningCard, v)
		}

		// regex to extract the card number
		re := regexp.MustCompile("[0-9]+")
		card.CardNumber, _ = strconv.Atoi(re.FindString(split2[0]))
		card.Copies = 1
		cardMap[card.CardNumber] = card

	}

	return cardMap
}

// Calculate the score of a card
// 1 match = 1 point and then doubles for each additional match
// assumes no duplicates in your card
func Question1() {
	cards := ReadData()

	total := 0
	for _, card := range cards {
		for _, v := range card.YourCard {
			if slices.Contains(card.WinningCard, v) {
				if card.Score == 0 {
					card.Score = 1
				} else {
					card.Score = card.Score * 2
				}
			}
		}
		total += card.Score
	}

	fmt.Println("Q1: ", total)
}

func Question2() {
	cards := ReadDataMap()

	for card := 1; card < len(cards); card++ {
		for copies := 0; copies < cards[card].Copies; copies++ {
			copyCounter := 0
			for _, v := range cards[card].YourCard {
				if slices.Contains(cards[card].WinningCard, v) {
					copyCounter++
					tmp := cards[card+copyCounter]
					tmp.Copies = tmp.Copies + 1
					cards[card+copyCounter] = tmp
				}
			}
		}
	}

	// Count all the copies
	total := 0
	for _, v := range cards {
		total += v.Copies
	}

	fmt.Println("Q2: ", total)
}
