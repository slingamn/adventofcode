package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	//"strings"
)

var (
	ErrBadLine = errors.New("bad line")
)

type empty struct{}

func readStdin() (out []int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}

	return
}

func validate(val int, buf []int) bool {
	for i := 0; i < len(buf); i++ {
		for j := 0; j < len(buf); j++ {
			if i != j && buf[i]+buf[j] == val {
				return true
			}
		}
	}
	return false
}

func solve(preambleLen int, input []int) (result int) {
	buf := make([]int, preambleLen)

	for i := preambleLen; i < len(input); i++ {
		copy(buf, input[i-preambleLen:i])
		if !validate(input[i], buf) {
			return input[i]
		}
	}

	return -1
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	solution := solve(25, input)
	fmt.Println(solution)
}
