package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Python-style modulus function
func modulus(a, m int) int {
	result := a % m
	if result < 0 {
		result += m
	}
	return result
}

// compute x**y % m, avoiding integer overflow for intermediate values
func modularExponentiate(x, y, m int) (result int) {
	negate := y < 0
	if negate {
		y = -y
	}
	y64 := uint64(y)

	result = 1
	for y64 != 0 {
		if (y64 & 1) != 0 {
			result = modulus(result*x, m)
		}
		x = modulus(x*x, m)
		y64 = y64 >> 1
	}

	if negate {
		result = modulus(-1*result, m)
	}
	return
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

const (
	sharedPrime = 20201227
	subj = 7
)

func solve(input []string) (result int) {
	cardPub := parseInt(input[0])
	doorPub := parseInt(input[1])

	success := false
	cardExp := 1
	cardPower := subj
	for ; cardExp < sharedPrime; {
		if cardPub == cardPower {
			success = true
			break
		}
		cardExp++
		cardPower = (cardPower * subj) % sharedPrime
	}

	if !success {
		panic("not found")
	}

	return modularExponentiate(doorPub, cardExp, sharedPrime)
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
