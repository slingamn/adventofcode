package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sort"
	//"strings"
)

var (
	ErrBadLine = errors.New("bad line")
)

type empty struct{}

func readStdin() (out []int, err error) {
	out = append(out, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}

	sort.Ints(out)
	out = append(out, out[len(out)-1]+3)

	return
}

func dynSolve(i int, input []int, cache []int) {
	result := 0

	if i == len(input) - 1 {
		result = 1
	}

	for j := i+1 ; j < len(input); j++ {
		if input[j] - input[i] <= 3 {
			result += cache[j]
		} else {
			break
		}
	}

	cache[i] = result
}

func solve(output []int) (result int) {
	cache := make([]int, len(output))
	for i := len(output) - 1; i >= 0; i-- {
		dynSolve(i, output, cache)
	}
	return cache[0]
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(input))
}
