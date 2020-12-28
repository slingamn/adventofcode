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

type empty struct{}

type Grid [][]byte

// Coordinate is a position on the map
type Coordinate struct {
	y int
	x int
}

func (g Grid) Print() {
	for _, line := range g {
		fmt.Printf("%s\n", line)
	}
}

// Returns the value at a map coordinate, or the zero byte if out of bounds
func (g Grid) Get(i, j int) (result byte) {
	if i < 0 || i >= len(g) || j < 0 || j >= len(g[i]) {
		return 0
	}
	return g[i][j]
}

func (g Grid) GetCoord(coord Coordinate) (result byte) {
	return g.Get(coord.y, coord.x)
}

func (g Grid) SetCoord(coord Coordinate, v byte) {
	g[coord.y][coord.x] = v
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
					mult := 1
					for {
						ydelthere := ydelt*mult
						xdelthere := xdelt*mult
						see := g.Get(i+ydelthere, j+xdelthere)
						done := false
						switch see {
						case 0, 'L':
							done = true
						case '#':
							occupiedNeighbors++
							done = true
						}
						if done {
							break
						}
						mult++
					}
				}
			}
			switch g[i][j] {
			case 'L':
				if occupiedNeighbors == 0 {
					updates = append(updates, Coordinate{i, j})
				}
			case '#':
				if occupiedNeighbors >= 5 {
					updates = append(updates, Coordinate{i, j})
				}
			}
		}
	}

	for _, coord := range updates {
		g.SetCoord(coord, munge(g.GetCoord(coord)))
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
