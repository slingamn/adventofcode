package main

import (
	"bytes"
	"fmt"
	"bufio"
	"os"
)

// Grid is a maze / map
type Grid [][]byte
// Coordinate is a position on the map
type Coordinate [2]int
type Coordinates [4]Coordinate
// Given a starting Coordinate, Distances is a set of shortest-path distances from it
// to other Coordinates
type Distances map[Coordinate]int

var (
	// AdjacentDeltas are the deltas from a coordinate to its four adjacent coordinates
	// (you're allowed to move directly up or down, left or right, one step)
	AdjacentDeltas = []Coordinate{Coordinate{-1, 0}, Coordinate{1, 0}, Coordinate{0, -1}, Coordinate{0, 1}}
)

func isKey(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isDoor(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func (c Coordinate) Add(d Coordinate) (sum Coordinate) {
	return Coordinate{c[0] + d[0], c[1] + d[1]}
}

func (g Grid) Copy() (r Grid) {
	r = make([][]byte, len(g))
	for i, line := range g {
		r[i] = make([]byte, len(line))
		copy(r[i], line)
	}
	return
}

func (g Grid) Replace(orig, replacement byte) {
	for _, line := range g {
		j := bytes.IndexByte(line, orig)
		if j != -1 {
			line[j] = replacement
		}
	}
}

// Modifies a Grid to reflect having picked up a key. We treat both keys and doors as
// impenetrable obstacles, so we have to replace both the key and the door with a '.'
func (g Grid) Unlock(key byte) {
	var door byte
	door = key - 32
	g.Replace(key, '.')
	g.Replace(door, '.')
}


// Returns the value at a map coordinate, or the zero byte if out of bounds
func (g Grid) Get(coord Coordinate) (result byte) {
	i, j := coord[0], coord[1]
	// this is Go so we have to ask permission
	if i < 0 || i >= len(g) || j < 0 || j >= len(g[i]) {
		return 0
	}
	return g[i][j]
}

// Returns the number of keys in the map
func (g Grid) Length() (result int) {
	var maxByte byte
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] > maxByte {
				maxByte = g[i][j]
			}
		}
	}
	return int(maxByte - ('a' - 1))
}

func (g Grid) Print() {
	for _, line := range g {
		fmt.Println(string(line))
	}
}

// ComputeDistances computes the shortest-path distance to all remaining keys
func (g Grid) ComputeDistances(start Coordinate) (result Distances) {
	// breadth-first search
	result = make(Distances)

	allDistances := make(map[Coordinate]int)
	allDistances[start] = 0
	var queue []Coordinate
	queue = append(queue, start)
	for len(queue) > 0 {
		// dequeue
		pos := queue[0]
		queue = queue[1:]
		curDist := allDistances[pos]

		for _, delta := range AdjacentDeltas {
			adjCoord := pos.Add(delta)
			value := g.Get(adjCoord)
			if _, found := allDistances[adjCoord]; found {
				continue
			}
			switch {
			case value == 0:
				// not found
			case value == '#':
				// can't visit
			case isKey(value):
				// found shortest path to key, but can't visit
				result[adjCoord] = curDist + 1
			case isDoor(value):
				// can't visit
			case value == '.' || value == '@':
				// enqueue for visit
				allDistances[adjCoord] = curDist + 1
				queue = append(queue, adjCoord)
			}
		}
	}
	return
}

// VisitedSet is a representation of what keys have been collected, as a value type
// if we didn't have the maximum number of doors as a constant, we could use a string
// of the visited keys, sorted.
type VisitedSet [26]bool

// Progress is all information about how much "progress" a tour has made, for pruning purposes:
// if tourA.Progress == tourB.Progress && tourA.cost < tourB.cost, then tourB can be ignored
type Progress struct {
	position Coordinates // where we are now (position of last key collected)
	visited  VisitedSet // which keys have been collected
}

// Tour is a partial tour of the keys
type Tour struct {
	Progress
	order   [26]byte // the first `length` of these are the collected keys, in order
	length  int      // number of keys collected
	cost    int
}

// Given a tour, returns the tour that results from extending it with a visit to a new node
func (t *Tour) Visit(robot int, position Coordinate, door byte, cost int) (result Tour) {
	result = *t
	result.position[robot] = position
	idx := door - 'a'
	result.visited[idx] = true
	result.order[t.length] = door
	result.length += 1
	result.cost += cost
	return
}

func (t *Tour) Print() {
	fmt.Printf("%d: %s\n", t.cost, string(t.order[:t.length]))
}

// Computes the shortest tour of the remaining keys
func (g Grid) ShortestTour(start Coordinates) (result Tour) {
	numKeys := g.Length()

	// map of Progress to the cost of the shortest tour seen so far that makes that Progress,
	// used to prune suboptimal tours
	progressToMinCost := make(map[Progress]int)

	// breadth-first search again
	var queue []Tour

	var initialTour Tour
	initialTour.position = start
	queue = append(queue, initialTour)

	for len(queue) > 0 {
		// dequeue a tour
		tour := queue[0]
		queue = queue[1:]
		// tour may have been dominated by a subsequently discovered tour
		minCost, found := progressToMinCost[tour.Progress]
		if found && minCost < tour.cost {
			continue
		}

		// copy the map and update it to reflect our progress
		// XXX this is O(n) work, avoidable?
		curGrid := g.Copy()
		for i := 0; i < tour.length; i++ {
			curGrid.Unlock(tour.order[i])
		}
		// examine all possible extensions of the current tour
		for robot := 0; robot < 4; robot++ {
			distances := curGrid.ComputeDistances(tour.position[robot])
			for coord, distance := range distances {
				door := curGrid.Get(coord)
				nextTour := tour.Visit(robot, coord, door, distance)

				// did we find all the keys?
				if nextTour.length == numKeys {
					// is this the best full-length tour seen so far?
					if result.cost == 0 || result.cost > nextTour.cost {
						result = nextTour
					}
					// either way, no need to enqueue for further exploration
					continue
				}

				// enqueue this partial tour, but only if it is not dominated by a previously enqueued tour
				minCost, found := progressToMinCost[nextTour.Progress]
				if !found || nextTour.cost < minCost {
					queue = append(queue, nextTour)
					progressToMinCost[nextTour.Progress] = nextTour.cost
				}
			}
		}
	}

	return
}

// read a Grid in from stdin
func readStdin() (result Grid, start Coordinates) {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	z := 0
	for scanner.Scan() {
		line := []byte(scanner.Text())
		if j := bytes.IndexByte(line, '@'); j != -1 {
			start[z] = Coordinate{i,j}
			start[z+1] = Coordinate{i,j+2}
			z += 2
		}
		result = append(result, line)
		i++
	}

	return
}

func main() {
	grid, start := readStdin()
	grid.Print()
	for _, pos := range start {
		if grid[pos[0]][pos[1]] != '@' {
			panic("fail")
		}
	}
	shortest := grid.ShortestTour(start)
	fmt.Printf("shortest tour: ")
	shortest.Print()
}
