package day02

import (
	"github.com/dwatson/adventofcode/2022/utils"
	"reflect"
	"testing"
)

func TestDay2_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day02_example.input"), 15},
	}

	for _, tt := range testlist {
		got := Day2_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 1 - Example 1: %d\n", got)
		}
	}
}

func TestDay2_2(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day02.input"), 8933},
	}

	for _, tt := range testlist {
		got := Day2_1(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 1 - Part 1: %d\n", got)
		}
	}
}

func TestDay2_3(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day02_example.input"), 12},
	}

	for _, tt := range testlist {
		got := Day2_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 1 - Example 2: %d\n", got)
		}
	}
}

func TestDay2_4(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("day02.input"), 11998},
	}

	for _, tt := range testlist {
		got := Day2_2(tt.input)
		if !reflect.DeepEqual(tt.out, got) {
			t.Fatalf("expected: %v, got: %v\n", tt.out, got)
		} else {
			t.Logf("Day 1 - Part 2: %d\n", got)
		}
	}
}
