package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Layout struct {
	Start     int
	Splitters [][]int
}

func main() {
	layout, err := parseInputFile("data/sample.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	splits, paths := computeLayout(layout)
	fmt.Printf("Part One Total: %+v\n\n", splits)
	fmt.Printf("Part Two Total: %+v\n\n", paths)
}

func computeLayout(layout Layout) (splits int, paths int) {
	pathMap := map[int]int{layout.Start: 1}
	totalSplits := 0

	for _, splitterRow := range layout.Splitters {
		for i, count := range pathMap {
			if count != 0 {
				if slices.Contains(splitterRow, i) {
					totalSplits++
					pathMap[i-1] += count
					pathMap[i+1] += count
					pathMap[i] = 0
				}
			}
		}
	}

	totalPaths := 0
	for _, count := range pathMap {
		totalPaths += count
	}
	return totalSplits, totalPaths
}

func parseInputFile(name string) (layout Layout, err error) {
	file, err := os.Open(name)
	if err != nil {
		return layout, fmt.Errorf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var splitters []int

		for index, char := range scanner.Text() {
			switch string(char) {
			case "S":
				layout.Start = index
			case "^":
				splitters = append(splitters, index)
			}
		}

		if len(splitters) > 0 {
			layout.Splitters = append(layout.Splitters, splitters)
		}
	}

	return layout, nil
}
