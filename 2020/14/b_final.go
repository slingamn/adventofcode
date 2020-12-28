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

func compile(mask string) (orMask, xMask int) {
	for i := 0; i < 36; i++ {
		switch mask[35-i] {
		case '1':
			orMask |= 1 << i
		case 'X':
			xMask |= 1 << i
		}
	}
	return
}

func solve(input []Instruction) (result int) {
	memory := make(map[int]int)
	var orMask, xMask int
	for _, instr := range input {
		if instr.Mask != "" {
			orMask, xMask = compile(instr.Mask)
			continue
		}

		addr := instr.Addr | orMask
		// add-with-carry to enumerate all 2^n XOR masks
		xorMask := 0
		for {
			memory[addr^xorMask] = instr.Val

			added := false
			for i := 0; i < 36; i++ {
				b := 1 << i
				if xMask&b != 0 {
					if xorMask&b == 0 {
						// add in this place, we're done
						xorMask |= b
						added = true
						break
					} else {
						// zero this bit, carry to the next place
						xorMask &= ^b
					}
				}
			}
			if !added {
				// overflow
				break
			}
		}
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
