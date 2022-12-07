package day05

import (
	"github.com/dwatson/adventofcode/2022/utils"
	"reflect"
	"testing"
)

func TestDay5_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   string
	}{
		{utils.ReadValues("day05_example.input"), "CMZ"},
	}

	for _, tt := range testlist {
		got := Day5_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 5 - Example 1: %s\n", got)
		}
	}
}

func TestDay5_2(t *testing.T) {
	var testlist = []struct {
		input []string
		out   string
	}{
		{utils.ReadValues("day05.input"), "RLFNRTNFB"},
	}

	for _, tt := range testlist {
		got := Day5_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 5 - Part 1: %s\n", got)
		}
	}
}

func TestDay5_3(t *testing.T) {
	var testlist = []struct {
		input []string
		out   string
	}{
		{utils.ReadValues("day05_example.input"), "MCD"},
	}

	for _, tt := range testlist {
		got := Day5_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 5 - Example 2: %s\n", got)
		}
	}
}

func TestDay5_4(t *testing.T) {
	var testlist = []struct {
		input []string
		out   string
	}{
		{utils.ReadValues("day05.input"), "MHQTLJRLB"},
	}

	for _, tt := range testlist {
		got := Day5_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 5 - Part 2: %s\n", got)
		}
	}
}
