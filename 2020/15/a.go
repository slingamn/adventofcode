package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readStdin() (input []int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	ints := strings.Split(line, ",")
	for _, s := range ints {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		input = append(input, i)
	}

	return
}

func solve(input []int) (result int) {
	rd := 1
	lastSpoken := 0
	numToLastTurn := make(map[int]int)

	for _, num := range input {
		numToLastTurn[num] = rd
		lastSpoken = num
		rd++
	}

	for {
		if rd > 2020 {
			break
		}

		nowSpeaking := 0
		x := numToLastTurn[lastSpoken]
		if x != 0 {
			nowSpeaking = (rd-1) - x
		}
		numToLastTurn[lastSpoken] = rd-1
		lastSpoken = nowSpeaking
		rd++
	}

	return lastSpoken
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(input))
}
