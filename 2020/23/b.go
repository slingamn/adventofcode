package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

const (
	roundCount = 10000000
	minCup     = 1
	maxCup     = 1000000
)

func solve(input []string) (result int) {
	var start []int
	for _, x := range input[0] {
		start = append(start, parseInt(string(x)))
	}

	// toNext[i] maps cup i to the cup that follows it in the sequence
	// in particular, toNext[0] is invalid because there is no cup 0
	toNext := make([]int, maxCup+1)
	for i := 0; i < len(start) - 1; i++ {
		toNext[start[i]] = start[i+1]
	}
	toNext[start[len(start)-1]] = len(start) + 1
	for j := len(start)+1; j < maxCup; j++ {
		toNext[j] = j+1
	}
	toNext[maxCup] = start[0]

	pos := start[0]
	for i := 0; i < roundCount; i++ {
		mov1 := toNext[pos]
		mov2 := toNext[mov1]
		mov3 := toNext[mov2]
		newPos := toNext[mov3]

		dst := pos
		for {
			dst -= 1
			if dst == 0 {
				dst = maxCup
			}
			if dst != mov1 && dst != mov2 && dst != mov3 {
				break
			}
		}

		dstNext := toNext[dst]
		toNext[dst] = mov1
		toNext[mov3] = dstNext
		toNext[pos] = newPos
		pos = newPos
	}

	result = 1
	successor := toNext[1]
	result *= successor
	successor = toNext[successor]
	result *= successor
	return
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	solution := solve(input)
	fmt.Println(solution)
}
