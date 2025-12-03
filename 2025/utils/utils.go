package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadValues(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	var values []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error scanning file: %v\n", err)
	}
	return values
}

func SliceFind(slice []string, term string) int {
	for i, n := range slice {
		if term == n {
			return i
		}
	}
	return len(slice)
}

func IsSliceAllEqual(s []int) bool {
	if len(s) == 0 {
		return true
	}

	firstItem := s[0]
	for _, i := range s {
		if firstItem != i {
			return false
		}
	}

	return true
}
