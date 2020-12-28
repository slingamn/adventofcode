package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	//"sort"
	"strconv"
	"strings"
)

var (
	ErrBadLine = errors.New("bad line")
)

// https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm
func extendedEuclideanAlgorithm(a, b int) (x, y int) {
	var r, s, t, old_r, old_s, old_t int

	old_r, r = a, b
	old_s, s = 1, 0
	old_t, t = 0, 1

	for {
		if r == 0 {
			return old_s, old_t
		}
		quotient := old_r / r
		old_r, r = r, (old_r - quotient * r)
		old_s, s = s, (old_s - quotient * s)
		old_t, t = t, (old_t - quotient * t)
	}
}

// Python-style modulus function (-1 % 5 == 4, not -1)
func modulus(i int, m int) (r int) {
	r = i % m
	if r < 0 {
		r += m
	}
	return
}

// a is congruent to c mod m
type Constraint struct {
	c int
	m int
}

func readStdin() (est int, buses []Constraint, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	estStr := scanner.Text()
	est, err = strconv.Atoi(estStr)
	if err != nil {
		return
	}
	scanner.Scan()
	busStr := scanner.Text()
	busVals := strings.Split(busStr, ",")
	var bus int
	for c, busVal := range busVals {
		bus, err = strconv.Atoi(busVal)
		if err != nil {
			continue
		}
		buses = append(buses, Constraint{modulus(bus - c, bus), bus})
	}

	return
}

func solve(est int, constraints []Constraint) (result int) {
	// https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Existence_(direct_construction)
	prod := 1
	for _, co := range constraints {
		prod *= co.m
	}
	n := make([]int, len(constraints))
	for i, co := range constraints {
		n[i] = prod/co.m
	}
	m := make([]int, len(constraints))
	for i, co := range constraints {
		bigm, _ := extendedEuclideanAlgorithm(n[i], co.m)
		m[i] = bigm
	}

	for i, co := range constraints {
		result += co.c * m[i] * n[i]
	}

	result = modulus(result, prod)

	/*
	for _, co := range constraints {
		if result % co.m != co.c {
			panic(fmt.Sprintf("%d did not satisfy %#v\n", result, co))
		}
	}
	*/

	return
}

func main() {
	est, buses, err := readStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(est, buses))
}
