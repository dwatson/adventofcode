package day06

type void struct{}

var empty void

func FindUniqueCharsGroup(s string, count int) int {
	for i := count; i < len(s)-1; i++ {
		candidate := s[i-count : i]
		fakeSet := make(map[rune]void)
		for _, r := range candidate {
			fakeSet[r] = empty
		}
		if len(fakeSet) == count {
			return i
		}
	}
	return 0
}

func Day6_1(value string) int {
	return FindUniqueCharsGroup(value, 4)
}

func Day6_2(value string) int {
	return FindUniqueCharsGroup(value, 14)
}
