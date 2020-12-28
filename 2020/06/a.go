package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
)

var (
	ErrBadLine = errors.New("bad line")
)

type groupForm struct {
	questions [26]bool
	count int
}

func (g *groupForm) Add(char byte) {
	idx := uint8(char) - uint8('a')
	if !g.questions[idx] {
		g.questions[idx] = true
		g.count++
	}
}

func readStdin() (input []groupForm, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	var g groupForm
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			input = append(input, g)
			g = groupForm{}
			continue
		}
		for i := 0; i < len(line); i++ {
			g.Add(line[i])
		}
	}
	input = append(input, g)

	return
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, g := range input {
		sum += g.count
	}
	fmt.Println(sum)
}
