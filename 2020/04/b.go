package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"errors"
	"regexp"
	"strconv"
)

var hairRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)

var pidRegex = regexp.MustCompile(`^[0-9]{9}$`)

type empty struct{}

var (
	fieldToIdx = map[string]int{
		"byr": 0,
		"iyr": 1,
		"eyr": 2,
		"hgt": 3,
		"hcl": 4,
		"ecl": 5,
		"pid": 6,
	}

	eyeColors = map[string]empty {
		"amb": empty{},
		"blu": empty{},
		"brn": empty{},
		"gry": empty{},
		"grn": empty{},
		"hzl": empty{},
		"oth": empty{},
	}

	ErrBadLine = errors.New("bad line")
)

const (
	numFields = 7
)

type passport [numFields]bool

func readStdin() (count int, err error) {
	var full passport
	for i := 0; i < numFields; i++ {
		full[i] = true
	}

	var p passport
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\r\n")

		if len(line) == 0 {
			if p == full {
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
			strValue := field[colonIdx+1:]
			switch t {
			case "byr":
				val, err := strconv.Atoi(strValue)
				p[0] = (err == nil && 1920 <= val && val <= 2002)
			case "iyr":
				val, err := strconv.Atoi(strValue)
				p[1] = (err == nil && 2010 <= val && val <= 2020)
			case "eyr":
				val, err := strconv.Atoi(strValue)
				p[2] = (err == nil && 2020 <= val && val <= 2030)
			case "hgt":
				var in, cm bool
				if strings.HasSuffix(strValue, "in") {
					strValue = strings.TrimSuffix(strValue, "in")
					in = true
				} else if strings.HasSuffix(strValue, "cm") {
					strValue = strings.TrimSuffix(strValue, "cm")
					cm = true
				}
				val, err := strconv.Atoi(strValue)
				result := false
				if err == nil {
					if in {
						result = (59 <= val && val <= 76)
					} else if cm {
						result = (150 <= val && val <= 193)
					}
				}
				p[3] = result
			case "hcl":
				p[4] = hairRegex.MatchString(strValue)
			case "ecl":
				_, p[5] = eyeColors[strValue]
			case "pid":
				p[6] = pidRegex.MatchString(strValue)
			}
		}
	}

	if p == full {
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
