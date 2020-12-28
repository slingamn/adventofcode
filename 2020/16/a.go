package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Constraint struct {
	min int
	max int
}

func (c Constraint) Satisfies(i int) bool {
	return c.min <= i && i <= c.max
}

type Constraints []Constraint

func (cs Constraints) Satisfies(i int) bool {
	for _, c := range cs {
		if c.Satisfies(i) {
			return true
		}
	}
	return false
}

type Fields []Constraints

type Ticket []int

func readStdin() (result int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	var fields Fields
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		colonIdx := strings.IndexByte(line, ':')
		if colonIdx == -1 {
			panic("fail")
		}
		line = line[colonIdx+2:]
		constStrs := strings.Split(line, " or ")
		var constraints Constraints
		for _, cs := range constStrs {
			var c Constraint
			dashIdx := strings.IndexByte(cs, '-')
			if dashIdx == -1 {
				panic("fail")
			}
			c.min, err = strconv.Atoi(cs[:dashIdx])
			if err != nil {
				panic(err)
			}
			c.max, err = strconv.Atoi(cs[dashIdx+1:])
			if err != nil {
				panic(err)
			}
			constraints = append(constraints, c)
		}
		fields = append(fields, constraints)
	}

	for i := 0; i < 4; i++ {
		scanner.Scan()
		scanner.Text()
	}

	for scanner.Scan() {
		var ticket Ticket
		tstrs := strings.Split(scanner.Text(), ",")
		for _, tstr := range tstrs {
			tint, err := strconv.Atoi(tstr)
			if err != nil {
				panic(err)
			}
			ticket = append(ticket, tint)
		}

		for _, f := range ticket {
			valid := false
			for _, constraint := range fields {
				if constraint.Satisfies(f) {
					valid = true
					break
				}
			}
			if !valid {
				result += f
			}
		}
	}

	return
}

func solve(input []string) (result int) {
	return len(input)
}

func main() {
	input, err := readStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(input)
}
