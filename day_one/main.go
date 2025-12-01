package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction string

type Rotation struct {
	Direction Direction
	Distance  int
}

const (
	DirectionLeft  Direction = "L"
	DirectionRight Direction = "R"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	zeroCount := 0
	position := 50

	for scanner.Scan() {
		rotation, err := parseRotation(scanner.Text())
		if err != nil {
			log.Fatalf("Error parsing rotation value: %s", err)
		}

		// these allow us to account for rotation distances greater than one trip around
		zeroCount += rotation.Distance / 100
		effectiveDistance := rotation.Distance % 100

		switch rotation.Direction {
		case DirectionLeft:
			newPosition := rotateLeft(position, effectiveDistance)

			// additional checks here: should not tick when starting from 0, should tick when ending on 0
			if position != 0 && (newPosition > position || newPosition == 0) {
				zeroCount++
			}
			position = newPosition
		case DirectionRight:
			newPosition := rotateRight(position, effectiveDistance)

			if newPosition < position {
				zeroCount++
			}
			position = newPosition
		}
	}

	if scanner.Err() != nil {
		log.Fatalf("Failed to read lines from file: %s", scanner.Err())
	}

	fmt.Printf("Password is: %d", zeroCount)
}

func parseRotation(rawRotation string) (Rotation, error) {
	chars := strings.Split(rawRotation, "")

	if distance, err := strconv.Atoi(strings.Join(chars[1:], "")); err != nil {
		return Rotation{}, err
	} else if Direction(chars[0]) != DirectionLeft && Direction(chars[0]) != DirectionRight {
		return Rotation{}, fmt.Errorf("unsupported direction value: %s", chars[0])
	} else {
		return Rotation{Direction: Direction(chars[0]), Distance: distance}, nil
	}
}

func rotateLeft(position int, distance int) int {
	var result int

	if position-distance < 0 {
		result = position - distance + 100
	} else {
		result = position - distance
	}

	return result
}

func rotateRight(position int, distance int) int {
	var result int

	if position+distance >= 100 {
		result = position + distance - 100
	} else {
		result = position + distance
	}

	return result
}
