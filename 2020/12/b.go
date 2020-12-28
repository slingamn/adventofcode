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

func rotate(wx, wy, angle int) (nwx, nwy int) {
	angle = modulus(angle, 360)
	switch angle {
	case 0:
		return wx, wy
	case 90:
		return -wy, wx
	case 180:
		return -wx, -wy
	case 270:
		return wy, -wx
	default:
		panic(angle)
	}
}

func solve(input []Instruction) (result int) {
	var x, y int
	wx := 10
	wy := 1

	for _, instr := range input {
		code := instr.Code

		switch code {
		case 'N':
			wy += instr.Arg
		case 'S':
			wy -= instr.Arg
		case 'E':
			wx += instr.Arg
		case 'W':
			wx -= instr.Arg
		case 'L':
			wx, wy = rotate(wx, wy, instr.Arg)
		case 'R':
			wx, wy = rotate(wx, wy, -instr.Arg)
		case 'F':
			x += instr.Arg*wx
			y += instr.Arg*wy
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
