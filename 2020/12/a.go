package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	//"sort"
	"strconv"
	//"strings"
)

var (
	ErrBadLine = errors.New("bad line")
)

type empty struct{}

type Instruction struct {
	Code byte
	Arg  int
}

func readStdin() (input []Instruction, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var instr Instruction
		instr.Code = line[0]
		instr.Arg, err = strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}
		input = append(input, instr)
	}

	return
}

func abs(i int) (r int) {
	if i < 0 {
		return -i
	}
	return i
}

func modulus(i int, m int) (r int) {
	r = i % m
	if r < 0 {
		r += m
	}
	return
}

func recomputeDir(dir byte, delta int) (newDir byte) {
	var angle int
	switch dir {
	case 'E':
		angle = 0
	case 'N':
		angle = 90
	case 'W':
		angle = 180
	case 'S':
		angle = 270
	}
	angle  = modulus(angle+delta, 360)
	switch angle {
	case 0:
		return 'E'
	case 90:
		return 'N'
	case 180:
		return 'W'
	case 270:
		return 'S'
	}
	return
}

func solve(input []Instruction) (result int) {
	var x, y int
	var dir byte
	dir = 'E'

	for _, instr := range input {
		code := instr.Code
		if code == 'F' {
			code = dir
		}

		switch code {
		case 'N':
			y += instr.Arg
		case 'S':
			y -= instr.Arg
		case 'E':
			x += instr.Arg
		case 'W':
			x -= instr.Arg
		case 'L':
			dir = recomputeDir(dir, instr.Arg)
		case 'R':
			dir = recomputeDir(dir, -instr.Arg)
		}
	}

	return abs(x)+abs(y)
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(input))
}
