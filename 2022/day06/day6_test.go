package day06

import (
	"github.com/dwatson/adventofcode/2022/utils"
	"reflect"
	"testing"
)

func TestDay6_1(t *testing.T) {
	var testlist = []struct {
		input string
		out   int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tt := range testlist {
		got := Day6_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 6 - Example 1: %d\n", got)
		}
	}
}

func TestDay6_2(t *testing.T) {
	var testlist = []struct {
		input string
		out   int
	}{
		{utils.ReadValues("day06.input")[0], 1235},
	}

	for _, tt := range testlist {
		got := Day6_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 6 - Part 1: %d\n", got)
		}
	}
}

func TestDay6_3(t *testing.T) {
	var testlist = []struct {
		input string
		out   int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tt := range testlist {
		got := Day6_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 6 - Example 2: %d\n", got)
		}
	}
}

func TestDay6_4(t *testing.T) {
	var testlist = []struct {
		input string
		out   int
	}{
		{utils.ReadValues("day06.input")[0], 3051},
	}

	for _, tt := range testlist {
		got := Day6_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 6 - Part 2: %d\n", got)
		}
	}
}
