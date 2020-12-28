package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type empty struct{}

func score(pile []int) (result int) {
	mult := len(pile)
	for _, c := range pile {
		result += mult*c
		mult--
	}
	return
}

func copyInts(s []int) (r []int) {
	r = make([]int, len(s))
	copy(r, s)
	return
}

type Decks [2][50]int
type DeckCache map[Decks]empty

func (d DeckCache) Has(x Decks) bool {
	_, ok := d[x]
	return ok
}

func (d DeckCache) Add(x Decks) {
	d[x] = empty{}
}

func makeKey(first, second []int) (x Decks) {
	copy(x[0][:], first)
	copy(x[1][:], second)
	return
}

func recursiveSolve(first, second []int) ([]int, []int, bool) {
	cache := make(DeckCache)
	for {
		if len(second) == 0 {
			return first, second, true
		} else if len(first) == 0 {
			return first, second, false
		}

		key := makeKey(first, second)
		if cache.Has(key) {
			return first, second, true
		}
		cache.Add(key)

		f := first[0]
		first = first[1:]
		s := second[0]
		second = second[1:]

		var roundWinner bool
		if f > len(first) || s > len(second) {
			roundWinner = f >= s
		} else {
			_, _, roundWinner = recursiveSolve(copyInts(first[:f]), copyInts(second[:s]))
		}
		if roundWinner {
			first = append(first, f)
			first = append(first, s)
		} else {
			second = append(second, s)
			second = append(second, f)
		}
	}
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

	first, second, winner := recursiveSolve(first, second)
	if winner {
		return score(first)
	} else {
		return score(second)
	}
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
