package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mazeLoop := Question1()
	Question2(mazeLoop)
}

type Maze []string

// Create a function to read the data
func ReadData() ([]string, [2]int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	s := bufio.NewScanner(file)
	s_id := 0
	maze := Maze{}
	SPosition := [2]int{}
	counter := 1
	for s.Scan() {
		// padding
		if s_id == 0 {
			maze = append(maze, strings.Repeat(".", len(s.Text())+2))
			s_id++
		}

		// read the line
		maze = append(maze, "."+s.Text()+".")

		// Find index of "S" in the line
		if i := strings.Index(s.Text(), "S"); i != -1 {
			SPosition[0] = counter
			SPosition[1] = i + 1

		}

		counter += 1

	}

	// padding
	maze = append(maze, strings.Repeat(".", len(maze[0])))

	return maze, SPosition
}

func Question1() [][]int {
	maze, StartPosition := ReadData()
	origStartRow, origStartCol := StartPosition[0], StartPosition[1]

	startRow, StartCol := StartPosition[0], StartPosition[1]
	mazeLoop := [][]int{}
	var prevRow, prevCol int

	for {
		//fmt.Println("SR", startRow, "SC", StartCol, "PR", prevRow, "PC", prevCol)

		if string(maze[startRow][StartCol]) == "S" {
			switch {
			case string(maze[startRow-1][StartCol]) == "|" || string(maze[startRow-1][StartCol]) == "F" || string(maze[startRow-1][StartCol]) == "7":
				prevRow = startRow
				prevCol = StartCol
				startRow--
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			case string(maze[startRow+1][StartCol]) == "|" || string(maze[startRow][StartCol+1]) == "L" || string(maze[startRow][StartCol+1]) == "J":
				prevRow = startRow
				prevCol = StartCol
				startRow++
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			case string(maze[startRow][StartCol-1]) == "-" || string(maze[startRow][StartCol-1]) == "L" || string(maze[startRow][StartCol-1]) == "F":
				prevRow = startRow
				prevCol = StartCol
				StartCol--
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			case string(maze[startRow][StartCol+1]) == "-" || string(maze[startRow][StartCol+1]) == "J" || string(maze[startRow][StartCol+1]) == "7":
				prevRow = startRow
				prevCol = StartCol
				StartCol++
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			}
		} else if string(maze[startRow][StartCol]) == "|" {
			switch {
			case (string(maze[startRow-1][StartCol]) == "|" || string(maze[startRow-1][StartCol]) == "7" || string(maze[startRow-1][StartCol]) == "F" || string(maze[startRow-1][StartCol]) == "S") && prevRow != startRow-1:
				prevRow = startRow
				prevCol = StartCol
				startRow--
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			case (string(maze[startRow+1][StartCol]) == "|" || string(maze[startRow+1][StartCol]) == "L" || string(maze[startRow+1][StartCol]) == "J" || string(maze[startRow+1][StartCol]) == "S") && prevRow != startRow+1:
				prevRow = startRow
				prevCol = StartCol
				startRow++
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			}
		} else if string(maze[startRow][StartCol]) == "-" {
			switch {
			case (string(maze[startRow][StartCol-1]) == "-" || string(maze[startRow][StartCol-1]) == "L" || string(maze[startRow][StartCol-1]) == "F" || string(maze[startRow][StartCol-1]) == "S") && prevCol != StartCol-1:
				prevRow = startRow
				prevCol = StartCol
				StartCol--
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			case (string(maze[startRow][StartCol+1]) == "-" || string(maze[startRow][StartCol+1]) == "J" || string(maze[startRow][StartCol+1]) == "7" || string(maze[startRow][StartCol+1]) == "S") && prevCol != StartCol+1:
				prevRow = startRow
				prevCol = StartCol
				StartCol++
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			}
		} else if string(maze[startRow][StartCol]) == "L" {
			switch {
			case (string(maze[startRow-1][StartCol]) == "|" || string(maze[startRow-1][StartCol]) == "F" || string(maze[startRow-1][StartCol]) == "7" || string(maze[startRow-1][StartCol]) == "S") && prevRow != startRow-1:
				prevRow = startRow
				prevCol = StartCol
				startRow--
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			case (string(maze[startRow][StartCol+1]) == "-" || string(maze[startRow][StartCol+1]) == "J" || string(maze[startRow][StartCol+1]) == "7" || string(maze[startRow][StartCol+1]) == "S") && prevCol != StartCol+1:
				prevRow = startRow
				prevCol = StartCol
				StartCol++
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			}
		} else if string(maze[startRow][StartCol]) == "J" {
			switch {
			case (string(maze[startRow-1][StartCol]) == "|" || string(maze[startRow-1][StartCol]) == "F" || string(maze[startRow-1][StartCol]) == "7" || string(maze[startRow-1][StartCol]) == "S") && prevRow != startRow-1:
				prevRow = startRow
				prevCol = StartCol
				startRow--
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			case (string(maze[startRow][StartCol-1]) == "-" || string(maze[startRow][StartCol-1]) == "L" || string(maze[startRow][StartCol-1]) == "F" || string(maze[startRow][StartCol-1]) == "S") && prevCol != StartCol-1:
				prevRow = startRow
				prevCol = StartCol
				StartCol--
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			}
		} else if string(maze[startRow][StartCol]) == "7" {
			switch {
			case (string(maze[startRow+1][StartCol]) == "|" || string(maze[startRow+1][StartCol]) == "L" || string(maze[startRow+1][StartCol]) == "J" || string(maze[startRow+1][StartCol]) == "S") && prevRow != startRow+1:
				prevRow = startRow
				prevCol = StartCol
				startRow++
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			case (string(maze[startRow][StartCol-1]) == "-" || string(maze[startRow][StartCol-1]) == "L" || string(maze[startRow][StartCol-1]) == "F" || string(maze[startRow][StartCol-1]) == "S") && prevCol != StartCol-1:
				prevRow = startRow
				prevCol = StartCol
				StartCol--
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			}
		} else if string(maze[startRow][StartCol]) == "F" {
			switch {
			case (string(maze[startRow][StartCol+1]) == "-" || string(maze[startRow][StartCol+1]) == "7" || string(maze[startRow][StartCol+1]) == "J" || string(maze[startRow][StartCol+1]) == "S") && prevCol != StartCol+1:
				prevRow = startRow
				prevCol = StartCol
				StartCol++
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			case (string(maze[startRow+1][StartCol]) == "|" || string(maze[startRow+1][StartCol]) == "J" || string(maze[startRow+1][StartCol]) == "L" || string(maze[startRow+1][StartCol]) == "S") && prevRow != startRow+1:
				prevRow = startRow
				prevCol = StartCol
				startRow++
				mazeLoop = append(mazeLoop, []int{startRow, StartCol})
			}
		}

		if startRow == origStartRow && StartCol == origStartCol {
			break
		}

	}
	fmt.Println("Q1: ", len(mazeLoop)/2)
	return mazeLoop

	// fmt.Println("Q1: ", answer)

}

func Question2(mazeLoop [][]int) {
	// Rebuild maze usiing solved maze
	// When LJ or JL is found, replace both with -. Same for F7 and 7F
	// When L7 or 7L is found, replace one with | and the other with -. Same for FJ and JF
	// So first, take solved maze and ignore "-", remove those from the solved maze solution
	// Then go through maze and look for the L7 or 7L and  FJ and JF. Replace one with | and the other with -
	// Finally, need ot understand what shape S is. Just manuualy do it
	maze, startVal := ReadData()
	for _, v := range maze {
		fmt.Println(v)
	}
	fmt.Println(startVal)

	//Manual Maze Intervention
	MazeLoopNoDash := [][]int{}
	// for i := range maze {
	// 	for j := range maze[i] {
	// 		if string(maze[i][j]) == "S" {
	// 			maze[i] = maze[i][:j] + "F" + maze[i][(j+1):]
	// 			MazeLoopNoDash = append(mazeLoop, []int{i, j})
	// 		}
	// 	}
	// }

	// remove "-" from mazeLoop
	for _, v := range mazeLoop {
		if string(maze[v[0]][v[1]]) != "-" {
			MazeLoopNoDash = append(MazeLoopNoDash, v)
		}
	}

	// Replace 7F, F7, JL, LJ
	for i := range MazeLoopNoDash {
		if string(maze[MazeLoopNoDash[i][0]][MazeLoopNoDash[i][1]]) == "7" {
			if string(maze[MazeLoopNoDash[i+1][0]][MazeLoopNoDash[i+1][1]]) == "F" {
				maze[MazeLoopNoDash[i][0]] = maze[MazeLoopNoDash[i][0]][:MazeLoopNoDash[i][1]] + "-" + maze[MazeLoopNoDash[i][0]][(MazeLoopNoDash[i][1]+1):]
				maze[MazeLoopNoDash[i+1][0]] = maze[MazeLoopNoDash[i+1][0]][:MazeLoopNoDash[i+1][1]] + "-" + maze[MazeLoopNoDash[i+1][0]][(MazeLoopNoDash[i+1][1]+1):]
			}
		}

		if string(maze[MazeLoopNoDash[i][0]][MazeLoopNoDash[i][1]]) == "F" {
			if string(maze[MazeLoopNoDash[i+1][0]][MazeLoopNoDash[i+1][1]]) == "7" {
				maze[MazeLoopNoDash[i][0]] = maze[MazeLoopNoDash[i][0]][:MazeLoopNoDash[i][1]] + "-" + maze[MazeLoopNoDash[i][0]][(MazeLoopNoDash[i][1]+1):]
				maze[MazeLoopNoDash[i+1][0]] = maze[MazeLoopNoDash[i+1][0]][:MazeLoopNoDash[i+1][1]] + "-" + maze[MazeLoopNoDash[i+1][0]][(MazeLoopNoDash[i+1][1]+1):]
			}
		}

		if string(maze[MazeLoopNoDash[i][0]][MazeLoopNoDash[i][1]]) == "J" {
			if string(maze[MazeLoopNoDash[i+1][0]][MazeLoopNoDash[i+1][1]]) == "L" {
				maze[MazeLoopNoDash[i][0]] = maze[MazeLoopNoDash[i][0]][:MazeLoopNoDash[i][1]] + "-" + maze[MazeLoopNoDash[i][0]][(MazeLoopNoDash[i][1]+1):]
				maze[MazeLoopNoDash[i+1][0]] = maze[MazeLoopNoDash[i+1][0]][:MazeLoopNoDash[i+1][1]] + "-" + maze[MazeLoopNoDash[i+1][0]][(MazeLoopNoDash[i+1][1]+1):]
			}
		}

		if string(maze[MazeLoopNoDash[i][0]][MazeLoopNoDash[i][1]]) == "L" {
			if string(maze[MazeLoopNoDash[i+1][0]][MazeLoopNoDash[i+1][1]]) == "J" {
				maze[MazeLoopNoDash[i][0]] = maze[MazeLoopNoDash[i][0]][:MazeLoopNoDash[i][1]] + "-" + maze[MazeLoopNoDash[i][0]][(MazeLoopNoDash[i][1]+1):]
				maze[MazeLoopNoDash[i+1][0]] = maze[MazeLoopNoDash[i+1][0]][:MazeLoopNoDash[i+1][1]] + "-" + maze[MazeLoopNoDash[i+1][0]][(MazeLoopNoDash[i+1][1]+1):]
			}
		}
	}

	// Replace L7, 7L, FJ, JF
	for i := range MazeLoopNoDash {
		if string(maze[MazeLoopNoDash[i][0]][MazeLoopNoDash[i][1]]) == "L" {
			if string(maze[MazeLoopNoDash[i+1][0]][MazeLoopNoDash[i+1][1]]) == "7" {
				maze[MazeLoopNoDash[i][0]] = maze[MazeLoopNoDash[i][0]][:MazeLoopNoDash[i][1]] + "|" + maze[MazeLoopNoDash[i][0]][(MazeLoopNoDash[i][1]+1):]
				maze[MazeLoopNoDash[i+1][0]] = maze[MazeLoopNoDash[i+1][0]][:MazeLoopNoDash[i+1][1]] + "-" + maze[MazeLoopNoDash[i+1][0]][(MazeLoopNoDash[i+1][1]+1):]
			}
		}

		if string(maze[MazeLoopNoDash[i][0]][MazeLoopNoDash[i][1]]) == "7" {
			if string(maze[MazeLoopNoDash[i+1][0]][MazeLoopNoDash[i+1][1]]) == "L" {
				maze[MazeLoopNoDash[i][0]] = maze[MazeLoopNoDash[i][0]][:MazeLoopNoDash[i][1]] + "|" + maze[MazeLoopNoDash[i][0]][(MazeLoopNoDash[i][1]+1):]
				maze[MazeLoopNoDash[i+1][0]] = maze[MazeLoopNoDash[i+1][0]][:MazeLoopNoDash[i+1][1]] + "-" + maze[MazeLoopNoDash[i+1][0]][(MazeLoopNoDash[i+1][1]+1):]
			}
		}
		if string(maze[MazeLoopNoDash[i][0]][MazeLoopNoDash[i][1]]) == "F" {
			if string(maze[MazeLoopNoDash[i+1][0]][MazeLoopNoDash[i+1][1]]) == "J" {
				maze[MazeLoopNoDash[i][0]] = maze[MazeLoopNoDash[i][0]][:MazeLoopNoDash[i][1]] + "|" + maze[MazeLoopNoDash[i][0]][(MazeLoopNoDash[i][1]+1):]
				maze[MazeLoopNoDash[i+1][0]] = maze[MazeLoopNoDash[i+1][0]][:MazeLoopNoDash[i+1][1]] + "-" + maze[MazeLoopNoDash[i+1][0]][(MazeLoopNoDash[i+1][1]+1):]
			}
		}
		if string(maze[MazeLoopNoDash[i][0]][MazeLoopNoDash[i][1]]) == "J" {
			if string(maze[MazeLoopNoDash[i+1][0]][MazeLoopNoDash[i+1][1]]) == "F" {
				maze[MazeLoopNoDash[i][0]] = maze[MazeLoopNoDash[i][0]][:MazeLoopNoDash[i][1]] + "|" + maze[MazeLoopNoDash[i][0]][(MazeLoopNoDash[i][1]+1):]
				maze[MazeLoopNoDash[i+1][0]] = maze[MazeLoopNoDash[i+1][0]][:MazeLoopNoDash[i+1][1]] + "-" + maze[MazeLoopNoDash[i+1][0]][(MazeLoopNoDash[i+1][1]+1):]
			}
		}
	}

	for _, v := range maze {
		fmt.Println(v)
	}

	Solver2(maze, mazeLoop)

}

// 544 too
func Solver2(maze []string, solvedPuzzle [][]int) int {
	// loop through maze row by row, left to right
	// Check if character is part of solved puzzle. If so, don't count
	// Otherwise, reading left to right, a point with an odd number of "|" on the left side counts
	// A point with an even number of || on the left side does not count

	total := 0
	for row := range maze {
		fmt.Println(maze[row])
		left := 0
	loop:
		for col := range maze[row] {

			for _, v := range solvedPuzzle {
				if row == v[0] && col == v[1] {
					if string(maze[row][col]) == "|" {
						left++
					}
					continue loop
				}
			}

			if left%2 != 0 {
				fmt.Println(row, col, string(maze[row][col]))
				total++
			}
		}
	}

	fmt.Println(total)
	return 0
}
