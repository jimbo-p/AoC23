package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	Question1()
	//Question2()
}

type Universe struct {
	Rows []string
}

// init universe
var u Universe = Universe{}

// Create a function to read the data
func ReadData() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	s := bufio.NewScanner(file)
	for s.Scan() {
		u.Rows = append(u.Rows, s.Text())
	}
}

func Question1() {
	ReadData()
	u.Expand(2, 2)
	u.FindAndLabel()
	u.Distance()

}

func (u *Universe) Distance() {
	// Find manhattan distance between each #
	// add up distance for each pair
	// pairs to only be counted once

	for _, v := range u.Rows {
		fmt.Println(v)
	}

	distance := 0
	for row, v := range u.Rows {
		for col := range v {
			if u.Rows[row][col] != '.' {
				// Find all other #'s
				for row2, v2 := range u.Rows {
					for col2, _ := range v2 {
						val1, _ := strconv.Atoi(string(u.Rows[row][col]))
						val2, _ := strconv.Atoi(string(u.Rows[row2][col2]))
						if u.Rows[row2][col2] != '.' && val1 < val2 {
							// Calculate manhattan distance
							// Add to total
							distance += int(math.Abs(manhattanDistance(row, col, row2, col2)))
							fmt.Println("Distance between", string(u.Rows[row][col]), "and", string(u.Rows[row2][col2]), "is", manhattanDistance(row, col, row2, col2))
						}
					}
				}
			}
		}
	}

	fmt.Println("Q1:", distance)
}

func manhattanDistance(row1, col1, row2, col2 int) float64 {
	return float64(math.Abs((float64(row1) - float64(row2))) + math.Abs((float64(col1) - float64(col2))))
}

func (u *Universe) FindAndLabel() {
	// Find each # in the map and label it with a number counting upwards from the first
	counter := 1
	for row, v := range u.Rows {
		for col, _ := range v {
			if u.Rows[row][col] == '#' {
				u.Rows[row] = u.Rows[row][:col] + strconv.Itoa(counter) + u.Rows[row][col+1:]
				counter++
			}
		}
	}
}

func (u *Universe) Expand(rowExp, colExp int) {
	expandedU := Universe{}
	colExp--

	// Expand rows
	for _, v := range u.Rows {
		if isAllSameCharacter(v) {
			for i := 0; i < rowExp; i++ {
				expandedU.Rows = append(expandedU.Rows, v)
			}
		} else {
			expandedU.Rows = append(expandedU.Rows, v)
		}
	}

	// Search each column for same character
	// If the characters are all the same, expand the column by colExp

	fmt.Println("Col expansion")
	totalLoops := len(expandedU.Rows[0])
	currentCol := 0
	for col := 0; col < totalLoops; col++ {
		sameCharCheck := ""
		for _, r := range u.Rows {
			sameCharCheck += string(r[col])
		}

		expandedURowsandCols := Universe{}
		if isAllSameCharacter(sameCharCheck) {
			fmt.Println(currentCol)
			for rowE, _ := range expandedU.Rows {
				newRow := expandedU.Rows[rowE][:(currentCol)] + strings.Repeat(".", colExp) + expandedU.Rows[rowE][(currentCol):]
				expandedURowsandCols.Rows = append(expandedURowsandCols.Rows, newRow)
			}
			expandedU.Rows = expandedURowsandCols.Rows
		}

		if isAllSameCharacter(sameCharCheck) {
			currentCol += 2
		} else {
			currentCol++
		}

	}

	u.Rows = expandedU.Rows
}

func isAllSameCharacter(s string) bool {
	if len(s) <= 1 {
		return true
	}

	firstChar := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] != firstChar {
			return false
		}
	}

	return true
}

func Question2(mazeLoop [][]int) {

}
