package main

import (
	"bufio"
	"fmt"
	"os"
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

const (
	roundCount = 100
	minCup = 1
	maxCup = 9
)

type Position [maxCup]int

func fixup(idx int) (res int) {
	switch idx {
	case -1:
		return maxCup-1
	case 1:
		return 4
	case 2, 3:
		panic(idx)
	case maxCup:
		return 0
	default:
		return idx
	}
}

func toString(pos Position) (str string) {
	var buf strings.Builder
	var start int
	for start = 0; start < len(pos); start++ {
		if pos[start] == 1 {
			break
		}
	}
	for j := start+1; j != start; j = ((j+1)%len(pos)) {
		buf.WriteString(strconv.Itoa(pos[j]))
	}
	return buf.String()
}

func solve(input []string) (result int) {
	var pos Position

	in := input[0]
	for i, x := range in {
		pos[i] = parseInt(string(x))
	}

	for i := 0; i < roundCount; i++ {
		cur := pos[0]
		dst := cur
		for {
			dst -= 1
			if dst < minCup {
				dst = maxCup
			}
			if dst != cur && dst != pos[1] && dst != pos[2] && dst != pos[3] {
				break
			}
		}
		var nextPos Position
		j := 0
		x := 4
		for ; j < len(nextPos); {
			nextPos[j] = pos[x]
			atDst := pos[x] == dst
			j++
			x = fixup(x+1)
			if atDst {
				for k := 0; k < 3; k++ {
					nextPos[j] = pos[k+1]
					j++
				}
			}
		}
		pos = nextPos
	}

	fmt.Println(toString(pos))

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
