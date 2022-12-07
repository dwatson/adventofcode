package day01

import (
	"sort"
	"strconv"
)

func Day1_1(values []string) int {
	elfCalories := CountCalories(values)

	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))
	return elfCalories[0]
}

func Day1_2(values []string) int {
	elfCalories := CountCalories(values)

	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))
	return elfCalories[0] + elfCalories[1] + elfCalories[2]
}

func CountCalories(input []string) []int {
	var running = 0
	var elfCalories = []int{}

	for _, value := range input {
		i, _ := strconv.Atoi(value)
		if i != 0 {
			running += i
		} else {
			elfCalories = append(elfCalories, running)
			running = 0
		}
	}
	elfCalories = append(elfCalories, running)
	return elfCalories
}
