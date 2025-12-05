package day05

import (
	"reflect"
	"testing"

	"github.com/dwatson/adventofcode/2025/utils"
)

func TestDay5_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("example.input"), 3},
	}

	for _, tt := range testlist {
		got := Day5_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 5 - Example 1: %d\n", got)
		}
	}
}

func TestDay5_2(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("input"), 758},
	}

	for _, tt := range testlist {
		got := Day5_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 5 - Part 1: %d\n", got)
		}
	}
}

func TestDay5_3(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("example.input"), 14},
	}

	for _, tt := range testlist {
		got := Day5_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 5 - Example 2: %d\n", got)
		}
	}
}

func TestDay5_4(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("input"), 343143696885053},
	}

	for _, tt := range testlist {
		got := Day5_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 5 - Part 2: %d\n", got)
		}
	}
}
