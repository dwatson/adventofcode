package day04

import (
	"strconv"
	"strings"
)

type ElfGroup struct {
	Assignments []Assignment
}

func (eg ElfGroup) Contained() bool {
	if eg.Assignments[0].Min <= eg.Assignments[1].Min && eg.Assignments[0].Max >= eg.Assignments[1].Max {
		return true
	}

	if eg.Assignments[1].Min <= eg.Assignments[0].Min && eg.Assignments[1].Max >= eg.Assignments[0].Max {
		return true
	}
	return false
}

func (eg ElfGroup) Overlapping() bool {
	if eg.Assignments[0].Max < eg.Assignments[1].Min || eg.Assignments[0].Min > eg.Assignments[1].Max {
		return false
	}
	if eg.Assignments[1].Max < eg.Assignments[0].Min || eg.Assignments[1].Min > eg.Assignments[0].Max {
		return false
	}
	return true
}

type Assignment struct {
	Data string
	Min  int
	Max  int
}

func Day4_1(values []string) int {
	egs := []ElfGroup{}
	fullycontained := 0
	for _, i := range values {
		eg := ElfGroup{}
		pair := strings.Split(i, ",")
		for _, j := range pair {
			sections := strings.Split(j, "-")
			min, _ := strconv.Atoi(sections[0])
			max, _ := strconv.Atoi(sections[1])
			eg.Assignments = append(eg.Assignments, Assignment{Data: j, Min: min, Max: max})
		}
		if eg.Contained() {
			fullycontained += 1
		}
		egs = append(egs, eg)
	}
	return fullycontained
}

func Day4_2(values []string) int {
	egs := []ElfGroup{}
	overlapped := 0
	for _, i := range values {
		eg := ElfGroup{}
		pair := strings.Split(i, ",")
		for _, j := range pair {
			sections := strings.Split(j, "-")
			min, _ := strconv.Atoi(sections[0])
			max, _ := strconv.Atoi(sections[1])
			eg.Assignments = append(eg.Assignments, Assignment{Data: j, Min: min, Max: max})
		}
		if eg.Overlapping() {
			overlapped += 1
		}
		egs = append(egs, eg)
	}
	return overlapped
}
