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
	ErrInfiniteLoop = errors.New("infinite loop")
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

func execute(instructions []Instruction) (accumulator int, err error) {
	visited := make(map[int]empty)
	pos := 0

	for {
		if pos >= len(instructions) {
			return accumulator, nil
		}
		if _, found := visited[pos]; found {
			return accumulator, ErrInfiniteLoop
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

	var accumulator int
	success := false
	for i := 0; i < len(input); i++ {
		code := input[i].Code
		if code == Acc {
			continue
		}
		if code == Jmp {
			input[i].Code = Nop
		} else  {
			input[i].Code = Jmp
		}
		accumulator, err = execute(input)
		if err == nil {
			success = true
			break
		}
		input[i].Code = code
	}
	if !success {
		panic("no")
	}
	fmt.Println(accumulator)
}
