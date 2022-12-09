package day09

import (
	"github.com/dwatson/adventofcode/2022/utils"
	"reflect"
	"testing"
)

func TestDay9_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day09_example.input"), 13},
	}

	for _, tt := range testlist {
		got := Day9_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 9 - Example 1: %d\n", got)
		}
	}
}

func TestDay9_2(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day09.input"), 6522},
	}

	for _, tt := range testlist {
		got := Day9_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 9 - Part 1: %d\n", got)
		}
	}
}

func TestDay9_3(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day09_example.input"), 1},
	}

	for _, tt := range testlist {
		got := Day9_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 9 - Example 2: %d\n", got)
		}
	}
}

func TestDay9_4(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day09_example2.input"), 36},
	}

	for _, tt := range testlist {
		got := Day9_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 9 - Example 2 Part 2: %d\n", got)
		}
	}
}

func TestDay9_5(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day09.input"), 2717},
	}

	for _, tt := range testlist {
		got := Day9_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 9 - Part 2: %d\n", got)
		}
	}
}
