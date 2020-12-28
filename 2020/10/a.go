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

	return
}

func solve(output []int) (result int) {
	sort.Ints(output)
	one := 0
	three := 0

	cur := output[0]
	for _, val := range output[1:] {
		switch val - cur {
		case 1:
			one++
		case 3:
			three++
		}
		cur = val
	}
	three++
	return one * three
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(input))
}
