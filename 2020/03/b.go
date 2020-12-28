package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"errors"
)

const (
	lineLen = 31
)

type lineRep [lineLen]bool

type slope struct {
	right int
	down  int
}

var (
	ErrBadLine = errors.New("bad line")
)

// read a Grid in from stdin
func readStdin() (grid []lineRep, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\r\n")
		if len(line) != lineLen {
			err = ErrBadLine
			return
		}
		var lrep lineRep
		for i := 0; i < lineLen; i++ {
			lrep[i] = (line[i] == '#')
		}
		grid = append(grid, lrep)
	}
	return
}

func solve(grid []lineRep, s slope) (result int) {
	hpos := 0
	vpos := 0
	for vpos < len(grid) {
		if grid[vpos][hpos] {
			result++
		}
		hpos += s.right
		hpos = hpos % lineLen
		vpos += s.down
	}
	return
}

func main() {
	grid, err := readStdin()
	if err != nil {
		panic(err)
	}
	result := 1
	result = result * solve(grid, slope{right: 1, down: 1})
	result = result * solve(grid, slope{right: 3, down: 1})
	result = result * solve(grid, slope{right: 5, down: 1})
	result = result * solve(grid, slope{right: 7, down: 1})
	result = result * solve(grid, slope{right: 1, down: 2})
	fmt.Println(result)
}
