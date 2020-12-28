package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
)

var (
	ErrBadLine = errors.New("bad line")
)

const (
	fullForm = (1 << 26) - 1
)

func count(g uint32) (result int) {
	for i := 0; i < 26; i++ {
		if g & (1 << i) != 0 {
			result++
		}
	}
	return
}

func readStdin() (result []int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	var groupForm uint32
	groupForm = fullForm
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			result = append(result, count(groupForm))
			groupForm = fullForm
			continue
		}
		var form uint32
		for i := 0; i < len(line); i++ {
			form = form | (1 << (line[i] - 'a'))
		}
		groupForm = groupForm & form
	}
	result = append(result, count(groupForm))

	return
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, g := range input {
		sum += g
	}
	fmt.Println(sum)

}
