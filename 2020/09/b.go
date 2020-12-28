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

func minmax(r []int) (min, max int) {
	for _, x := range r {
		if min == 0 || x < min {
			min = x
		}
		if max == 0 || x > max {
			max = x
		}
	}
	return
}

func solve(val int, input []int) (result int) {
	for i := 0; i < len(input); i++ {
		candidate := input[i]
		for j := i + 1; j < len(input); j++ {
			candidate += input[j]
			if candidate == val {
				r := input[i : j+1]
				min, max := minmax(r)
				return min + max
			} else if candidate > val {
				break
			}
		}
	}

	return -1
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	solution := solve(1212510616, input)
	fmt.Println(solution)
}
