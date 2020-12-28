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
