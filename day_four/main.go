package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type RollMap [][]bool

func main() {
	input, err := parseInputFile("data/input.txt")
	if err != nil {
		log.Fatalf("Failed parse input from file: %s", err)
	}

	// part 1, also processes first state for part 2
	state, removed := processPaperRollMap(input)
	fmt.Printf("\nInitial count is: %d\n\n", removed)

	// part 2
	totalRemoved := removed
	for removed > 0 {
		state, removed = processPaperRollMap(state)
		totalRemoved += removed
	}
	fmt.Printf("\nTotal count is: %d\n\n", totalRemoved)
}

func processPaperRollMap(currentState RollMap) (nextState RollMap, removed int) {
	var countCache [][]int

	for rowIndex, row := range currentState {
		var rowCache []int
		var nextRowCache []int
		var nextStateRow []bool

		for index, roll := range row {
			count := 0

			// add count from previous row
			if len(countCache) >= rowIndex && rowIndex != 0 {
				count += countCache[rowIndex-1][index]
			}

			// add count from current row and add to cache if it doesnt exist
			if len(countCache) > rowIndex {
				count += countCache[rowIndex][index]
			} else {
				xCount := calculateXAdjacentCount(row, index)
				rowCache = append(rowCache, xCount)
				count += xCount
			}

			// add count from next row and add it to cache if it doesn't exist
			if len(countCache) > rowIndex+1 {
				count += countCache[rowIndex+1][index]
			} else if len(currentState) > len(countCache) {
				xCount := calculateXAdjacentCount(currentState[rowIndex+1], index)
				nextRowCache = append(nextRowCache, xCount)
				count += xCount
			}

			// increment the removed total if the surrounding roll count is under 5 (extra accounts for the center roll)
			if roll && count < 5 {
				removed++
				nextStateRow = append(nextStateRow, false)
			} else if roll {
				nextStateRow = append(nextStateRow, true)
			} else {
				nextStateRow = append(nextStateRow, false)
			}
		}

		// Update caches and next state
		if len(rowCache) > 0 {
			countCache = append(countCache, rowCache)
		}

		if len(nextRowCache) > 0 {
			countCache = append(countCache, nextRowCache)
		}

		nextState = append(nextState, nextStateRow)
	}

	return nextState, removed
}

func calculateXAdjacentCount(row []bool, index int) int {
	count := 0

	if row[index] {
		count++
	}

	// Check element before if it exists
	if index != 0 && row[index-1] {
		count++
	}

	// And after
	if index != len(row)-1 && row[index+1] {
		count++
	}
	return count
}

func parseInputFile(name string) (result RollMap, err error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file, %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line []bool

		for _, s := range strings.Split(scanner.Text(), "") {
			if s == "@" {
				line = append(line, true)
			} else {
				line = append(line, false)
			}
		}
		result = append(result, line)
	}
	return result, nil
}
