package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Bound struct {
	Min int
	Max int
}

type IngredientDB struct {
	FreshRanges          []Bound
	AvailableIngredients []int
}

func main() {
	db, err := parseInputFile("data/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	mergedFreshRanges := mergeOverlappingRanges(db.FreshRanges)

	// Part One
	freshCount := 0
	for _, ingredient := range db.AvailableIngredients {
		for _, freshRange := range mergedFreshRanges {
			if ingredient >= freshRange.Min && ingredient <= freshRange.Max {
				freshCount++
				break
			}
		}
	}

	fmt.Printf("\nTotal number of fresh ingredients: %d\n\n", freshCount)

	// Part Two
	totalIdCount := 0
	for _, freshRange := range mergedFreshRanges {
		totalIdCount += freshRange.Max - freshRange.Min + 1
	}

	fmt.Printf("\nTotal number of fresh IDs: %d\n\n", totalIdCount)
}

func mergeOverlappingRanges(freshRanges []Bound) (result []Bound) {
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i].Min <= freshRanges[j].Min
	})

	for _, freshRange := range freshRanges {
		if len(result) > 0 && result[len(result)-1].Max >= freshRange.Min {

			// When the next range is entirely contained within the previous, it can be discarded
			if freshRange.Max >= result[len(result)-1].Max {
				result[len(result)-1].Max = freshRange.Max
			}
		} else {
			result = append(result, freshRange)
		}
	}

	return result
}

func parseInputFile(name string) (result IngredientDB, err error) {
	file, err := os.Open(name)
	if err != nil {
		return result, fmt.Errorf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanningFreshRanges := true

	for scanner.Scan() {
		// A line break divides the two different data types in the input
		if scanner.Text() == "" {
			scanningFreshRanges = false
			continue
		}

		if scanningFreshRanges {
			bound, err := parseFreshRange(scanner.Text())
			if err != nil {
				return result, err
			}
			result.FreshRanges = append(result.FreshRanges, bound)
		} else {
			ingredient, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return result, fmt.Errorf("invalid ingredient: %s %s", scanner.Text(), err)
			}
			result.AvailableIngredients = append(result.AvailableIngredients, ingredient)
		}
	}
	return result, nil
}

// convert fresh ranges into a struct that is easier to work with
func parseFreshRange(line string) (bound Bound, err error) {
	var nums []int

	for _, s := range strings.Split(line, "-") {
		num, err := strconv.Atoi(s)
		if err != nil {
			return bound, fmt.Errorf("invalid bound: %s %s", bound, err)
		}
		nums = append(nums, num)
	}

	if len(nums) != 2 {
		return bound, fmt.Errorf("bound does not have one min and one max: %+v", bound)
	}

	return Bound{Min: nums[0], Max: nums[1]}, nil
}
