package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type empty struct{}

type Pile []uint8

func score(pile Pile) (result int) {
	mult := len(pile)
	for _, c := range pile {
		result += mult * int(c)
		mult--
	}
	return
}

func copyPile(s Pile) (r Pile) {
	r = make(Pile, len(s))
	copy(r, s)
	return
}

type StringSet map[string]empty

func (s StringSet) Add(str string) {
	s[str] = empty{}
}

func (s StringSet) Has(str string) bool {
	_, ok := s[str]
	return ok
}

// convert the state of the game into a hashable key
func makeKey(first, second Pile) (x string) {
	var buf strings.Builder
	buf.Grow(len(first) + len(second) + 1)
	buf.Write([]byte(first[:]))
	// there is no zero card, use it as the delimiter
	buf.WriteByte(0)
	buf.Write([]byte(second[:]))
	return buf.String()
}

func recursiveSolve(first, second Pile) (winningDeck Pile, firstWins bool) {
	cache := make(StringSet)
	for {
		if len(second) == 0 {
			return first, true
		} else if len(first) == 0 {
			return second, false
		}

		key := makeKey(first, second)
		if cache.Has(key) {
			return first, true
		}
		cache.Add(key)

		f := first[0]
		first = first[1:]
		s := second[0]
		second = second[1:]

		var firstWinsRound bool
		if int(f) > len(first) || int(s) > len(second) {
			firstWinsRound = f >= s
		} else {
			_, firstWinsRound = recursiveSolve(copyPile(first[:f]), copyPile(second[:s]))
		}
		if firstWinsRound {
			first = append(first, f)
			first = append(first, s)
		} else {
			second = append(second, s)
			second = append(second, f)
		}
	}
}

func solve(input []string) (result int) {
	var first, second Pile
	isFirst := true
	for _, line := range input {
		if line == "" {
			isFirst = false
			continue
		}
		num, err := strconv.Atoi(line)
		if err == nil {
			if num <= 0 || num > 255 {
				panic(num)
			}
			if isFirst {
				first = append(first, uint8(num))
			} else {
				second = append(second, uint8(num))
			}
		}
	}

	winningPile, _ := recursiveSolve(first, second)
	return score(winningPile)
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
