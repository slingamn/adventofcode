package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"errors"
)

var (
	ErrBadLine = errors.New("bad line")
)

func max(rows []int) (result int) {
	result = -1
	for _, row := range rows {
		if row > result {
			result = row
		}
	}
	return result
}

func shifter(in string, set byte) (result int) {
	end := len(in) - 1
	for i := 0; i <= end; i++ {
		if in[i] == set {
			result = result | 1
		}
		if i != end {
			result = result << 1
		}
	}
	return
}

func readStdin() (seats []int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\r\n")
		if len(line) != 10 {
			return nil, ErrBadLine
		}

		row := shifter(line[:7], 'B')
		seatnum := shifter(line[7:10], 'R')

		seats = append(seats, (row*8) + seatnum)
	}

	return
}

func main() {
	rows, err := readStdin()
	if err != nil {
		panic(err)
	}
	fmt.Println(max(rows))
}
