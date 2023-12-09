package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//Question1()
	Question2()
}

// Create a function to read the data
func ReadData() [][]int {
	file, _ := os.Open("input.txt")
	defer file.Close()
	s := bufio.NewScanner(file)
	data := [][]int{}

	for s.Scan() {
		dataline := []int{}
		line := s.Text()
		split := strings.Split(line, " ") // split1[0] = map key, split1[1] = map value
		for _, v := range split {
			val, _ := strconv.Atoi(v)
			dataline = append(dataline, val)
		}
		data = append(data, dataline)
	}
	return data
}

func Question1() {
	data := ReadData()

	fullData := [][][]int{}
	for _, values := range data {
		pattData := [][]int{}
		pattData = append(pattData, values)
		v := values
		for {
			patt := FindDiff(v)
			v = patt
			if slices.Equal(v, make([]int, len(v))) {
				pattData = append(pattData, append(patt, 0))
				break
			} else {
				pattData = append(pattData, patt)
			}
		}
		slices.Reverse(pattData)
		//fmt.Println(pattData)
		fullData = append(fullData, pattData)
	}

	sol := FindPattern(fullData)
	answer := 0
	for i := range sol {
		answer += sol[i][len(sol[i])-1][len(sol[i][len(sol[i])-1])-1]
	}

	fmt.Println("Q1: ", answer)

}

func FindDiff(values []int) []int {
	diff := []int{}
	for i := 1; i < len(values); i++ {
		diff = append(diff, values[i]-values[i-1])
	}
	return diff
}

func FindPattern(fullData [][][]int) [][][]int {
	for fullIndex := range fullData {
		for focusIndex := range fullData[fullIndex] {
			if focusIndex == len(fullData[fullIndex])-1 {
				break
			}
			newVal := fullData[fullIndex][focusIndex][len(fullData[fullIndex][focusIndex])-1] + fullData[fullIndex][focusIndex+1][len(fullData[fullIndex][focusIndex+1])-1]
			fullData[fullIndex][focusIndex+1] = append(fullData[fullIndex][focusIndex+1], newVal)
		}
	}

	return fullData
}

func Question2() {
	data := ReadData()

	fullData := [][][]int{}
	for _, values := range data {
		pattData := [][]int{}
		pattData = append(pattData, values)
		v := values
		for {
			patt := FindDiff(v)
			v = patt
			if slices.Equal(v, make([]int, len(v))) {
				pattData = append(pattData, append([]int{0}, patt...))
				break
			} else {
				pattData = append(pattData, patt)
			}
		}
		slices.Reverse(pattData)
		fullData = append(fullData, pattData)
	}

	sol := FindPattern2(fullData)
	answer := 0
	for i := range sol {
		answer += sol[i][len(sol[i])-1][0]
	}

	fmt.Println("Q2: ", answer)

}

func FindPattern2(fullData [][][]int) [][][]int {
	for fullIndex := range fullData {
		for focusIndex := range fullData[fullIndex] {
			if focusIndex == len(fullData[fullIndex])-1 {
				break
			}
			newVal := fullData[fullIndex][focusIndex+1][0] - fullData[fullIndex][focusIndex][0]
			fullData[fullIndex][focusIndex+1] = append([]int{newVal}, fullData[fullIndex][focusIndex+1]...)
		}
	}

	return fullData
}
