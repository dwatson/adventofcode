package day03

import (
	"fmt"
)

type cell struct {
	x, y int
}

func DisplayCell(grid map[int]map[int]string, x int, y int) {
	fmt.Print("---\n")
	for i := x - 1; i <= x+1; i++ {
		fmt.Print("\n")
		for j := y - 1; j <= y+1; j++ {
			fmt.Print(grid[i][j])
		}
	}
	fmt.Print("---\n")
}

func isAccessible(grid map[int]map[int]string, x int, y int) bool {
	count := 0
	if grid[x-1][y-1] == "@" {
		count++
	}
	if grid[x-1][y] == "@" {
		count++
	}
	if grid[x-1][y+1] == "@" {
		count++
	}
	if grid[x][y-1] == "@" {
		count++
	}
	if grid[x][y+1] == "@" {
		count++
	}
	if grid[x+1][y-1] == "@" {
		count++
	}
	if grid[x+1][y] == "@" {
		count++
	}
	if grid[x+1][y+1] == "@" {
		count++
	}
	return count < 4
}

func countAccessible(grid map[int]map[int]string) (int, []cell) {
	count := 0
	cells := []cell{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			DisplayCell(grid, i, j)
			if grid[i][j] == "@" {
				if isAccessible(grid, i, j) {
					cells = append(cells, cell{i, j})
					count++
				}
			}
		}
	}
	return count, cells
}

func buildGrid(values []string) map[int]map[int]string {
	grid := map[int]map[int]string{}
	for i, row := range values {
		grid[i] = map[int]string{}
		for j, col := range row {
			grid[i][j] = string(col)
		}
	}
	return grid
}

func Day4_1(values []string) int {
	grid := buildGrid(values)
	accessible, _ := countAccessible(grid)
	return accessible
}

func Day4_2(values []string) int {
	grid := buildGrid(values)
	count := 0

	for ok := true; ok; {
		accessible, cells := countAccessible(grid)
		count += accessible
		if accessible > 0 {
			for i := 0; i < len(cells); i++ {
				grid[cells[i].x][cells[i].y] = "."
			}
		} else {
			break
		}
	}

	return count
}
