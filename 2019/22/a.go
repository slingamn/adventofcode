package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"log"
)

// Implements Advent of Code 2019, day 22, part a
// https://adventofcode.com/2019/day/22

type ShuffleType uint

const (
	Reverse ShuffleType = iota
	Cut
	Increment
)

type Shuffle struct {
	Type ShuffleType
	Param int
}

func (s Shuffle) String() string {
	return fmt.Sprintf("Type=%d, Param=%d", s.Type, s.Param)
}

func reposition(pos, deckSize int, shuffle Shuffle) int {
	switch shuffle.Type {
	case Reverse:
		return (deckSize - 1) - pos
	case Cut:
		pos = pos - shuffle.Param
		if pos < 0 {
			pos += deckSize
		} else if pos >= deckSize {
			pos -= deckSize
		}
		return pos
	case Increment:
		return (pos * shuffle.Param) % deckSize
	default:
		panic("unreachable")
	}
}

// read a Grid in from stdin
func readStdin() (result []Shuffle) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Fields(strings.TrimSpace(scanner.Text()))
		var shuffle Shuffle
		if parts[0] == "cut" {
			param, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			shuffle = Shuffle{Type: Cut, Param: int(param)}
		} else if parts[0] == "deal" {
			if parts[1] == "into" {
				shuffle = Shuffle{Type: Reverse}
			} else if parts[1] == "with" {
				param, err := strconv.Atoi(parts[3])
				if err != nil {
					log.Fatal(err)
				}
				shuffle = Shuffle{Type: Increment, Param: param}
			} else {
				log.Fatal("bad deal", parts[1])
			}
		} else {
			log.Fatal("bad line", parts[0])
		}

		result = append(result, shuffle)
	}

	return
}

func main() {
	pos := 2019
	deckSize := 10007
	if len(os.Args) > 1 {
		var err error
		pos, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		deckSize, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}
	shuffles := readStdin()
	for _, shuffle := range shuffles {
		pos = reposition(pos, deckSize, shuffle)
	}
	fmt.Println(pos)
}
