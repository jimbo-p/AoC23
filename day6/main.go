package main

import "fmt"

type Race struct {
	Time     int
	Distance int
}

func main() {
	Question1()
	Question2()
}

// Create a function to read the data
func ReadData() []Race {
	// Races := []Race{
	// 	{Time: 7, Distance: 9},
	// 	{Time: 15, Distance: 40},
	// 	{Time: 30, Distance: 200},
	// }

	Races := []Race{
		{Time: 56, Distance: 546},
		{Time: 97, Distance: 1927},
		{Time: 78, Distance: 1131},
		{Time: 75, Distance: 1139},
	}

	return Races
}

// Create a function to read the data
func ReadData2() Race {
	// race := Race{
	// 	Time: 71530, Distance: 940200,
	// }

	race := Race{
		Time: 56977875, Distance: 546192711311139,
	}

	return race
}

// Calculate the score of a card
// 1 match = 1 point and then doubles for each additional match
// assumes no duplicates in your card
func Question1() {
	races := ReadData()

	waysToBeat := []int{}
	for _, race := range races {
		record := race.Distance
		for i := 0; i < race.Time; i++ {
			distance := (race.Time - i) * i
			if distance > record {
				beats := race.Time - 2*i

				waysToBeat = append(waysToBeat, beats+1)

				break
			}
		}
	}

	ans := 1
	for _, way := range waysToBeat {
		ans *= way
	}
	fmt.Println("Q1: ", ans)
}

func Question2() {
	race := ReadData2()

	waysToBeat := 0
	for i := 0; i < race.Time; i++ {
		distance := (race.Time - i) * i
		if distance > race.Distance {
			waysToBeat = race.Time - 2*i + 1
			break
		}
	}

	fmt.Println("Q2: ", waysToBeat)
}
