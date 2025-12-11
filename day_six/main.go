package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PartOneProblems struct {
	Rows        [][]int
	OperatorRow []string
}

type PartTwoProblems struct {
	Rows        []string
	OperatorRow []string
}

func main() {
	problems, err := parseInputFile("data/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	total := 0
	for index, operator := range problems.OperatorRow {
		switch operator {
		case "+":
			total += problems.SumAtRowIndex(index)
		case "*":
			total += problems.MultiplyAtRowIndex(index)
		}
	}

	fmt.Printf("Total: %d\n\n", total)

	partTwoResult, err := parseFileForPartTwo("data/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	fmt.Printf("Part Two Total: %d\n\n", partTwoResult)
}

func (m PartOneProblems) SumAtRowIndex(index int) int {
	result := 0
	for _, row := range m.Rows {
		if len(row) > index {
			result += row[index]
		}
	}
	return result
}

func (m PartOneProblems) MultiplyAtRowIndex(index int) int {
	result := 1
	for _, row := range m.Rows {
		if len(row) > index {
			result *= row[index]
		}
	}
	return result
}

func parseInputFile(name string) (result PartOneProblems, err error) {
	file, err := os.Open(name)
	if err != nil {
		return result, fmt.Errorf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rowSlice := strings.Fields(scanner.Text())
		isOperatorRow := len(rowSlice) > 0 && rowSlice[0] == "*" || rowSlice[0] == "+"

		if isOperatorRow {
			result.OperatorRow = rowSlice
		} else {
			var intSlice []int

			for _, s := range rowSlice {
				if i, err := strconv.Atoi(s); err != nil {
					return result, fmt.Errorf("failed to parse int from numeric row: %s %s", s, err)
				} else {
					intSlice = append(intSlice, i)
				}
			}

			result.Rows = append(result.Rows, intSlice)
		}
	}

	return result, nil
}

func parseFileForPartTwo(name string) (result int, err error) {
	file, err := os.Open(name)
	if err != nil {
		return result, fmt.Errorf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var problems PartTwoProblems

	for scanner.Scan() {
		rowSlice := strings.Fields(scanner.Text())
		isOperatorRow := len(rowSlice) > 0 && rowSlice[0] == "*" || rowSlice[0] == "+"

		if isOperatorRow {
			problems.OperatorRow = rowSlice
		} else {
			problems.Rows = append(problems.Rows, scanner.Text())
		}
	}

	var operation []string
	operationIndex := 0
	total := 0

	// for each 0 to len(row):
	for i := 0; i < len(problems.Rows[0]); i++ {
		var chars []string
		// get a slice of each rows character value at that index, join them and remove whitespace
		for _, row := range problems.Rows {
			s := string(row[i])
			if s != " " {
				chars = append(chars, s)
			}
		}

		if len(chars) > 0 {
			operation = append(operation, strings.Join(chars, ""))
		}

		if len(chars) == 0 || i == len(problems.Rows[0])-1 {
			// If all are empty, it is the end of an operation the operation can be processed
			var opInts []int

			for _, s := range operation {
				if i, err := strconv.Atoi(s); err != nil {
					return result, fmt.Errorf("failed to parse int from numeric row: %s %s", s, err)
				} else {
					opInts = append(opInts, i)
				}
			}

			// Apply next operation and add to total
			switch problems.OperatorRow[operationIndex] {
			case "+":
				total += applyAddition(opInts)
			case "*":
				total += applyMultiplication(opInts)
			}

			// Reset operation buffer and increment operation
			operation = operation[:0]
			operationIndex++
		}
	}
	return total, nil
}

func applyAddition(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func applyMultiplication(nums []int) int {
	result := 1
	for _, num := range nums {
		result *= num
	}
	return result
}
