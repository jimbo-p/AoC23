package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/example/hello/reverse"
)

type Schematic []string

func main() {
	Question1()
	Question2()
}

// Create a function to read the data
func ReadData() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()
	s := bufio.NewScanner(file)
	s_id := 0
	schematic := []string{}
	for s.Scan() {
		fmt.Println(s.Text())
		// padding
		if s_id == 0 {
			schematic = append(schematic, strings.Repeat(".", len(s.Text())+2))
			s_id++
		}

		// read the line
		schematic = append(schematic, "."+s.Text()+".")
	}
	// padding
	schematic = append(schematic, strings.Repeat(".", len(schematic[0])))
	return schematic
}

func Question1() {
	s := ReadData()
	// regex numbers, '[0-9]'
	nums := regexp.MustCompile("[0-9]")
	// regex not a period, '.', or a number, '[0-9]'
	syms := regexp.MustCompile("[^0-9.]")
	// regex not a number, '[^0-9]'
	notNums := regexp.MustCompile("[^0-9]")

	numbers := []int{}
	for row, line := range s {
		if row == 0 || row == len(s)-1 {
			continue
		}
		IncludeNumber := false
		number := ""
		for col, char := range line {
			if nnums := notNums.Match([]byte(string(char))); nnums {
				if IncludeNumber {
					v, _ := strconv.Atoi(number)
					numbers = append(numbers, v)
				}
				number = ""
				IncludeNumber = false
			}

			// get character that is not a number or a period
			if symbol := nums.Match([]byte(string(char))); symbol {
				number += string(char)

				// check if any character around the number is a symbol include diagonals
				// if so, include the number
				if syms.Match([]byte(string(s[row][col-1]))) ||
					syms.Match([]byte(string(s[row][col+1]))) ||
					syms.Match([]byte(string(s[row-1][col]))) ||
					syms.Match([]byte(string(s[row+1][col]))) ||
					syms.Match([]byte(string(s[row-1][col-1]))) ||
					syms.Match([]byte(string(s[row-1][col+1]))) ||
					syms.Match([]byte(string(s[row+1][col-1]))) ||
					syms.Match([]byte(string(s[row+1][col+1]))) {
					IncludeNumber = true
				}
			}
		}
	}

	fmt.Println(numbers)
	// Sum the slice
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	fmt.Println("Q1: ", sum)
}

func Question2() {
	s := ReadData()
	// regex just the * symbol, '*'
	nostar := regexp.MustCompile("[^*]")
	// regex a number, '[0-9]'
	nums := regexp.MustCompile("[0-9]")
	total := 0

	//numbers := []int{}
	for row, line := range s {
		if row == 0 || row == len(s)-1 {
			continue
		}
		for col, char := range line {
			if notStar := nostar.Match([]byte(string(char))); notStar {
				continue
			} else {
				numLeft, numRight, numUp, numDown, numUpLeft, numUpRight, numDownLeft, numDownRight := 0, 0, 0, 0, 0, 0, 0, 0
				numLeft = SearchLeft(line[:col])
				numRight = SearchRight(line[col+1:])

				// if a number is not directly above the star
				if nums.Match([]byte{s[row-1][col]}) {
					numUp = SearchUpOrDown(s[row-1], col)
				} else {
					numUpLeft = SearchLeft(s[row-1][:col])
					numUpRight = SearchRight(s[row-1][col+1:])
				}

				if nums.Match([]byte{s[row+1][col]}) {
					numDown = SearchUpOrDown(s[row+1], col)
				} else {
					numDownLeft = SearchLeft(s[row+1][:col])
					numDownRight = SearchRight(s[row+1][col+1:])
				}

				// Check if two of the numbers are > 0
				numMap := map[int]int{
					numLeft:      numLeft,
					numRight:     numRight,
					numUp:        numUp,
					numDown:      numDown,
					numUpLeft:    numUpLeft,
					numUpRight:   numUpRight,
					numDownLeft:  numDownLeft,
					numDownRight: numDownRight,
				}

				fmt.Println(numMap, len(numMap))

				calcVal := 1
				if len(numMap) == 3 {
					for _, v := range numMap {
						if v != 0 {
							calcVal *= v
						}
					}
					total += calcVal
				}

			}
		}

	}

	// 84900915 too high
	fmt.Println("Q2: ", total)
}

// fmt.Println("Q1: ", sum)

// Given a string, find if there is a nubmer directly left of the * symbol
func SearchLeft(l string) int {
	r := regexp.MustCompile("[0-9]")
	number := ""
	for _, char := range reverse.String(l) {
		if r.Match([]byte(string(char))) {
			number += string(char)
		} else {
			break
		}
	}
	num, _ := strconv.Atoi(reverse.String(number))

	return num
}

func SearchRight(l string) int {
	r := regexp.MustCompile("[0-9]")
	number := ""
	for _, char := range l {
		if r.Match([]byte(string(char))) {
			number += string(char)
		} else {
			break
		}
	}
	num, _ := strconv.Atoi(number)

	return num
}

func SearchUpOrDown(l string, i int) int {
	number := string(l[i])
	notNum := regexp.MustCompile("[^0-9]")

	index := 0
	for {
		index++
		if notNum.Match([]byte{l[i-index]}) {
			break
		} else {
			number = string(l[i-index]) + number
		}
	}

	index = 0
	for {
		index++
		if notNum.Match([]byte{l[i+index]}) {
			break
		} else {
			number = number + string(l[i+index])
		}
	}

	num, _ := strconv.Atoi(string(number))

	return num
}
