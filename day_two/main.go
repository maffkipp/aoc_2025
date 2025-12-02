package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("data/input.txt")
	if err != nil {
		log.Fatalf("Failed to open file, %s", err)
	}

	idRanges := strings.Split(string(input), ",")

	partOneTotal, err := partOne(idRanges)
	if err != nil {
		log.Fatalf("Error calculating part one: %s", err)
	}

	partTwoTotal, err := partTwo(idRanges)
	if err != nil {
		log.Fatalf("Error calculating part two: %s", err)
	}

	fmt.Printf("\nTotal of invalid IDs for part one is: %d\n", partOneTotal)
	fmt.Printf("\nTotal of invalid IDs for part two is: %d\n", partTwoTotal)
}

func partOne(idRanges []string) (int, error) {
	total := 0

	for _, idRange := range idRanges {
		bounds, err := getBounds(idRange)
		if err != nil {
			return total, err
		}

		for num := bounds.Min; num <= bounds.Max; num++ {

			// If the string is an even length, check the first half against the second half
			if str := strconv.Itoa(num); len(str)%2 == 0 {
				halfway := len(str) / 2
				if str[:halfway] == str[halfway:] {
					total += num
				}
			}
		}

	}
	return total, nil
}

func partTwo(idRanges []string) (int, error) {
	total := 0

	for _, idRange := range idRanges {
		bounds, err := getBounds(idRange)
		if err != nil {
			return total, err
		}

		for num := bounds.Min; num <= bounds.Max; num++ {
			str := strconv.Itoa(num)

			// If a subsection's length evenly divides into the length of the string, we can check further
			for i := 1; i <= len(str)/2; i++ {
				if len(str)%i == 0 {

					// Compare our original string against the subsection repeated up to the original's length
					comparison := strings.Repeat(str[:i], len(str)/i)

					if str == comparison {
						total += num
						break
					}
				}
			}
		}

	}
	return total, nil
}

type Bounds struct {
	Min int
	Max int
}

func getBounds(idRange string) (Bounds, error) {
	if bounds := strings.Split(idRange, "-"); len(bounds) != 2 {
		return Bounds{}, fmt.Errorf("Invalid format for ID range: %s", bounds)
	} else if min, err := strconv.Atoi(bounds[0]); err != nil {
		return Bounds{}, fmt.Errorf("Invalid minimum bound: %s: %s", bounds[0], err)
	} else if max, err := strconv.Atoi(bounds[1]); err != nil {
		return Bounds{}, fmt.Errorf("Invalid maximum bound: %s: %s", bounds[1], err)
	} else {
		return Bounds{Min: min, Max: max}, nil
	}
}
