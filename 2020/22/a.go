package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func score(pile []int) (result int) {
	mult := len(pile)
	for _, c := range pile {
		result += mult*c
		mult--
	}
	return
}

func solve(input []string) (result int) {
	var first, second []int
	isFirst := true
	for _, line := range input {
		if line == "" {
			isFirst = false
			continue
		}
		num, err := strconv.Atoi(line)
		if err == nil {
			if isFirst {
				first = append(first, num)
			} else {
				second = append(second, num)
			}
		}
	}

	for {
		if len(first) == 0 {
			return score(second)
		} else if len(second) == 0 {
			return score(first)
		}

		f := first[0]
		first = first[1:]
		s := second[0]
		second = second[1:]
		if f >= s {
			first = append(first, f)
			first = append(first, s)
		} else {
			second = append(second, s)
			second = append(second, f)
		}
	}

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
