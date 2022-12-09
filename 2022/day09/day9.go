package day09

import (
	"strconv"
	"strings"
)

type void struct{}

var empty void

type move struct {
	direction vector
	count     int
}

// int version of math.Abs
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type vector struct {
	x, y int
}

func (v vector) add(z vector) vector {
	return vector{x: v.x + z.x, y: v.y + z.y}
}

func (v vector) subtract(z vector) vector {
	return vector{x: v.x - z.x, y: v.y - z.y}
}

func (v vector) Rounded() vector {
	var z vector

	switch {
	case v.x < 0:
		z.x = -1
	case v.x > 0:
		z.x = 1
	}

	switch {
	case v.y < 0:
		z.y = -1
	case v.y > 0:
		z.y = 1
	}
	return z
}

func moveTail(tail vector, head vector) vector {
	delta := head.subtract(tail)

	if Abs(delta.x) <= 1 && Abs(delta.y) <= 1 {
		// no move needed
		return tail
	}

	//move tail
	return tail.add(delta.Rounded())

}

func GetMoves(values []string) []move {
	moves := []move{}
	for _, item := range values {
		parts := strings.Split(item, " ")

		m := move{}
		m.count, _ = strconv.Atoi(parts[1])

		switch parts[0] {
		case "U":
			m.direction = vector{0, -1}
		case "D":
			m.direction = vector{0, 1}
		case "L":
			m.direction = vector{-1, 0}
		case "R":
			m.direction = vector{1, 0}
		}
		moves = append(moves, m)
	}
	return moves
}

func CountVisits(moves []move, rope []vector, knots int) int {
	tailVisited := make(map[vector]void)
	tailVisited[rope[knots-1]] = empty

	for _, m := range moves {
		for c := 0; c < m.count; c++ {
			rope[0] = rope[0].add(m.direction)

			for i := 1; i < knots; i++ {
				rope[i] = moveTail(rope[i], rope[i-1])
			}

			tailVisited[rope[knots-1]] = empty
		}
	}

	return len(tailVisited)
}

func Day9_1(values []string) int {
	moves := GetMoves(values)

	// rope has two knots
	rope := make([]vector, 2)
	return CountVisits(moves, rope, 2)

}

func Day9_2(values []string) int {
	moves := GetMoves(values)

	// rope has ten knots
	rope := make([]vector, 10)
	return CountVisits(moves, rope, 10)
}
