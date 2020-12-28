package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
	"strconv"
	//"strings"
)

var (
	ErrBadLine = errors.New("bad line")
)

type empty struct{}

type Opcode int

const (
	Nop Opcode = iota
	Acc
	Jmp
)

type Instruction struct {
	Code Opcode
	Arg int
}

func readStdin() (out []Instruction, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var instr Instruction
		switch line[:3] {
		case "nop":
			instr.Code = Nop
		case "acc":
			instr.Code = Acc
		case "jmp":
			instr.Code = Jmp
		default:
			return nil, ErrBadLine
		}
		instr.Arg, err = strconv.Atoi(line[4:])
		if err != nil {
			return nil, err
		}
		out = append(out, instr)
	}

	return
}

func execute(instructions []Instruction) (accumulator int) {
	visited := make(map[int]empty)
	pos := 0

	for {
		if _, found := visited[pos]; found {
			return accumulator
		}
		visited[pos] = empty{}
		instr := instructions[pos]
		switch instr.Code {
		case Nop:
			pos++
		case Acc:
			accumulator += instr.Arg
			pos++
		case Jmp:
			pos += instr.Arg
		}
	}
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(execute(input))
}
