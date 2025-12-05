package day05

import (
	"sort"
	"strconv"
	"strings"
)

type RangeList []Range

type Range struct {
	Start, End int64
}

func isFresh(ranges RangeList, ingredient int64) bool {
	for _, r := range ranges {
		if ingredient >= r.Start && ingredient <= r.End {
			return true
		}
	}
	return false
}

func Day5_1(values []string) int {
	ranges := RangeList{}
	ingredients := []string{}
	count := 0

	for _, value := range values {
		if strings.Contains(value, "-") {
			// add to ranges
			StartEnd := strings.Split(value, "-")
			start, _ := strconv.ParseInt(StartEnd[0], 10, 64)
			end, _ := strconv.ParseInt(StartEnd[1], 10, 64)
			ranges = append(ranges, Range{start, end})
		} else if value == "" {
			// ignore
		} else {
			// add to ingredients
			ingredients = append(ingredients, value)
		}
	}

	for _, i := range ingredients {
		ingred, _ := strconv.ParseInt(i, 10, 64)

		if isFresh(ranges, ingred) {
			count++
		}

	}

	return count
}

func combineRanges(ranges RangeList) RangeList {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := ranges[:0]
	for _, r := range ranges {
		if r.End < r.Start {
			continue
		}

		n := len(merged)
		if n == 0 {
			merged = append(merged, r)
			continue
		}

		last := &merged[n-1]
		if r.Start <= last.End+1 {
			if r.End > last.End {
				last.End = r.End
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}

func Day5_2(values []string) int {
	ranges := RangeList{}

	for _, value := range values {
		if strings.Contains(value, "-") {
			// add to ranges
			StartEnd := strings.Split(value, "-")
			start, _ := strconv.ParseInt(StartEnd[0], 10, 64)
			end, _ := strconv.ParseInt(StartEnd[1], 10, 64)
			ranges = append(ranges, Range{start, end})
		}
	}

	combined := combineRanges(ranges)

	count := 0
	for _, i := range combined {
		count += int(i.End - i.Start + 1)
	}

	return count
}
