package day08

import (
	"strconv"
)

func isVisible(grid map[int]map[int]int, x int, y int) bool {
	visible := 4
	// up - x smaller
	if x != 0 {
		v := true
		for i := x - 1; i >= 0; i-- {
			if grid[i][y] >= grid[x][y] {
				v = false
			}
		}
		if !v {
			visible -= 1
		}
	}
	// down - x bigger
	if x != len(grid) {
		v := true
		for i := x + 1; i < len(grid); i++ {
			if grid[i][y] >= grid[x][y] {
				v = false
			}
		}
		if !v {
			visible -= 1
		}
	}
	//left - y smaller
	if y != 0 {
		v := true
		for i := y - 1; i >= 0; i-- {
			if grid[x][i] >= grid[x][y] {
				v = false
			}
		}
		if !v {
			visible -= 1
		}
	}

	//right - y bigger
	if x != len(grid) {
		v := true
		for i := y + 1; i < len(grid); i++ {
			if grid[x][i] >= grid[x][y] {
				v = false
			}
		}
		if !v {
			visible -= 1
		}
	}
	return visible > 0
}

func ScoreUp(grid map[int]map[int]int, x int, y int) int {
	visibleTrees := 0
	for i := x - 1; i >= 0; i-- {
		if grid[i][y] < grid[x][y] {
			visibleTrees += 1
		} else if grid[i][y] >= grid[x][y] {
			visibleTrees += 1
			return visibleTrees
		}
	}
	return visibleTrees
}

map[int]string


func ScoreDown(grid map[int]map[int]int, x int, y int) int {
	visibleTrees := 0
	for i := x + 1; i < len(grid); i++ {
		if grid[i][y] < grid[x][y] {
			visibleTrees += 1
		} else if grid[i][y] >= grid[x][y] {
			visibleTrees += 1
			return visibleTrees
		}
	}
	return visibleTrees
}

func ScoreLeft(grid map[int]map[int]int, x int, y int) int {
	visibleTrees := 0
	for i := y - 1; i >= 0; i-- {
		if grid[x][i] < grid[x][y] {
			visibleTrees += 1
		} else if grid[x][i] >= grid[x][y] {
			visibleTrees += 1
			return visibleTrees
		}
	}
	return visibleTrees
}

func ScoreRight(grid map[int]map[int]int, x int, y int) int {
	visibleTrees := 0
	for i := y + 1; i < len(grid); i++ {
		if grid[x][i] < grid[x][y] {
			visibleTrees += 1
		} else if grid[x][i] >= grid[x][y] {
			visibleTrees += 1
			return visibleTrees
		}
	}
	return visibleTrees
}

func treeScore(grid map[int]map[int]int, x int, y int) int {
	score := 0
	u := ScoreUp(grid, x, y)
	d := ScoreDown(grid, x, y)
	l := ScoreLeft(grid, x, y)
	r := ScoreRight(grid, x, y)

	score = u * d * l * r
	return score
}

func Day8_1(values []string) int {
	grid := map[int]map[int]int{}
	for i, row := range values {
		grid[i] = map[int]int{}
		for j, col := range row {
			height, _ := strconv.Atoi(string(col))
			grid[i][j] = height
		}
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isVisible(grid, i, j) {
				count += 1
			}
		}
	}

	return count
}

func Day8_2(values []string) int {
	grid := map[int]map[int]int{}
	for i, row := range values {
		grid[i] = map[int]int{}
		for j, col := range row {
			height, _ := strconv.Atoi(string(col))
			grid[i][j] = height
		}
	}

	bestscore := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			currentTree := treeScore(grid, i, j)
			if currentTree > bestscore {
				bestscore = currentTree
			}
		}
	}
	return bestscore
}
