package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// mercifully abandoned attempt to brute-force part 2

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

const (
	roundCount = 10000000
	minCup     = 1
	maxCup     = 1000000
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

func (pos *Position) toResult() (res int) {
	res = 1
	var start int
	for start = 0; start < len(pos); start++ {
		if pos[start] == 1 {
			break
		}
	}
	j := (start+1) % len(pos)
	res *= pos[j]
	j = (start+2) % len(pos)
	res *= pos[j]
	return
}

func solve(input []string) (result int) {
	var pos Position

	in := input[0]
	for i, x := range in {
		pos[i] = parseInt(string(x))
	}
	for i := 9; i < len(pos); i++ {
		pos[i] = i+1
	}

	last := time.Now()
	for i := 0; i < roundCount; i++ {
		if i % 1000 == 0 {
			fmt.Println(100 * float64(i) / roundCount)
			now := time.Now()
			fmt.Println(now.Sub(last))
			last = now
		}
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

	return pos.toResult()
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
