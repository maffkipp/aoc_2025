package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type JoltageList [][]int

func main() {
	input, err := parseInputFile("data/input.txt")
	if err != nil {
		log.Fatalf("Failed parse input from file: %s", err)
	}

	partOneTotal, err := calculateMaxJoltage(input, 2)
	if err != nil {
		log.Fatalf("Failed execute part one: %s", err)
	}

	partTwoTotal, err := calculateMaxJoltage(input, 12)
	if err != nil {
		log.Fatalf("Failed execute part one: %s", err)
	}

	fmt.Printf("\nTotal max joltage calculated in part one: %d\n", partOneTotal)
	fmt.Printf("\nTotal max joltage calculated in part two: %d\n", partTwoTotal)
}

func calculateMaxJoltage(input JoltageList, numDigits int) (result int, err error) {
	for _, bank := range input {
		var digitList []int
		var offset int

		// Loop through until our list of digits for the row is the correct length
		for len(digitList) < numDigits {
			maxIndexForDigit := len(bank) - (numDigits - len(digitList))
			highest := 0
			highestIndex := 0

			// We want to search between the index of our last found digit and the highest index the next digit could be
			for index, num := range bank[offset : maxIndexForDigit+1] {
				if num > highest {
					highest = num
					highestIndex = index
				}
			}
			digitList = append(digitList, highest)
			offset = offset + highestIndex + 1
		}

		// Collect our row into a concatenated string
		var resultString string

		for _, num := range digitList {
			resultString += strconv.Itoa(num)
		}

		// Then convert it back into an int and add it to our running total
		if combined, err := strconv.Atoi(resultString); err != nil {
			return result, err
		} else {
			result += combined
		}
	}

	return result, nil
}

func parseInputFile(name string) (result JoltageList, err error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file, %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line []int

		for _, s := range strings.Split(scanner.Text(), "") {
			num, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("Error parsing integer from: %s", s)
			}
			line = append(line, num)
		}
		result = append(result, line)
	}
	return result, nil
}
