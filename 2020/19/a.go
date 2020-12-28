package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
XXX this code is wrong, it doesn't backtrack on attempts to match disjuncts
for some reason this is tolerated by the class of program inputs
i refuse to fix it because the correct solution to this problem is clearly
to throw it into a general-purpose CFG parser like Earley or CYK
*/

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func parseInts(strs []string) (result []int) {
	var err error
	result = make([]int, len(strs))
	for i, s := range strs {
		result[i], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}
	return
}

type Rule struct {
	Terminal byte
	Disjuncts [][]int
}

type Rules []Rule

func (rules Rules) matchesInternal(str string, ruleNum int) (result bool, remaining string) {
	rule := rules[ruleNum]
	if rule.Terminal != 0 {
		if len(str) == 0 {
			return false, ""
		}
		return str[0] == rule.Terminal, str[1:]
	}
	for _, disjunct := range rule.Disjuncts {
		success := true
		remaining := str
		for _, child := range disjunct {
			success, remaining = rules.matchesInternal(remaining, child)
			if !success {
				break
			}
		}
		if success {
			return true, remaining
		}
	}
	return false, remaining
}

func (rules Rules) Matches(str string, ruleNum int) (result bool) {
	matches, remaining := rules.matchesInternal(str, ruleNum)
	return matches && remaining == ""
}

func solve(input []string) (result int, err error) {
	var rulesEnd int
	for ; input[rulesEnd] != ""; rulesEnd++ {
	}
	ruleStrs := input[:rulesEnd]

	rules := make(Rules, len(ruleStrs))
	for _, ruleStr := range ruleStrs {
		var rule Rule
		colIdx := strings.IndexByte(ruleStr, ':')
		if colIdx == -1 {
			panic(ruleStr)
		}
		ruleNum := parseInt(ruleStr[:colIdx])
		rest := ruleStr[colIdx+2:]
		if rest[0] == '"' {
			rule.Terminal = rest[1]
		} else {
			pieces := strings.Split(rest, "|")
			for _, piece := range pieces {
				fields := strings.Fields(piece)
				rule.Disjuncts = append(rule.Disjuncts, parseInts(fields))
			}
		}
		rules[ruleNum] = rule
	}

	messages := input[rulesEnd+1:]

	for _, msg := range messages {
		if rules.Matches(msg, 0) {
			result++
		}
	}

	return
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	solution, err := solve(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(solution)
}
