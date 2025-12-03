package day01

import (
	"strconv"
)

type Dial struct {
	position  int
	zeroCount int
}

func (eg Dial) Move(direction string, distance int) int {

	return 0
}

func getMove(move string) (string, int) {
	direction := string(move[0])
	distance, _ := strconv.Atoi(move[1:])
	return direction, distance
}

func doMove(start int, direction string, distance int) (int, int) {
	var count int = 0
	for i := 0; i < distance; i++ {
		if direction == "R" {
			start = start + 1
			if start > 99 {
				start = 0
			}
		} else {
			start = start - 1
			if start < 0 {
				start = 99
			}
		}
		if start == 0 {
			count++
		}
	}
	return start, count
}

func Day1_1(values []string) int {
	var zeroCount int = 0
	dial := 50

	for _, value := range values {
		dir, dist := getMove(value)
		position, _ := doMove(dial, dir, dist)
		dial = position
		if dial == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func Day1_2(values []string) int {
	var zeroCount int = 0
	var dial int = 50

	for _, value := range values {
		dir, dist := getMove(value)
		position, count := doMove(dial, dir, dist)
		dial = position
		zeroCount = zeroCount + count
		if dial == 0 {
		}

	}

	return zeroCount
}
