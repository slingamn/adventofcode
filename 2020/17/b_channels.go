package main

// XXX this is a bad version of b_final.go that uses channels instead of an
// unsynchronized iterator type, for microbenchmarking the synch overhead imposed
// by channels. buffering helps a little but i'm still seeing a 5x-10x slowdown

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

// iterate over a "box" or "n-orthotope" of Coordinate's;
// mins are the minimum values for each coordinate, maxes are the maximum values
func BoxIterator(mins, maxes Coordinate) (chan Coordinate) {
	c := make(chan Coordinate)
	go boxIterate(c, mins, maxes)
	return c
}

// iterate over a coordinate's neighbors
// XXX caller must skip the original coordinate manually
func NeighborIterator(coord Coordinate) (chan Coordinate) {
	var mins, maxes Coordinate
	for i := 0; i < len(coord); i++ {
		mins[i] = coord[i] - 1
		maxes[i] = coord[i] + 2
	}
	return BoxIterator(mins, maxes)
}

func boxIterate(c chan Coordinate, mins Coordinate, maxes Coordinate){
	pos := mins

	for {
		c <- pos

		done := true
		// add with carry
		for i := 0; i < len(pos); i++ {
			pos[i]++
			if pos[i] == maxes[i] {
				pos[i] = mins[i]
			} else {
				done = false
				break
			}
		}

		if done {
			close(c)
			return
		}
	}
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
	for neigh := range NeighborIterator(coord) {
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
		for pos := range BoxIterator(newMins, newMaxes) {
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
