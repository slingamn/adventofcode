package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	i int
	j int
	k int
	l int
}

type Grid map[Coordinate]byte

func readStdin() (input Grid, err error) {
	input = make(Grid)

	scanner := bufio.NewScanner(os.Stdin)
	k := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		for l, b := range line {
			input[Coordinate{0, 0, k, l}] = b
		}
		k++
	}

	return
}

func (t Grid) Copy() (r Grid) {
	r = make(Grid, len(t))
	for x, y := range t {
		r[x] = y
	}
	return
}

func countActive(t Grid, coord Coordinate) (result int) {
	for iDelt := -1; iDelt <= 1; iDelt++ {
		for jDelt := -1; jDelt <= 1; jDelt++ {
			for kDelt := -1; kDelt <= 1; kDelt++ {
				for lDelt := -1; lDelt <= 1; lDelt++ {
					if iDelt == 0 && jDelt == 0 && kDelt == 0 && lDelt == 0 {
						continue
					}
					if t[Coordinate{coord.i + iDelt, coord.j + jDelt, coord.k + kDelt, coord.l + lDelt}] == '#' {
						result++
					}
				}
			}
		}
	}
	return
}

func nextVal(count int, b byte) (next byte) {
	cur := b == '#'
	if cur {
		if count == 2 || count == 3 {
			next = '#'
		} else {
			next = '.'
		}
	} else {
		if count == 3 {
			next = '#'
		} else {
			next = '.'
		}
	}
	return
}

func iterate(grid Grid) (nextGrid Grid) {
	nextGrid = grid.Copy()

	for coord, b := range grid {
		cnt := countActive(grid, coord)
		nextGrid[coord] = nextVal(cnt, b)

		for iDelt := -1; iDelt <= 1; iDelt++ {
			for jDelt := -1; jDelt <= 1; jDelt++ {
				for kDelt := -1; kDelt <= 1; kDelt++ {
					for lDelt := -1; lDelt <= 1; lDelt++ {
						if iDelt == 0 && jDelt == 0 && kDelt == 0 && lDelt == 0 {
							continue
						}
						neigh := Coordinate{coord.i + iDelt, coord.j + jDelt, coord.k + kDelt, coord.l + lDelt}
						if _, found := grid[neigh]; found {
							continue
						}
						cnt := countActive(grid, neigh)
						nextGrid[neigh] = nextVal(cnt, 0)
					}
				}
			}
		}
	}
	return
}

func solve(grid Grid) (result int) {
	iterCount := 6

	for iter := 0; iter < iterCount; iter++ {
		grid = iterate(grid)
	}

	for _, b := range grid {
		if b == '#' {
			result++
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
