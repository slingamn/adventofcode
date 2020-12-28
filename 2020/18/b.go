package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

var (
	number = regexp.MustCompile(`^[0-9]+`)
)

func consumeExpr(remaining string) (firstOperand int, rest string) {
		digit := number.FindString(remaining)
		if digit != "" {
			firstOperand = parseInt(digit)
			remaining = remaining[len(digit):]
		} else {
			if remaining[0] != '(' {
				panic(remaining)
			}
			openParenCount := 1
			endIdx := 1
			for ; endIdx < len(remaining); endIdx++ {
				switch remaining[endIdx] {
				case '(':
					openParenCount++
				case ')':
					openParenCount--
				}
				if openParenCount == 0 {
					break
				}
			}
			if openParenCount != 0 {
				panic(remaining)
			}
			firstOperand = evaluate(remaining[1:endIdx])
			remaining = remaining[endIdx+1:]
		}
		return firstOperand, remaining
}

func evaluate(expr string) (result int) {
	if len(expr) == 0 {
		panic(expr)
	}

	remaining := expr

	var reduced []int
	var operators []byte

	for len(remaining) != 0 {
		var firstOperand int
		firstOperand, remaining = consumeExpr(remaining)
		reduced = append(reduced, firstOperand)

		if len(remaining) == 0 {
			break
		}

		remaining = strings.TrimLeft(remaining, " ")
		operator := remaining[0]
		remaining = remaining[1:]
		operators = append(operators, operator)
		remaining = strings.TrimLeft(remaining, " ")
	}

	if len(reduced) == 1 {
		return reduced[0]
	}
	var multiplicands []int
	var accumulator int
	accumulator = reduced[0]
	reduced = reduced[1:]
	for i, o := range operators {
		switch o {
		case '+':
			accumulator += reduced[i]
		case '*':
			multiplicands = append(multiplicands, accumulator)
			accumulator = reduced[i]
		default:
			panic(expr)
		}
	}
	multiplicands = append(multiplicands, accumulator)

	result = 1
	for _, m := range multiplicands {
		result *= m
	}
	return
}

func solve(input []string) (result int, err error) {
	for _, i := range input {
		result += evaluate(i)
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

	solution, err := solve(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(solution)
}
