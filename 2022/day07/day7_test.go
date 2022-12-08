package day07

import (
	"github.com/dwatson/adventofcode/2022/utils"
	"reflect"
	"testing"
)

func TestDay7_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day07_example.input"), 95437},
	}

	for _, tt := range testlist {
		got := Day7_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 7 - Example 1: %d\n", got)
		}
	}
}

func TestDay7_2(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day07.input"), 2031851},
	}

	for _, tt := range testlist {
		got := Day7_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 7 - Part 1: %d\n", got)
		}
	}
}

func TestDay7_3(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day07_example.input"), 24933642},
	}

	for _, tt := range testlist {
		got := Day7_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 7 - Example 2: %d\n", got)
		}
	}
}

func TestDay7_4(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day07.input"), 2568781},
	}

	for _, tt := range testlist {
		got := Day7_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 7 - Part 2: %d\n", got)
		}
	}
}
