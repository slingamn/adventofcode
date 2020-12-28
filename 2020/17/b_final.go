package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	// this is fully parametrized by dimension; change this to 3 and you get
	// part a again
	dimensions = 4

	numIterations = 6
)

type Coordinate [dimensions]int

type BoxIterator struct {
	pos   Coordinate
	mins  Coordinate
	maxes Coordinate
}

// iterate over a "box" or "n-orthotope" of Coordinate's;
// mins are the minimum values for each coordinate, maxes are the maximum values
func (n *BoxIterator) Initialize(mins, maxes Coordinate) {
	n.pos = mins
	n.mins = mins
	n.maxes = maxes
}

// iterate over a coordinate's neighbors
// XXX caller must skip the original coordinate manually
func (n *BoxIterator) InitializeNeighbors(coord Coordinate) {
	for i := 0; i < len(coord); i++ {
		n.pos[i] = coord[i] - 1
		n.mins[i] = coord[i] - 1
		n.maxes[i] = coord[i] + 2
	}
}

// return the next coordinate; if `done`, then iteration is complete
// and the returned coordinate is invalid
func (n *BoxIterator) Next() (result Coordinate, done bool) {
	if n.pos[0] >= n.maxes[0] {
		done = true
		return
	}

	result = n.pos

	n.advance()

	return
}

func (n *BoxIterator) advance() {
	done := true

	// add with carry
	for i := 0; i < len(n.pos); i++ {
		n.pos[i]++
		if n.pos[i] == n.maxes[i] {
			n.pos[i] = n.mins[i]
		} else {
			done = false
			break
		}
	}

	if done {
		n.pos[0] = n.maxes[0]
	}

	return
}

type Grid struct {
	mins Coordinate
	maxes Coordinate
	storage []bool
}

func (g *Grid) Initialize(mins, maxes Coordinate) {
	prod := 1
	for i := 0; i < dimensions; i++ {
		prod *= (maxes[i] - mins[i])
		if prod <= 0 {
			panic("invalid")
		}
	}

	g.mins = mins
	g.maxes = maxes
	g.storage = make([]bool, prod)
}

func (g Grid) coordToIdx(c Coordinate) (idx int) {
	span := 1

	for i := 0; i < dimensions; i++ {
		if c[i] < g.mins[i] {
			return -1
		}
		if c[i] >= g.maxes[i] {
			return -1
		}
		idx += span * (c[i] - g.mins[i])
		span *= (g.maxes[i] - g.mins[i])
	}

	return idx
}

func (g Grid) Get(c Coordinate) bool {
	idx := g.coordToIdx(c)
	if idx == -1 {
		return false
	}

	return g.storage[idx]
}

func (g Grid) Set(c Coordinate, b bool) {
	idx := g.coordToIdx(c)
	g.storage[idx] = b
}

func countActiveNeighbors(t Grid, coord Coordinate) (result int) {
	var iter BoxIterator
	iter.InitializeNeighbors(coord)

	for {
		neigh, done := iter.Next()
		if done {
			break
		}
		if neigh == coord {
			continue
		}
		if t.Get(neigh) {
			result++
		}
	}

	return
}

func recomputeAlive(neighborCount int, currentStatus bool) (next bool) {
	switch neighborCount {
	case 2:
		return currentStatus
	case 3:
		return true
	default:
		return false
	}
}

func solve(grid Grid) (result int) {

	for iter := 0; iter < numIterations; iter++ {
		// expand space of interest by 1 in all directions
		var nextGrid Grid
		newMins := grid.mins
		newMaxes := grid.maxes
		for i := 0; i < dimensions; i++ {
			newMins[i]--
			newMaxes[i]++
		}
		nextGrid.Initialize(newMins, newMaxes)

		// iterate over all spaces in the expanded grid, compute new alive status
		var iter BoxIterator
		iter.Initialize(newMins, newMaxes)

		for {
			pos, done := iter.Next()
			if done {
				break
			}
			cnt := countActiveNeighbors(grid, pos)
			alive := recomputeAlive(cnt, grid.Get(pos))
			nextGrid.Set(pos, alive)
		}

		grid = nextGrid
	}

	for _, b := range grid.storage {
		if b {
			result++
		}
	}

	return
}

func toGrid(input []string) (grid Grid) {
	var mins, maxes Coordinate
	for i := 0; i < dimensions; i++ {
		maxes[i] = 1
	}
	maxes[dimensions-2] = len(input)
	maxes[dimensions-1] = len(input[0])
	grid.Initialize(mins, maxes)

	for i, line := range input {
		for j := 0; j < len(line); j++ {
			var coord Coordinate
			coord[dimensions-2] = i
			coord[dimensions-1] = j
			if line[j] == '#' {
				grid.Set(coord, true)
			}
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

	grid := toGrid(input)

	fmt.Println(solve(grid))
}
