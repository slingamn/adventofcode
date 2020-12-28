package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"errors"
)

var (
	fieldToIdx = map[string]int{
		"byr": 0,
		"iyr": 1,
		"eyr": 2,
		"hgt": 3,
		"hcl": 4,
		"ecl": 5,
		"pid": 6,
		"cid": 7,
	}

	ErrBadLine = errors.New("bad line")
)

const (
	numFields = 8
)

type passport [numFields]bool

func readStdin() (count int, err error) {
	var full, cidless passport
	for i := 0; i < numFields; i++ {
		full[i] = true
	}
	cidless = full
	cidless[fieldToIdx["cid"]] = false

	var p passport
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\r\n")

		if len(line) == 0 {
			if p == full || p == cidless {
				count++
			}
			p = passport{}
			continue
		}

		for _, field := range strings.Fields(line) {
			colonIdx := strings.IndexByte(field, ':')
			if colonIdx == -1 {
				err = ErrBadLine
				return
			}
			t := field[:colonIdx]
			idx, ok := fieldToIdx[t]
			if ok {
				p[idx] = true
			} else {
				err = ErrBadLine
				return
			}
		}
	}

	if p == full || p == cidless {
		count++
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
