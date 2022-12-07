package day01

import (
	"github.com/dwatson/adventofcode/2022/utils"
	"reflect"
	"testing"
)

func TestDay1_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day01_example.input"), 24000},
	}

	for _, tt := range testlist {
		got := Day1_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 1 - Example 1: %d\n", got)
		}
	}
}

func TestDay1_2(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day01.input"), 70764},
	}

	for _, tt := range testlist {
		got := Day1_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 1 - Part 1: %d\n", got)
		}
	}
}

func TestDay1_3(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day01_example.input"), 45000},
	}

	for _, tt := range testlist {
		got := Day1_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 1 - Example 2: %d\n", got)
		}
	}
}

func TestDay1_4(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day01.input"), 203905},
	}

	for _, tt := range testlist {
		got := Day1_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 1 - Part 2: %d\n", got)
		}
	}
}
