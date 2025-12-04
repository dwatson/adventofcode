package day03

import (
	"reflect"
	"testing"

	"github.com/dwatson/adventofcode/2025/utils"
)

func TestDay4_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("example.input"), 13},
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
		{utils.ReadValues("input"), 1457},
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
		{utils.ReadValues("example.input"), 43},
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
		{utils.ReadValues("input"), 8310},
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
