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

func ParseTicket(line string) (ticket Ticket) {
	tstrs := strings.Split(line, ",")
	for _, tstr := range tstrs {
		tint, err := strconv.Atoi(tstr)
		if err != nil {
			panic(err)
		}
		ticket = append(ticket, tint)
	}
	return
}

func readStdin() (result int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	var fields Fields
	var departureFields []int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if strings.HasPrefix(line, "departure ") {
			departureFields = append(departureFields, len(fields))
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

	scanner.Scan()

	scanner.Scan()
	myTicket := ParseTicket(scanner.Text())

	scanner.Scan()
	scanner.Scan()

	var nearby []Ticket
	for scanner.Scan() {
		nearby = append(nearby, ParseTicket(scanner.Text()))
	}

	constraintToTicketIdx := make(map[int]int)
	unassignedConstraints := make(map[int]bool)
	unassignedFields := make(map[int]bool)
	if len(fields) != len(myTicket) {
		panic("invalid")
	}
	for i := 0; i < len(fields); i++ {
		unassignedConstraints[i] = true
		unassignedFields[i] = true
	}

	for len(unassignedConstraints) != 0 {
		progress := false
		for cIdx := range unassignedConstraints {
			constraints := fields[cIdx]
			possibleCnt := 0
			candidate := -1
			for i := range unassignedFields {
				valid := true
				for _, nt := range nearby {
					if !constraints.Satisfies(nt[i]) {
						valid = false
						break
					}
				}
				if valid {
					possibleCnt++
					candidate = i
				}
			}
			if possibleCnt == 1 {
				constraintToTicketIdx[cIdx] = candidate
				delete(unassignedConstraints, cIdx)
				delete(unassignedFields, candidate)
				progress = true
				break
			}
		}
		if !progress {
			panic("failed to make progress")
		}
	}

	result = 1
	for _, fi := range departureFields {
		result *= myTicket[constraintToTicketIdx[fi]]
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
