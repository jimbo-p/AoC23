package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	ReadData()
	Question1()
	Question2()
}

// Create struct to hold the data
type Game struct {
	Id   int
	Sets []Set
}

type Set struct {
	Red   int
	Blue  int
	Green int
}

// Create a function to read the data
func ReadData() []Game {
	file, _ := os.Open("input.txt")
	defer file.Close()
	re := regexp.MustCompile("[0-9]+")

	s := bufio.NewScanner(file)
	Games := make([]Game, 0)
	g_id := 0
	for s.Scan() {
		g_id += 1
		text := strings.Split(s.Text(), ":")
		sets := strings.Split(text[1], ";")
		g := Game{}
		for _, set := range sets {
			colors := strings.Split(set, ",")
			set := Set{}
			for _, color := range colors {
				if strings.Contains(color, "red") {
					r, _ := strconv.ParseInt(strings.Join(re.FindAllString(color, -1), ""), 10, 64)
					set.Red = int(r)
				} else if strings.Contains(color, "blue") {
					b, _ := strconv.ParseInt(strings.Join(re.FindAllString(color, -1), ""), 10, 64)
					set.Blue = int(b)
				} else if strings.Contains(color, "green") {
					gn, _ := strconv.ParseInt(strings.Join(re.FindAllString(color, -1), ""), 10, 64)
					set.Green = int(gn)
				}
			}
			g.Sets = append(g.Sets, set)
		}
		g.Id = g_id
		Games = append(Games, g)
	}
	return Games
}

func Question1() {
	// Bag in question
	r, g, b := 12, 13, 14
	games := ReadData()

	// Loop through each game and see if it's possible
	sum := 0
	for _, game := range games {
		game_possible := true
		for _, set := range game.Sets {
			if set.Red > r || set.Blue > b || set.Green > g {
				game_possible = false
			}
		}
		if game_possible {
			sum += game.Id
		}
	}

	fmt.Println("Q1: ", sum)
}

func Question2() {
	// Loop through each game and to find minimal number of cubes needed to play
	games := ReadData()

	sum := 0
	for _, game := range games {
		set_min := Set{}
		for _, set := range game.Sets {
			if set_min.Red < set.Red {
				set_min.Red = set.Red
			}

			if set_min.Blue < set.Blue {
				set_min.Blue = set.Blue
			}

			if set_min.Green < set.Green {
				set_min.Green = set.Green
			}
		}
		sum += set_min.Red * set_min.Blue * set_min.Green
	}

	fmt.Println("Q2: ", sum)
}
