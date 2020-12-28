package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Implements Advent of Code 2019, day 22, part b
// https://adventofcode.com/2019/day/22

// Python-style modulus function
func modulus(a, m int64) int64 {
	result := a % m
	if result < 0 {
		result += m
	}
	return result
}

// compute x*y % m, avoiding integer overflow for intermediate values
func modularMultiply(x, y, m int64) (result int64) {
	negate := y < 0
	if negate {
		y = -y
	}
	y64 := uint64(y)

	for y64 != 0 {
		if (y64 & 1) != 0 {
			result = modulus(result+x, m)
		}
		x = modulus(x*2, m)
		y64 = y64 >> 1
	}

	if negate {
		result = modulus(-1 * result, m)
	}
	return
}

// https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm
func extendedEuclideanAlgorithm(a, b int64) (x, y int64) {
	var r, s, t, old_r, old_s, old_t int64

	old_r, r = a, b
	old_s, s = 1, 0
	old_t, t = 0, 1

	for {
		if r == 0 {
			return old_s, old_t
		}
		quotient := old_r / r
		old_r, r = r, (old_r - quotient * r)
		old_s, s = s, (old_s - quotient * s)
		old_t, t = t, (old_t - quotient * t)
	}
}

// compute the multiplicative inverse of a, mod p
func multiplicativeInverse(a, p int64) int64 {
	x, _ := extendedEuclideanAlgorithm(a, p)
	return modulus(x, p)
}

// all shuffles are ax + b (mod p)
type Shuffle struct {
	a int64
	b int64
}

func (s Shuffle) String() string {
	return fmt.Sprintf("f(x) = %dx + (%d)", s.a, s.b)
}

func (s Shuffle) Evaluate(x, deckSize int64) int64 {
	return modulus(modularMultiply(s.a, x, deckSize) + s.b, deckSize)
}

// return the function composition, i.e., f such that f(x) = second(first(x))
func compose(first, second Shuffle, deckSize int64) (result Shuffle) {
	coefficient := modularMultiply(second.a, first.a, deckSize)
	constant := modularMultiply(second.a, first.b, deckSize) + second.b
	return Shuffle{a: modulus(coefficient, deckSize), b: modulus(constant, deckSize)}
}

// return the function composition of s with itself, exp times
func exponentiate(s Shuffle, exp uint64, deckSize int64) (result Shuffle) {
	exponentiatedShuffle := s
	result = Shuffle{1, 0}
	for i := 0; i < 64; i++ {
		if (exp & (1 << i)) != 0 {
			result = compose(result, exponentiatedShuffle, deckSize)
		}
		exponentiatedShuffle = compose(exponentiatedShuffle, exponentiatedShuffle, deckSize)
	}
	return
}

// return the inverse function of s
func inverse(s Shuffle, deckSize int64) (result Shuffle) {
	a_inverse := multiplicativeInverse(s.a, deckSize)
	return Shuffle{a_inverse, modulus(-1 * modularMultiply(a_inverse, s.b, deckSize), deckSize)}
}

func parseInt64(str string) (result int64) {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return
}

func readStdin() (result []Shuffle) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Fields(strings.TrimSpace(scanner.Text()))
		var shuffle Shuffle
		if parts[0] == "cut" {
			param := parseInt64(parts[1])
			shuffle = Shuffle{1, -1 * param}
		} else if parts[0] == "deal" {
			if parts[1] == "into" {
				shuffle = Shuffle{-1, -1}
			} else if parts[1] == "with" {
				param := parseInt64(parts[3])
				shuffle = Shuffle{param, 0}
			} else {
				panic(parts[1])
			}
		} else {
			panic(parts[0])
		}

		result = append(result, shuffle)
	}

	return
}

func main() {
	var pos int64 = 2020
	var deckSize int64 = 119315717514047
	var exponent uint64 = 101741582076661

	shuffles := readStdin()
	shuffle := Shuffle{1, 0}
	for _, newShuffle := range shuffles {
		shuffle = compose(shuffle, newShuffle, deckSize)
	}
	fmt.Printf("composed shuffles to %s\n", shuffle.String())

	expShuffle := exponentiate(shuffle, exponent, deckSize)
	fmt.Printf("exponentiated shuffle to %s\n", expShuffle.String())

	inverseShuffle := inverse(expShuffle, deckSize)
	fmt.Printf("inverted shuffle to %s\n", inverseShuffle)
	answer := inverseShuffle.Evaluate(pos, deckSize)
	fmt.Printf("answer: %d\n", answer)

	invertedAnswer := expShuffle.Evaluate(answer, deckSize)
	if invertedAnswer != pos {
		panic(invertedAnswer)
	}
}
