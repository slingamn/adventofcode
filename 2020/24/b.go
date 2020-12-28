package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type empty struct{}

// Coordinate is a position on a Grid, i.e., grid[y][x]
type Coordinate struct {
	y int
	x int
}

type Direction uint8

const (
	East Direction = iota
	Southeast
	Southwest
	West
	Northwest
	Northeast
)

var dirToCode = map[Direction]string{
	East:      "e",
	Southeast: "se",
	Southwest: "sw",
	West:      "w",
	Northwest: "nw",
	Northeast: "ne",
}

func lex(str string) (result []Direction) {
	for len(str) != 0 {
		found := false
		for dir, code := range dirToCode {
			if strings.HasPrefix(str, code) {
				str = strings.TrimPrefix(str, code)
				result = append(result, dir)
				found = true
				break
			}
		}
		if !found {
			panic(str)
		}
	}
	return
}

func neighbors(c Coordinate) (n [6]Coordinate) {
	return [6]Coordinate{
		{c.y, c.x + 1},
		{c.y + 1, c.x},
		{c.y + 1, c.x - 1},
		{c.y, c.x - 1},
		{c.y - 1, c.x + 1},
		{c.y - 1, c.x},
	}
}

func next(black bool, nCount int) (result bool) {
	if black {
		if nCount == 0 || nCount > 2 {
			return false
		} else {
			return true
		}
	} else {
		if nCount == 2 {
			return true
		} else {
			return false
		}
	}
}

func iterate(current map[Coordinate]bool) (result map[Coordinate]bool) {
	result = make(map[Coordinate]bool)
	edges := make(map[Coordinate]empty)

	for c, black := range current {
		nCount := 0
		for _, n := range neighbors(c) {
			if current[n] {
				nCount++
			}
			if _, ok := current[n]; !ok {
				edges[n] = empty{}
			}
		}
		result[c] = next(black, nCount)
	}

	for c := range edges {
		nCount := 0
		for _, n := range neighbors(c) {
			if current[n] {
				nCount++
			}
		}
		result[c] = next(false, nCount)
	}

	return
}

func solve(input []string) (result int) {
	state := make(map[Coordinate]bool)

	for _, str := range input {
		directions := lex(str)
		var y, x int
		for _, dir := range directions {
			switch dir {
			case East:
				x++
			case Southeast:
				y++
			case Southwest:
				x--
				y++
			case West:
				x--
			case Northeast:
				x++
				y--
			case Northwest:
				y--
			}
		}
		coord := Coordinate{y, x}
		state[coord] = !state[coord]
	}

	for i := 0; i < 100; i++ {
		state = iterate(state)
	}

	for _, black := range state {
		if black {
			result++
		}
	}
	return
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	solution := solve(input)
	fmt.Println(solution)
}
