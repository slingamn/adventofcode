package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type empty struct{}

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

	if len(fields) != len(myTicket) {
		panic("invalid")
	}

	scanner.Scan()
	scanner.Scan()

	var nearby []Ticket
	for scanner.Scan() {
		ticket := ParseTicket(scanner.Text())

		// only consider tickets where each value is valid for some field
		valid := true
		for _, val := range ticket {
			field_valid := false
			for _, constraint := range fields {
				if constraint.Satisfies(val) {
					field_valid = true
					break
				}
			}
			if !field_valid {
				valid = false
				break
			}
		}
		if valid {
			nearby = append(nearby, ParseTicket(scanner.Text()))
		}
	}

	// adjacencyMatrix[i][j] = "is field i compatible with position j?"
	adjacencyMatrix := make([][]bool, len(myTicket))
	for i, _ := range adjacencyMatrix {
		constraints := fields[i]
		adjacencyMatrix[i] = make([]bool, len(myTicket))
		for j := 0; j < len(myTicket); j++ {
			valid := true
			for _, n := range nearby {
				if !constraints.Satisfies(n[j]) {
					valid = false
					break
				}
			}
			adjacencyMatrix[i][j] = valid
		}
	}

	fieldToIndex := make(map[int]int)
	unassignedFields := make(map[int]empty)
	unassignedIndices := make(map[int]empty)
	for i := 0; i < len(fields); i++ {
		unassignedFields[i] = empty{}
		unassignedIndices[i] = empty{}
	}

	// greedy algorithm: during each round, look for the field
	// that can only match to a single remaining index, then assign it
	// (this is actually correct in general, assuming that a unique
	// perfect matching exists)
	for len(unassignedFields) != 0 {
		progress := false
		for i := range unassignedFields {
			possibleCnt := 0
			candidate := -1
			for j := range unassignedIndices {
				if adjacencyMatrix[i][j] {
					candidate = j
					possibleCnt++
				}
			}
			if possibleCnt == 1 {
				fieldToIndex[i] = candidate
				delete(unassignedFields, i)
				delete(unassignedIndices, candidate)
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
		result *= myTicket[fieldToIndex[fi]]
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
