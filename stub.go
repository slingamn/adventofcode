package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ExerciseImports() {
	// XXX stop the compiler from complaining about unused imports:
	// sort, strconv, strings
	sort.Ints(nil)
	strconv.Atoi(strings.TrimPrefix("11", "1"))
}

var (
	ErrBadLine = errors.New("bad line")
)

type empty struct{}

// BEGIN LIBRARY ROUTINES

func countSetBits(v int) (c int) {
	// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan
	for c = 0; v != 0; c++ {
		v &= v - 1 // clear the least significant bit set
	}
	return
}

// Grid for cellular automata, mazes, etc.
type Grid [][]byte

// Coordinate is a position on a Grid, i.e., grid[y][x]
type Coordinate struct {
	y int
	x int
}

func (g Grid) Print() {
	for _, line := range g {
		fmt.Printf("%s\n", line)
	}
}

func (g Grid) Copy() (c Grid) {
	c = make(Grid, len(g))
	for i, line := range g {
		c[i] = make([]byte, len(line))
		copy(c[i], line)
	}
	return
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

// Absolute value of an integer (math.Abs is for float64 only)
func abs(i int) (r int) {
	if i < 0 {
		return -i
	}
	return i
}

// Python-style modulus function (-1 % 5 == 4, not -1)
func modulus(i int, m int) (r int) {
	r = i % m
	if r < 0 {
		r += m
	}
	return
}

// https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm
// compute gcd(a, b), return it with x and y such that ax + by = gcd(a, b)
func extendedEuclideanAlgorithm(a, b int) (gcd, x, y int) {
	var r, s, t, old_r, old_s, old_t int

	old_r, r = a, b
	old_s, s = 1, 0
	old_t, t = 0, 1

	for {
		if r == 0 {
			return old_r, old_s, old_t
		}
		quotient := old_r / r
		old_r, r = r, (old_r - quotient*r)
		old_s, s = s, (old_s - quotient*s)
		old_t, t = t, (old_t - quotient*t)
	}
}

// return x such that a*x % m == 1
func modularMultiplicativeInverse(a, m int) (ainv int) {
	gcd, x, _ := extendedEuclideanAlgorithm(a, m)
	if gcd != 1 {
		return 0 // could panic instead?
	}
	return modulus(x, m)
}

// compute x*y % m, avoiding integer overflow for intermediate values
func modularMultiply(x, y, m int) (result int) {
	negate := y < 0
	if negate {
		y = -y
	}
	y64 := uint64(y)

	for y64 != 0 {
		if (y64 & 1) != 0 {
			result = modulus(result+x, m)
		}
		x = modulus(x*2, m)
		y64 = y64 >> 1
	}

	if negate {
		result = modulus(-1*result, m)
	}
	return
}

// compute x**y % m, avoiding integer overflow for intermediate values
func modularExponentiate(x, y, m int) (result int) {
	negate := y < 0
	if negate {
		y = -y
	}
	y64 := uint64(y)

	result = 1
	for y64 != 0 {
		if (y64 & 1) != 0 {
			result = modulus(result*x, m)
		}
		x = modulus(x*x, m)
		y64 = y64 >> 1
	}

	if negate {
		result = modulus(-1*result, m)
	}
	return
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func parseInts(strs []string) (result []int) {
	var err error
	result = make([]int, len(strs))
	for i, s := range strs {
		result[i], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}
	return
}

// go doesn't have a set type or generics, don't @ me

type StringSet map[string]empty

func (s StringSet) Add(str string) {
	s[str] = empty{}
}

func (s StringSet) Has(str string) bool {
	_, ok := s[str]
	return ok
}

func (s StringSet) Copy() (c StringSet) {
	c = make(StringSet, len(s))
	for x := range s {
		c.Add(x)
	}
	return
}

func (s StringSet) Intersection(o StringSet) {
	for x := range s {
		if !o.Has(x) {
			delete(s, x)
		}
	}
	return
}

func (s StringSet) Union(o StringSet) {
	for x := range o {
		s.Add(x)
	}
}

func (s StringSet) Peek() (val string) {
	for x := range s {
		return x
	}
	return
}

type IntSet map[int]empty

func (s IntSet) Add(x int) {
	s[x] = empty{}
}

func (s IntSet) Has(x int) bool {
	_, ok := s[x]
	return ok
}

func (s IntSet) Copy() (c IntSet) {
	c = make(IntSet, len(s))
	for x := range s {
		c.Add(x)
	}
	return
}

func (s IntSet) Intersection(o IntSet) {
	for x := range s {
		if !o.Has(x) {
			delete(s, x)
		}
	}
	return
}

func (s IntSet) Union(o IntSet) {
	for x := range o {
		s.Add(x)
	}
}

func (s IntSet) Peek() (val int) {
	for x := range s {
		return x
	}
	return
}

// END LIBRARY ROUTINES

func solve(input []string) (result int) {
	result = len(input)
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
