package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Policy struct {
	char byte
	min  int
	max  int
}

func (p Policy) Complies(password string) bool {
	var count int
	for i := 0; i < len(password); i++ {
		if password[i] == p.char {
			count++
		}
	}
	return p.min <= count && count <= p.max
}

// read a Grid in from stdin
func readStdin() (count int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var p Policy
		line := []byte(strings.Trim(scanner.Text(), "\r\n"))
		fields := strings.Fields(string(line))
		minmax := strings.Split(fields[0], "-")
		p.min, err = strconv.Atoi(minmax[0])
		if err != nil {
			return
		}
		p.max, err = strconv.Atoi(minmax[1])
		if err != nil {
			return
		}
		p.char = fields[1][0]
		password := fields[2]
		if p.Complies(password) {
			count++
		}
	}
	return
}

func main() {
	count, err := readStdin()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
