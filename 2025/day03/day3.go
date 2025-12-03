package day03

import (
	"fmt"
	"strconv"
	"strings"
)

func findLargest(values string, startPoint int, endPoint int) (int, int) {
	fmt.Println("findLargest", startPoint, endPoint)
	highestValue := 0
	highestIndex := startPoint

	for i := startPoint; i < endPoint; i++ {
		numval, _ := strconv.Atoi(string(values[i]))
		if numval > highestValue {
			highestValue = numval
			highestIndex = i
		}
		if highestValue == 9 {
			break
		}
	}

	return highestValue, highestIndex
}

func Day3_1(values []string) int {
	runningTotal := 0

	for _, value := range values {
		valLen := len(value)
		startIndex := 0
		digitCount := 2

		var sb strings.Builder
		for i := digitCount - 1; i >= 0; i-- {
			fmt.Println("i", i)
			highVal, highIndex := findLargest(value, startIndex, valLen-i)
			startIndex = highIndex + 1
			fmt.Println(highVal)
			sb.WriteString(strconv.Itoa(highVal))
		}
		val, _ := strconv.Atoi(sb.String())
		fmt.Println("val", val)

		runningTotal += val
	}

	fmt.Println(runningTotal)
	return runningTotal
}

func Day3_2(values []string) int {
	runningTotal := 0

	for _, value := range values {
		valLen := len(value)
		startIndex := 0
		digitCount := 12

		var sb strings.Builder
		for i := digitCount - 1; i >= 0; i-- {
			fmt.Println("i", i)
			highVal, highIndex := findLargest(value, startIndex, valLen-i)
			startIndex = highIndex + 1
			fmt.Println(highVal)
			sb.WriteString(strconv.Itoa(highVal))
		}
		val, _ := strconv.Atoi(sb.String())
		fmt.Println("val", val)

		runningTotal += val
	}

	fmt.Println(runningTotal)
	return runningTotal
}
