package day05

import (
	"strconv"
	"strings"
)

type stack []string

type move struct {
	count  int
	source int
	dest   int
}

func parseInput(values []string) (map[int]stack, []move) {
	var stacks []string
	var moves []string

	for i, j := range values {
		if len(j) == 0 {
			stacks = values[:i-1]
			moves = values[i+1:]
		}
	}
	return parseStacks(stacks), parseMoves(moves)
}

func parseStacks(stackData []string) map[int]stack {
	stacks := map[int]stack{}

	for i := len(stackData); i >= 1; i-- {
		s := 1
		for j := 1; j <= len(stackData[i-1]); j += 4 {
			if string(stackData[i-1][j]) != " " {
				stacks[s] = append(stacks[s], string(stackData[i-1][j]))
			}
			s += 1
		}
	}
	return stacks
}

func parseMoves(moveData []string) []move {
	moves := []move{}
	for _, m := range moveData {
		sections := strings.Split(m, " ")
		count, _ := strconv.Atoi(sections[1])
		source, _ := strconv.Atoi(sections[3])
		dest, _ := strconv.Atoi(sections[5])
		moves = append(moves, move{source: source, count: count, dest: dest})
	}
	return moves
}

func GetTop(stacks map[int]stack) string {
	top := ""
	for i := 1; i <= len(stacks); i++ {
		top = top + stacks[i][len(stacks[i])-1]
	}
	return top
}

func Process(stacks map[int]stack, moves []move) string {
	for _, m := range moves {
		for i := 0; i < m.count; i++ {
			item := stacks[m.source][len(stacks[m.source])-1]
			stacks[m.source] = stacks[m.source][:len(stacks[m.source])-1]
			stacks[m.dest] = append(stacks[m.dest], item)
		}
	}
	return GetTop(stacks)
}

func ProcessOrdered(stacks map[int]stack, moves []move) string {
	for _, m := range moves {
		items := stacks[m.source][len(stacks[m.source])-m.count : len(stacks[m.source])]
		stacks[m.source] = stacks[m.source][:len(stacks[m.source])-m.count]
		stacks[m.dest] = append(stacks[m.dest], items...)
	}

	return GetTop(stacks)
}

func Day5_1(values []string) string {
	stacks, moves := parseInput(values)
	result := Process(stacks, moves)

	return result
}

func Day5_2(values []string) string {
	stacks, moves := parseInput(values)
	result := ProcessOrdered(stacks, moves)

	return result
}
