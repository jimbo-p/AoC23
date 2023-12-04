package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	Question1()
	//Question2()
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
