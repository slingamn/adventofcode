package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	//"sort"
	"strconv"
	"strings"
)

var (
	ErrBadLine = errors.New("bad line")
)

type empty struct{}

type Instruction struct {
	Addr int
	Val  int
	Mask string
}

func readStdin() (input []Instruction, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	maskCode := "mask = "
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, maskCode) {
			input = append(input, Instruction{Mask: strings.TrimPrefix(line, maskCode)})
			continue
		}

		startIdx := strings.IndexByte(line, '[')
		endIdx := strings.IndexByte(line, ']')
		addr, err := strconv.Atoi(line[startIdx+1 : endIdx])
		if err != nil {
			panic(err)
		}
		val, err := strconv.Atoi(line[endIdx+4:])
		if err != nil {
			panic(err)
		}
		input = append(input, Instruction{Addr: addr, Val: val})
	}

	return
}

func solve(input []Instruction) (result int) {
	memory := make(map[int]int)
	var mask string
	for _, instr := range input {
		if instr.Mask != "" {
			mask = instr.Mask
			continue
		}
		var val int
		for i := 0; i < 36; i++ {
			switch mask[35-i] {
			case 'X':
				val |= (instr.Val & (1 << i))
			case '0':
				// nothing
			case '1':
				val |= (1 << i)
			}
		}

		memory[instr.Addr] = val
	}

	sum := 0
	for _, val := range memory {
		sum += val
	}
	return sum
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(input))
}
