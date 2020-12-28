package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Python-style modulus function (-1 % 5 == 4, not -1)
func modulus(i int, m int) (r int) {
	r = i % m
	if r < 0 {
		r += m
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

// Returns the value at a map coordinate, or the zero byte if out of bounds
func (g Grid) Get(i, j int) (result byte) {
	if i < 0 || i >= len(g) || j < 0 || j >= len(g[i]) {
		return 0
	}
	return g[i][j]
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

func (g Grid) Equal(o Grid) bool {
	if len(g) != len(o) {
		return false
	}
	for i := 0; i < len(g); i++ {
		// according to bytes.go, this does not allocate
		if string(g[i]) != string(o[i]) {
			return false
		}
	}
	return true
}

func (g Grid) rotateOnce() (r Grid) {
	r = g.Copy()
	yMod := len(g)
	for i := 0; i < yMod; i++ {
		xMod := len(g[i])
		for j := 0; j < xMod; j++ {
			r[j][modulus(-1-i, yMod)] = g[i][j]
		}
	}
	return
}

func (g Grid) flipOnce() (r Grid) {
	r = g.Copy()
	for i := 0; i < len(g); i++ {
		mod := len(g[i])
		for j := 0; j < mod; j++ {
			r[i][j] = g[i][modulus(-1-j, mod)]
		}
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

const (
	sideLen = 10
)

func reverseBits(num int) (res int) {
	for i := 0; i < sideLen; i++ {
		if (num & (1 << ((sideLen - 1) - i))) != 0 {
			res |= 1 << i
		}
	}
	return
}

func intSqrt(num int) (res int) {
	res = int(math.Sqrt(float64(num)))
	if res*res != num {
		panic(num)
	}
	return
}

// the top, right, bottom, and left edges, in clockwise order,
// each time starting at the least significant bit of the integer
type Tile [4]int

func (t Tile) Rotate() (r Tile) {
	return Tile{t[3], t[0], t[1], t[2]}
}

func (t Tile) Flip() (r Tile) {
	return Tile{reverseBits(t[0]), reverseBits(t[3]), reverseBits(t[2]), reverseBits(t[1])}
}

func (t Tile) DihedralFour() (result [8]Tile) {
	result[0] = t
	result[1] = result[0].Rotate()
	result[2] = result[1].Rotate()
	result[3] = result[2].Rotate()
	result[4] = t.Flip()
	result[5] = result[4].Rotate()
	result[6] = result[5].Rotate()
	result[7] = result[6].Rotate()
	return
}

type TileArrangement [][]Tile

func compatible(a TileArrangement, y, x int, t Tile) bool {
	if y > 0 {
		if a[y-1][x][2] != reverseBits(t[0]) {
			return false
		}
	}
	if x > 0 {
		if a[y][x-1][1] != reverseBits(t[3]) {
			return false
		}
	}
	return true
}

func recursiveBacktrack(tiles map[int]Tile, result [][]int, rotations [][]uint8, tileArrangement [][]Tile, y, x int) (success bool) {
	if len(tiles) == 0 {
		return true
	}

	curTiles := make([]int, 0, len(tiles))
	for tileID := range tiles {
		curTiles = append(curTiles, tileID)
	}

	nextY := y
	nextX := (x + 1) % len(result[0])
	if nextX == 0 {
		nextY = (y + 1) % len(result)
	}

	for _, tileId := range curTiles {
		tile := tiles[tileId]
		delete(tiles, tileId)

		for rotIdx, r := range tile.DihedralFour() {
			if compatible(tileArrangement, y, x, r) {
				result[y][x] = tileId
				tileArrangement[y][x] = r
				rotations[y][x] = uint8(rotIdx)
				if recursiveBacktrack(tiles, result, rotations, tileArrangement, nextY, nextX) {
					return true
				}
			}
		}

		tiles[tileId] = tile
	}

	return false
}

func arrange(tiles map[int]Tile) (result [][]int, rotations [][]uint8) {
	squareSide := intSqrt(len(tiles))
	result = make([][]int, squareSide)
	rotations = make([][]uint8, squareSide)
	ta := make(TileArrangement, squareSide)
	for i := 0; i < squareSide; i++ {
		result[i] = make([]int, squareSide)
		ta[i] = make([]Tile, squareSide)
		rotations[i] = make([]uint8, squareSide)
	}

	success := recursiveBacktrack(tiles, result, rotations, ta, 0, 0)
	if !success {
		panic("fail")
	}
	return
}

func rotateGrid(grid Grid, rotateIdx uint8) (result Grid) {
	// XXX make sure we make a deep copy of the grid
	if rotateIdx >= 4 {
		rotateIdx -= 4
		result = grid.flipOnce()
	} else {
		result = grid.Copy()
	}

	var i uint8
	for i = 0; i < rotateIdx; i++ {
		result = result.rotateOnce()
	}
	return
}

func countSharps(grid Grid) (result int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				result++
			}
		}
	}
	return
}

var seaMonster []Coordinate = []Coordinate{
	{0, 18}, {1, 0}, {1, 5}, {1, 6}, {1, 11}, {1, 12}, {1, 17}, {1, 18}, {1, 19}, {2, 1}, {2, 4}, {2, 7}, {2, 10}, {2, 13}, {2, 16},
}

func countSeaMonsters(g Grid) (result int) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			success := true
			for _, delta := range seaMonster {
				if g.Get(i+delta.y, j+delta.x) != '#' {
					success = false
					break
				}
			}
			if success {
				result++
			}
			// TODO prevent overlapping monsters?
		}
	}
	return
}

func solve(input []string) (result int, err error) {
	tiles := make(map[int]Tile)
	grids := make(map[int]Grid)

	for len(input) > 0 {
		line := input[0]
		input = input[1:]
		tileNum := parseInt(line[5 : len(line)-1])
		var grid Grid
		for {
			grid = append(grid, []byte(input[0]))
			input = input[1:]
			if input[0] == "" {
				break
			}
		}

		var tile Tile
		for i := 0; i < sideLen; i++ {
			if grid[0][i] == '#' {
				tile[0] |= 1 << i
			}
		}
		for i := 0; i < sideLen; i++ {
			if grid[i][sideLen-1] == '#' {
				tile[1] |= 1 << i
			}
		}
		for i := 0; i < sideLen; i++ {
			if grid[sideLen-1][(sideLen-1)-i] == '#' {
				tile[2] |= 1 << i
			}
		}
		for i := 0; i < sideLen; i++ {
			if grid[(sideLen-1)-i][0] == '#' {
				tile[3] |= 1 << i
			}
		}

		tiles[tileNum] = tile
		grids[tileNum] = grid
		input = input[1:]
	}

	arrangement, rotations := arrange(tiles)

	squareSide := intSqrt(len(grids))
	cutSideLen := sideLen - 2
	picSide := cutSideLen * squareSide
	pic := make(Grid, picSide)
	for i := 0; i < len(pic); i++ {
		pic[i] = make([]byte, picSide)
	}

	for i := 0; i < len(arrangement); i++ {
		for j := 0; j < len(arrangement[i]); j++ {
			yPos := cutSideLen * i
			xPos := cutSideLen * j
			curPic := rotateGrid(grids[arrangement[i][j]], rotations[i][j])
			for yDelt := 0; yDelt < cutSideLen; yDelt++ {
				copy(pic[yPos+yDelt][xPos:], curPic[yDelt+1][1:sideLen-1])
			}
		}
	}

	sharpCount := countSharps(pic)

	var rot uint8
	for rot = 0; rot < 8; rot++ {
		r := rotateGrid(pic, rot)
		m := countSeaMonsters(r)
		if m != 0 {
			return sharpCount - (len(seaMonster) * m), nil
		}
	}

	panic("sea monsters not found")
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	solution, err := solve(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(solution)
}
