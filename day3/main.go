package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type RowCol struct {
	Row int
	Col int
}

// create map, key = index, value = number
type NumMap map[RowCol]int

func main() {
	Question1()
	//Question2()
}

// Create a function to read the data
func ReadData() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	// regex numbers, '[0-9]'
	nums := regexp.MustCompile("[0-9]+")
	// regex not a period, '.', or a number, '[0-9]'
	syms := regexp.MustCompile("[^0-9.]")

	s := bufio.NewScanner(file)
	numMap := make(NumMap)
	s_id := 0
	for s.Scan() {
		s_id += 1
		rowCol := RowCol{Row: s_id, Col: 0}

		text := s.Text()
		numbers := nums.FindAllString(text, -1)
		symbols := syms.FindAllString(text, -1)

		for i, char := range text {
			if string(char) == "." {
				s.Symbols = append(s.Symbols, Symbol{Symbol: string(char), Index: i})
			} else {
				n := Number{}
				n.Number = int(char)
				n.Index = i
				s.Numbers = append(s.Numbers, n)
			}
		}
		Schematics = append(Schematics, s)
	}
	return Schematics
}

func Question1() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	s := bufio.NewScanner(file)

	schematic := []string{}
	for s.Scan() {
		schematic = append(schematic, s.Text())
	}

	// regex numbers, '[0-9]'
	nums := regexp.MustCompile("[0-9]+")
	// regex not a period, '.', or a number, '[0-9]'
	symbols := regexp.MustCompile("[^0-9.]")

	fmt.Println(nums.FindAllString(schematic[0], -1))
	fmt.Println(symbols.FindAllString(schematic[1], -1))

	// unique symbols from regex query

	//fmt.Println("Q1: ", sum)
}

// func Question2() {

// 	fmt.Println("Q2: ", sum)
// }
