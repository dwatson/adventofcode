package day04

import (
	"github.com/dwatson/adventofcode/utils"
	"reflect"
	"testing"
)

func TestDay4_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day04_example.input"), 2},
	}

	for _, tt := range testlist {
		got := Day4_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 4 - Example 1: %d\n", got)
		}
	}
}

func TestDay4_2(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day04.input"), 450},
	}

	for _, tt := range testlist {
		got := Day4_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 4 - Part 1: %d\n", got)
		}
	}
}

func TestDay4_3(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day04_example.input"), 4},
	}

	for _, tt := range testlist {
		got := Day4_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 4 - Example 2: %d\n", got)
		}
	}
}

func TestDay4_4(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day04.input"), 837},
	}

	for _, tt := range testlist {
		got := Day4_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 4 - Part 2: %d\n", got)
		}
	}
}
