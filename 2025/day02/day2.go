package day02

import (
	"math"
	"strconv"
	"strings"

	"github.com/dwatson/adventofcode/2025/utils"
)

func getRanges(value string) []string {
	return strings.Split(value, ",")
}

func checkRangeSimple(valueRange string) int {
	runningTotal := 0
	StartEnd := strings.Split(valueRange, "-")
	start, _ := strconv.Atoi(StartEnd[0])
	end, _ := strconv.Atoi(StartEnd[1])
	for i := start; i < end+1; i++ {
		strValue := strconv.Itoa(i)
		if len(strValue)%2 == 0 {
			l := len(strValue)
			first := strValue[:l/2]
			second := strValue[l/2:]
			if first == second {
				runningTotal = runningTotal + i
			}
		}
	}

	return runningTotal
}

func checkRangeComplex(valueRange string) int {
	runningTotal := 0
	StartEnd := strings.Split(valueRange, "-")
	start, _ := strconv.Atoi(StartEnd[0])
	end, _ := strconv.Atoi(StartEnd[1])

	for i := start; i < end+1; i++ {
		if testNumber(i) {
			runningTotal += i
		}
	}

	return runningTotal
}

func factorsWithoutNum(num int) []int {
	factors := make([]int, 0)

	for i := 1; i < num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

func testNumber(num int) bool {
	digits := int(math.Floor(math.Log10(float64(num)))) + 1

	factors := factorsWithoutNum(digits)
	for _, chunkSize := range factors {
		nums := make([]int, 0)
		for i := chunkSize; i <= digits; i += chunkSize {
			divisor := int(math.Pow10(i))
			remainderDivisor := int(math.Pow10(i - chunkSize))

			remainder := num % divisor
			remainder = remainder / remainderDivisor

			nums = append(nums, remainder)
		}

		if utils.IsSliceAllEqual(nums) {
			return true
		}
	}

	return false
}

func Day2_1(values []string) int {
	var runningTotal int = 0

	valueRange := getRanges(values[0])

	for _, value := range valueRange {
		runningTotal = runningTotal + checkRangeSimple(value)
	}

	return runningTotal
}

func Day2_2(values []string) int {
	var runningTotal int = 0

	valueRange := getRanges(values[0])

	for _, value := range valueRange {
		runningTotal = runningTotal + checkRangeComplex(value)
	}

	return runningTotal
}
