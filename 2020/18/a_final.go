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

func trimSpace(str string) (remaining string) {
	return strings.TrimLeft(str, " ")
}

func atEnd(str string) bool {
	return len(str) == 0 || str[0] == ')'
}

func consumeExpr(remaining string) (firstOperand int, rest string) {
	digits := number.FindString(remaining)
	if digits != "" {
		firstOperand = parseInt(digits)
		remaining = remaining[len(digits):]
	} else {
		if remaining[0] != '(' {
			panic(remaining)
		}
		firstOperand, remaining = evaluate(remaining[1:])
		remaining = trimSpace(remaining)
		if remaining != "" {
			if remaining[0] != ')' {
				panic(remaining)
			} else {
				remaining = remaining[1:]
			}
		}
	}
	return firstOperand, remaining
}

func evaluate(expr string) (result int, remaining string) {
	remaining = trimSpace(expr)

	result, remaining = consumeExpr(remaining)
	remaining = trimSpace(remaining)
	if atEnd(remaining) {
		return
	}

	for {
		var operand int
		operator := remaining[0]
		remaining = remaining[1:]

		remaining = trimSpace(remaining)
		operand, remaining = consumeExpr(remaining)

		switch operator {
		case '+':
			result += operand
		case '*':
			result *= operand
		default:
			panic(operator)
		}

		remaining  = trimSpace(remaining)
		if atEnd(remaining) {
			return
		}
	}
}

func solve(input []string) (result int, err error) {
	for _, i := range input {
		val, remaining := evaluate(i)
		if remaining != "" {
			panic(remaining)
		}
		result += val
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
