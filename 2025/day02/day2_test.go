package day02

import (
	"reflect"
	"testing"

	"github.com/dwatson/adventofcode/2025/utils"
)

func TestDay2_1(t *testing.T) {
	var testlist = []struct {
		input []string
		out   int
	}{
		{utils.ReadValues("example.input"), 1227775554},
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
		{utils.ReadValues("input"), 31839939622},
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
		{utils.ReadValues("example.input"), 4174379265},
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
		{utils.ReadValues("input"), 41662374059},
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
