package day08

import (
	"github.com/dwatson/adventofcode/2022/utils"
	"reflect"
	"testing"
)

func TestDay8_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day08_example.input"), 21},
	}

	for _, tt := range testlist {
		got := Day8_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 8 - Example 1: %d\n", got)
		}
	}
}

func TestDay8_2(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day08.input"), 1717},
	}

	for _, tt := range testlist {
		got := Day8_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 8 - Part 1: %d\n", got)
		}
	}
}

func TestDay8_3(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day08_example.input"), 8},
	}

	for _, tt := range testlist {
		got := Day8_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 8 - Example 2: %d\n", got)
		}
	}
}

func TestDay8_4(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day08.input"), 321975},
	}

	for _, tt := range testlist {
		got := Day8_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 8 - Part 2: %d\n", got)
		}
	}
}
