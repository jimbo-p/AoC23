package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	Question1()
	Question2()
}

func Question1() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("[0-9]+")

	calVals := make([]int, 0)
	for scanner.Scan() {
		l := scanner.Text()
		nums := strings.Join(re.FindAllString(l, -1), "")

		if len(nums) > 1 {
			v1, _ := strconv.Atoi(string(nums[0]))
			v2, _ := strconv.Atoi(string(nums[len(nums)-1]))
			calVals = append(calVals, v1*10+v2)
		} else {
			v1, _ := strconv.Atoi(string(nums[0]))
			calVals = append(calVals, v1*10+v1)
		}
	}

	// Sum the slice
	sum := 0
	for _, v := range calVals {
		sum += v
	}
	fmt.Println("Q1: ", sum)
}

func ReplaceWords(s string) string {
	digitMap := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
		"4": "four",
		"5": "five",
		"6": "six",
		"7": "seven",
		"8": "eight",
		"9": "nine",
	}

	for k, v := range digitMap {
		s = strings.Replace(s, v, k, -1)
	}

	return s
}

func Question2() {
	assoc := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	total := 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		chars := strings.Split(str, "")
		nums := make([]int, 0)

		for index, char := range chars {
			if n, err := strconv.Atoi(char); err == nil {
				nums = append(nums, n)
				continue
			}

			for key, value := range assoc {
				if len(key)+index > len(chars) {
					continue
				} else if strings.Join(chars[index:len(key)+index], "") == key {
					nums = append(nums, value)
				}
			}
		}
		part, err := strconv.Atoi(strconv.Itoa(nums[0]) + strconv.Itoa(nums[len(nums)-1]))
		if err != nil {
			log.Fatal(err)
		}
		total += part
	}
	fmt.Println("Q2: ", total)
}
