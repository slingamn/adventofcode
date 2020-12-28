package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	//"sort"
	//"strconv"
	//"strings"
)

var (
	ErrBadLine = errors.New("bad line")
)

type Grid [][]byte
// Coordinate is a position on the map
type Coordinate [2]int

func (g Grid) Print() {
	for _, line := range g {
		fmt.Println(string(line))
	}
}

type empty struct{}

// Returns the value at a map coordinate, or the zero byte if out of bounds
func (g Grid) Get(i, j int) (result byte) {
	// this is Go so we have to ask permission
	if i < 0 || i >= len(g) || j < 0 || j >= len(g[i]) {
		return 0
	}
	return g[i][j]
}

func munge(b byte) byte {
	if b == 'L' {
		return '#'
	}
	return 'L'
}

func (g Grid) Update() (updated bool) {
	var updates []Coordinate

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			occupiedNeighbors := 0
			for ydelt := -1; ydelt <= 1; ydelt++ {
				for xdelt := -1; xdelt <= 1; xdelt++ {
					if xdelt == 0 && ydelt == 0 {
						continue
					}
					if g.Get(i + ydelt, j + xdelt) == '#' {
						occupiedNeighbors++
					}
				}
			}
			switch g[i][j] {
			case 'L':
				if occupiedNeighbors == 0 {
					updates = append(updates, Coordinate{i, j})
				}
			case '#':
				if occupiedNeighbors >= 4 {
					updates = append(updates, Coordinate{i, j})
				}
			}
		}
	}

	for _, coord := range updates {
		i, j := coord[0], coord[1]
		g[i][j] = munge(g[i][j])
	}

	return len(updates) != 0
}

func readStdin() (input Grid, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []byte(line))
	}

	return
}

func solve(g Grid) (result int) {
	for g.Update() {
	}

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '#' {
				result++
			}
		}
	}
	return
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(input))
}
