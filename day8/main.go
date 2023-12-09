package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Inst map[string]struct {
	Left  string
	Right string
}

var steps int = 0
var nextLoc string = ""

func main() {
	// LR, inst := ReadData()

	// if string(LR[0]) == "L" {
	// 	nextLoc = inst["AAA"].Left
	// } else {
	// 	nextLoc = inst["AAA"].Right
	// }

	// Question1(LR, inst)
	// fmt.Println("Q1: ", steps)

	Question2()
}

// Create a function to read the data
func ReadData() (string, Inst) {
	//LR := "LR"
	LR := "LRLRLLRRLLRRLRRLRRRLLRLRLRLRLRRLRRRLRLRRLRLLRRLLRLRRLRLRRLLRRRLRLRLRRRLRLLRRRLLLLLLRRRLRRLLLRRLRLRRLRRLRLRRLRRLLRRLRRRLRRRLLRLRLLLRRLLLRRLLRRLRLLRRRLRRRLRRRLRLRRLRRLLLRRRLRRLLRRLRRRLRLRLRRLRRLRRRLRRRLRLLLLRRRLRLRRRLRRRLLRLRRLRRLLRLLLRRLRLRRLRRRLRRRLRRRLLRRRLRLLRRRLRRRLRRRLRRRLRRLRRRLLRRLLRLRLRRRLRRRLRLRRRR"
	file, _ := os.Open("input.txt")
	defer file.Close()
	s := bufio.NewScanner(file)

	inst := Inst{}
	for s.Scan() {

		line := s.Text()
		split1 := strings.Split(line, " = ") // split1[0] = map key, split1[1] = map value

		lr := struct {
			Left  string
			Right string
		}{
			Left:  split1[1][1:4],
			Right: split1[1][6:9],
		}

		inst[split1[0]] = lr
	}

	return LR, inst
}

func Question1(LR string, inst Inst) {
	for _, v := range LR[1:] {

		if string(v) == "L" {
			nextLoc = inst[nextLoc].Left
		} else {
			nextLoc = inst[nextLoc].Right
		}
		steps++

		if nextLoc == "ZZZ" {
			steps++
			break
		}
	}

	if nextLoc != "ZZZ" {
		if string(LR[0]) == "L" {
			nextLoc = inst[nextLoc].Left
		} else {
			nextLoc = inst[nextLoc].Right
		}
		steps++
		Question1(LR, inst)
	}
}

type Node struct {
	StartNode         string
	NextNode          string
	FoundZ            bool
	CharToNextZ       int
	CharToGetToFirstZ int
}

// 10,000,000,000 too low
func Question2() {
	LR, inst := ReadData()
	nodes := []Node{}
	for k := range inst {
		if k[2] == 'A' {
			nodes = append(nodes, Node{StartNode: k, NextNode: ""})
		}
	}

	solvedNodes := []Node{}
	record := []string{}
	for _, node := range nodes {
		record = []string{}
		q2Steps := 0
		for {
			record = append(record, node.StartNode)
			direction := string(LR[q2Steps%(len(LR))]) // L or R
			node = OneStep(direction, inst, node)
			node.StartNode = node.NextNode
			q2Steps++

			if node.StartNode[2] == 'Z' {
				if node.FoundZ {
					node.CharToNextZ = q2Steps
					solvedNodes = append(solvedNodes, node)
					break
				} else {
					node.FoundZ = true
					node.CharToGetToFirstZ = q2Steps
					q2Steps = 0
				}
			}

		}
	}

	//fmt.Println(record)
	fmt.Println(solvedNodes)

	solveQ2(solvedNodes)
}

func OneStep(LR string, inst Inst, node Node) Node {

	if LR == "L" {
		node.NextNode = inst[node.StartNode].Left
	} else {
		node.NextNode = inst[node.StartNode].Right
	}

	return node
}

func solveQ2(nodes []Node) {
	vals := []int{}
	for _, node := range nodes {
		vals = append(vals, node.CharToNextZ)
	}

	// Initialize lcm with the first number
	result := vals[0]

	for _, num := range vals[1:] {
		result = lcm(result, num)
	}

	fmt.Println("LCM is:", result)

}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
