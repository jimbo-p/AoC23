package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	// Question1()
	Question2()
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
	intVerse := u.FindAndLabel()
	u.Distance(intVerse)
}

// 560822911938
// 82000210 too low
func Question2() {
	u = Universe{}
	ReadData()
	m := u.Expand2(1000000-1, 1000000-1)

	points := getKeys(m)
	keyCombs := keyCombinations(points)

	total := 0
	for _, kc := range keyCombs {
		// Get the values corresponding to each combination of keys
		val1, ok1 := m[kc[0]]
		val2, ok2 := m[kc[1]]
		if ok1 && ok2 {
			total += manhattanDistance2(val1.row, val1.col, val2.row, val2.col)
		}
	}

	fmt.Println(total)
}

type Point struct {
	row int
	col int
}

func (u *Universe) Expand2(rowExp, colExp int) map[Point]Point {
	// map of old point and new point
	expMap := make(map[Point]Point)
	for row, v := range u.Rows {
		for col, _ := range v {
			if u.Rows[row][col] == '#' {
				expMap[Point{row, col}] = Point{row, col}
			}
		}
	}

	fmt.Println(expMap)

	// Expand cols
	for row, v := range u.Rows {
		if isAllSameCharacter(v) {
			for k, val := range expMap {
				if k.row > row {
					expMap[k] = Point{val.row + rowExp, val.col}
				}
			}
		}
	}

	fmt.Println(expMap)

	// expand rows
	for col := 0; col < len(u.Rows[0]); col++ {
		sameCharCheck := ""
		for _, r := range u.Rows {
			sameCharCheck += string(r[col])
		}

		if isAllSameCharacter(sameCharCheck) {
			for k, v := range expMap {
				if k.col > col {
					expMap[k] = Point{v.row, v.col + colExp}
				}
			}
		}
	}

	fmt.Println(expMap)
	return expMap
}

// Function to collect keys from the map
func getKeys(m map[Point]Point) []Point {
	var keys []Point
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Function to generate unique combinations of keys
func keyCombinations(keys []Point) [][2]Point {
	var combs [][2]Point

	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			combs = append(combs, [2]Point{keys[i], keys[j]})
		}
	}

	return combs
}

func manhattanDistance2(row1, col1, row2, col2 int) int {
	return int(math.Abs((float64(row1) - float64(row2))) + math.Abs((float64(col1) - float64(col2))))
}

func (u *Universe) Distance(matrix [][]int16) {
	// Find manhattan distance between each #
	// add up distance for each pair
	// pairs to only be counted once

	distance := 0
	for row, v := range matrix {
		fmt.Println("Completed row: ", row)
		for col := range v {
			if matrix[row][col] != 0 {
				// Find all other #'s
				for row2, v2 := range matrix {
					for col2 := range v2 {
						val1 := matrix[row][col]
						val2 := matrix[row2][col2]
						if matrix[row2][col2] != 0 && val1 < val2 {
							// Calculate manhattan distance
							// Add to total
							distance += int(math.Abs(manhattanDistance(int16(row), int16(col), int16(row2), int16(col2))))
							// fmt.Println("Distance between", matrix[row][col], "and", matrix[row2][col2], "is", manhattanDistance(row, col, row2, col2))
						}
					}
				}
			}
		}
	}

	fmt.Println("Q1:", distance)
}

func manhattanDistance(row1, col1, row2, col2 int16) float64 {
	return float64(math.Abs((float64(row1) - float64(row2))) + math.Abs((float64(col1) - float64(col2))))
}

func (u *Universe) FindAndLabel() [][]int16 {
	// Find each # in the map and label it with a number counting upwards from the first
	matrix := make([][]int16, len(u.Rows))

	for i := range matrix {
		// Initialize each sub-slice (row) with 10 ints, all set to 0
		matrix[i] = make([]int16, len(u.Rows[i]))
	}

	var counter int16 = 1
	for row, v := range u.Rows {
		for col, _ := range v {
			if u.Rows[row][col] == '#' {
				matrix[row][col] = counter
				counter++
			} else {
				matrix[row][col] = 0
			}
		}
	}
	return matrix
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
