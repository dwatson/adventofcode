package day03

import (
	"reflect"
	"testing"

	"github.com/dwatson/adventofcode/2025/utils"
)

func TestDay3_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("example.input"), 357},
	}

	for _, tt := range testlist {
		got := Day3_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 3 - Example 1: %d\n", got)
		}
	}
}

func TestDay3_2(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("input"), 17229},
	}

	for _, tt := range testlist {
		got := Day3_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 3 - Part 1: %d\n", got)
		}
	}
}

func TestDay3_3(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("example.input"), 3121910778619},
	}

	for _, tt := range testlist {
		got := Day3_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 3 - Example 2: %d\n", got)
		}
	}
}

func TestDay3_4(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("input"), 170520923035051},
	}

	for _, tt := range testlist {
		got := Day3_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 3 - Part 2: %d\n", got)
		}
	}
}
