package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DSR struct {
	Dest   int
	Source int
	Range  int
}

type DSR2 struct {
	Dest   uint32
	Source uint32
	Range  uint32
}

type DS struct {
	Dest   int
	Source int
}

type DS2 struct {
	Dest   uint32
	Source uint32
}

func main() {
	Question1()
	Question2()
}

func Process(seeds []int, dsr []DSR) []int {
	ds := DS{}

	newSeeds := []int{}
	for _, seed := range seeds {
		seedInDSR := false
		for _, x := range dsr {
			if seed >= x.Source && seed < (x.Source+x.Range) {
				seedInDSR = true
				ds.Source = seed
				ds.Dest = seed + (x.Dest - x.Source)
			}
		}

		if !seedInDSR {
			ds.Source = seed
			ds.Dest = seed
		}
		newSeeds = append(newSeeds, ds.Dest)
	}

	return newSeeds

}

func Process2(seeds []uint32, dsr []DSR2) []uint32 {
	ds := DS2{}

	newSeeds := []uint32{}
	for _, seed := range seeds {
		seedInDSR := false
		for _, x := range dsr {
			if seed >= x.Source && seed < (x.Source+x.Range) {
				seedInDSR = true
				ds.Source = seed
				ds.Dest = seed + (x.Dest - x.Source)
			}
		}

		if !seedInDSR {
			ds.Source = seed
			ds.Dest = seed
		}
		newSeeds = append(newSeeds, ds.Dest)
	}

	return newSeeds

}

// Create a function to read the data
func Question1() {
	seeds := []int{4043382508, 113348245, 3817519559, 177922221, 3613573568, 7600537, 773371046, 400582097, 2054637767, 162982133, 2246524522, 153824596, 1662955672, 121419555, 2473628355, 846370595, 1830497666, 190544464, 230006436, 483872831}
	file, _ := os.Open("input.txt")
	defer file.Close()
	s := bufio.NewScanner(file)

	cont_again := true
	dsr := []DSR{}
	for s.Scan() {

		if s.Text() == "" {
			// Call function
			seeds = Process(seeds, dsr)

			cont_again = true
			dsr = []DSR{}
			continue
		}

		if cont_again {
			cont_again = false
			continue
		}

		nums := strings.Split(s.Text(), " ")
		dest, _ := strconv.Atoi(nums[0])
		source, _ := strconv.Atoi(nums[1])
		rangeVal, _ := strconv.Atoi(nums[2])
		d := DSR{Dest: dest, Source: source, Range: rangeVal}
		dsr = append(dsr, d)

	}

	// get minimum seed
	minSeed := 99999999999
	for _, seed := range seeds {
		if seed < minSeed {
			minSeed = seed
		}
	}
	fmt.Println(minSeed)
}

// Create a function to read the data
func Question2() {
	counter := 0
	seeder := []uint32{4043382508, 3817519559, 3613573568, 773371046, 2054637767, 2246524522, 1662955672, 2473628355, 1830497666, 230006436}
	seederRanges := []uint32{113348245, 177922221, 7600537, 400582097, 162982133, 153824596, 121419555, 846370595, 190544464, 483872831}
	minSeedSlice := []uint32{}

	for i, seed := range seeder {
		seeds := []uint32{}

		fmt.Println("Seeds creates: ", i)
		for j := seed; j < seed+seederRanges[i]; j++ {
			seeds = append(seeds, j)
		}

		file, _ := os.Open("input.txt")
		defer file.Close()
		s := bufio.NewScanner(file)

		cont_again := true
		dsr := []DSR2{}

		for s.Scan() {

			if s.Text() == "" {
				counter++
				fmt.Println("Text progress: ", counter)
				// Call function
				seeds = Process2(seeds, dsr)

				cont_again = true
				dsr = []DSR2{}
				continue
			}

			if cont_again {
				cont_again = false
				continue
			}

			nums := strings.Split(s.Text(), " ")
			dest, _ := strconv.Atoi(nums[0])
			source, _ := strconv.Atoi(nums[1])
			rangeVal, _ := strconv.Atoi(nums[2])
			d := DSR2{Dest: uint32(dest), Source: uint32(source), Range: uint32(rangeVal)}
			dsr = append(dsr, d)

		}

		// get minimum seed
		var minSeed uint32 = 4294967295
		for _, seed := range seeds {
			if seed < minSeed {
				minSeed = seed
			}
		}
		minSeedSlice = append(minSeedSlice, minSeed)
	}
	fmt.Println(minSeedSlice)
}

// func Question2() {

// 	fmt.Println("Q2: ", total)
// }
